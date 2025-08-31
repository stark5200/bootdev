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

	//"github.com/golang-jwt/jwt/v5"
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
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	jwtExpiresInStr := os.Getenv("JWT_EXPIRES_IN")
	if jwtExpiresInStr == "" {
		log.Fatal("JWT_EXPIRES_IN environment variable is not set")
	}
	jwtExpiresIn, err := time.ParseDuration(jwtExpiresInStr)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRES_IN duration: %v", err)
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

	apiCfg := apiConfig{fileserverHits: atomic.Int32{}, db: dbQueries, platform: platform, jwtSecret: jwtSecret, jwtExpiresIn: jwtExpiresIn}

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
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefreshToken)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevokeToken)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerUpdateUser)

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
	platform       string
	jwtSecret      string
	jwtExpiresIn   time.Duration
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
	Token 	 string    `json:"token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
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


	// Try to get token from Authorization header
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or missing token", err)
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or expired token", err)
		return
	}

	if len(p.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil) // Respond with 400 Bad Request
		log.Printf("Chirp is too long: %s", p.Body)
		return
	}

	chirpData := database.CreateChirpParams{
		Body:  p.Body,
		UserID: userID,
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

	userToken, err := auth.MakeJWT(userFromDB.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create JWT", err)
		return
	}

	refreshToken := auth.MakeRefreshToken()

	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID:    userFromDB.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().UTC().Add(60 * 24 * time.Hour), // 60 days
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to store refresh token", err)
		return
	}

	loggedInUser := userResponse{
		ID:        userFromDB.ID,
		CreatedAt: userFromDB.CreatedAt,
		UpdatedAt: userFromDB.UpdatedAt,
		Email:     userFromDB.Email,
		Token:     userToken,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, loggedInUser)
}

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Try to get token from Authorization header
	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or missing token", err)
		return
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or expired token", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	p := params{}
	err = decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userFromDB, err := cfg.db.GetUserByID(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user", err)
		return
	}

	if p.Email != "" {
		userFromDB.Email = p.Email
	}
	if p.Password != "" {
		hashedPassword, err := auth.HashPassword(p.Password)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to hash password", err)
			return
		}
		userFromDB.PasswordHash = hashedPassword
	}
	userFromDB.UpdatedAt = time.Now().UTC()
	/*
	err = auth.CheckPasswordHash(p.Password, userFromDB.PasswordHash) // just to use the password package
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}
	*/

	updatedUserData := database.UpdateUserParams{
		ID:           userFromDB.ID,
		Email:        userFromDB.Email,
		PasswordHash: userFromDB.PasswordHash,
		CreatedAt:    userFromDB.CreatedAt,
		UpdatedAt:    userFromDB.UpdatedAt,
	}

	_, err = cfg.db.UpdateUser(r.Context(), updatedUserData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, userResponse{
		ID:        updatedUserData.ID,
		CreatedAt: updatedUserData.CreatedAt,
		UpdatedAt: updatedUserData.UpdatedAt,
		Email:     updatedUserData.Email,
	})
}

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	type refreshResponse struct {
		RefreshToken string `json:"token"`
	}
	
	refresh_token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing refresh token", nil)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refresh_token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get user for refresh token", err)
		return
	}

	storedToken, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		time.Hour,
	)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}

	refresh_response := refreshResponse{
		RefreshToken: storedToken,
	}

	respondWithJSON(w, http.StatusOK, refresh_response)
}

func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke session", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
