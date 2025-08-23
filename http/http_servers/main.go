package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"os"
	"database/sql"
	"chirpy/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//"path/filepath"

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	port := "8080"
	filepath := "."
	
	fileServer := http.FileServer(http.Dir(filepath))

	apiCfg := apiConfig{fileserverHits: atomic.Int32{}, db: *dbQueries}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", fileServer)))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)

	http_server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("✅ created server and actively listening on port %s! and fileroot %s", port, filepath)
	log.Fatal(http_server.ListenAndServe())

}

type apiConfig struct {
	fileserverHits atomic.Int32
	db             database.Queries
}

type chirp struct {
	Body    string `json:"body"`
}
type cleanedChirp struct {
	Cleaned_body    string `json:"cleaned_body"`
}
type returnBody struct {
	Valid    bool `json:"valid"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
} 

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
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

	/*
	respBody := returnBody{
		Valid: true,
	}
	*/

	respondWithJSON(w, http.StatusOK, chirpData) // Respond with 200 OK
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

	c, ok := payload.(chirp)
	if !ok {
		log.Printf("Invalid payload type: %T", payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cleaned := cleanedChirp{
		Cleaned_body: cleanText(c.Body),
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
