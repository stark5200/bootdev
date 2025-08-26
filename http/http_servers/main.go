package main

import (
	"chirpy/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	//"hash"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	//"context"
	"chirpy/internal/auth"
	"time"

	"github.com/google/uuid"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//"path/filepath"

func main() {
	port := "8080"
	filepath := "."

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM environment variable is not set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()
	dbQueries := database.New(dbConn)
	
	fileServer := http.FileServer(http.Dir(filepath))

	apiCfg := apiConfig{fileserverHits: atomic.Int32{}, db: dbQueries, platform: platform}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", fileServer)))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerGetChirpsID)
	//mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerCreateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)

	http_server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("✅ created server and actively listening on port %s! and fileroot %s", port, filepath)
	log.Fatal(http_server.ListenAndServe())
}

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform			 string
}

type chirp struct {
	Body    string `json:"body"`
	UserId uuid.UUID `json:"user_id"`
}

type cleanedChirp struct {
	CleanedBody    string `json:"cleaned_body"`
	UserId         uuid.UUID `json:"user_id"`
}
type returnBody struct {
	Valid    bool `json:"valid"`
}

type email struct {
	Email    string `json:"email"`
}

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	HashedPassword string   `json:"hashed_password"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
} 

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.db.Reset(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset the database: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state."))
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	numVisits := cfg.fileserverHits.Load()
    html := fmt.Sprintf(`
    <html>
      <body>
        <h1>Welcome, Chirpy Admin</h1>
        <p>Chirpy has been visited %d times!</p>
      </body>
    </html>
    `, numVisits)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
	log.Printf("✅ served metrics page with %d hits", numVisits)
} 


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
/*
func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	const maxChirpLength = 140

	decoder := json.NewDecoder(r.Body)
	chirpData := chirp{}
	err := decoder.Decode(&chirpData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode chirpData body", err)
		log.Printf("Error decoding chirpData body: %s", err)
		return
	}

	if len(chirpData.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil) // Respond with 400 Bad Request
		log.Printf("Chirp is too long: %s", chirpData.Body)
		return
	}
	
	// respBody := returnBody{
	//	 Valid: true,
	// }
	
	respondWithChirpJSON(w, http.StatusOK, chirpData) // Respond with 200 OK
}
*/

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {

	type fullChirp struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Body     string     `json:"body"`
		UserID  uuid.UUID  `json:"user_id"`
	}

	chirpsFromDB, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get chirps", err)
		return
	}

	chirps := make([]fullChirp, len(chirpsFromDB))
	for i, c := range chirpsFromDB {
		chirps[i] = fullChirp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:     c.Body,
			UserID:   c.UserID,
		}
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerGetChirpsID(w http.ResponseWriter, r *http.Request) {

	ID := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid chirp ID format", err)
		return
	}

	type fullChirp struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Body     string     `json:"body"`
		UserID  uuid.UUID  `json:"user_id"`
	}

	chirpFromDB, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to get chirp by ID", err)
		return
	}

	chirp := fullChirp{
		ID:        chirpFromDB.ID,
		CreatedAt: chirpFromDB.CreatedAt,
		UpdatedAt: chirpFromDB.UpdatedAt,
		Body:     chirpFromDB.Body,
		UserID:   chirpFromDB.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	const maxChirpLength = 140

	type params struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}

	type fullChirp struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Body     string     `json:"body"`
		UserID  uuid.UUID  `json:"user_id"`
	}

	type response struct {
		fullChirp
	}

	decoder := json.NewDecoder(r.Body)
	p := params{}
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode params (chirp) body", err)
		log.Printf("Error decoding params (chirp) body: %s", err)
		return
	}

	if len(p.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil) // Respond with 400 Bad Request
		log.Printf("Chirp is too long: %s", p.Body)
		return
	}

	chirpData := database.CreateChirpParams{
		Body:  p.Body,
		UserID: p.UserId,
	}

	chirpFromDB, err := cfg.db.CreateChirp(r.Context(), chirpData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create chirp", err)
		return
	}

	newChirp := fullChirp{
		ID:        chirpFromDB.ID,
		CreatedAt: chirpFromDB.CreatedAt,
		UpdatedAt: chirpFromDB.UpdatedAt,
		Body:     chirpFromDB.Body,
		UserID:  chirpFromDB.UserID,
	}

	respondWithJSON(w, http.StatusCreated, newChirp)
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	type response struct {
		userResponse
	}
	
	decoder := json.NewDecoder(r.Body)
	p := params{}
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(p.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to hash password", err)
		return
	}

	userData := database.CreateUserParams{
		Email: p.Email,
		PasswordHash: hashedPassword,
	}

	userFromDB, err := cfg.db.CreateUser(r.Context(), userData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	newUserResponse := userResponse{
		ID:        userFromDB.ID,
		CreatedAt: userFromDB.CreatedAt,
		UpdatedAt: userFromDB.UpdatedAt,
		Email:     userFromDB.Email,
	}

	respondWithJSON(w, http.StatusCreated, newUserResponse)
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	decoder := json.NewDecoder(r.Body)
	p := params{}
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userFromDB, err := cfg.db.LoginUserByEmail(r.Context(), p.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	err = auth.CheckPasswordHash(p.Password, userFromDB.PasswordHash)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	loggedInUser := userResponse{
		ID:        userFromDB.ID,
		CreatedAt: userFromDB.CreatedAt,
		UpdatedAt: userFromDB.UpdatedAt,
		Email:     userFromDB.Email,
	}

	respondWithJSON(w, http.StatusOK, loggedInUser)
}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithChirpJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	c, ok := payload.(chirp)
	if !ok {
		log.Printf("Invalid payload type: %T", payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cleaned := cleanedChirp{
		CleanedBody: cleanText(c.Body),
	}

	dat, err := json.Marshal(cleaned)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func cleanText(input string) string {
	banned := []string{"kerfuffle", "sharbert", "fornax"} // expand as needed
	words := strings.Fields(input) // split by whitespace
    for i, w := range words {
        lw := strings.ToLower(w)
        for _, b := range banned {
            if lw == b {
                words[i] = "****"
            }
        }
    }
    return strings.Join(words, " ")
}
