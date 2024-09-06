package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	mux := http.NewServeMux()
	apiCfg := apiConfig{}
	mux.Handle("GET /app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	
	mux.HandleFunc("GET /api/healthz", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

	mux.HandleFunc("GET /api/reset", apiCfg.handlerReset)

	mux.HandleFunc("POST /api/chirps", createChirp)
	mux.HandleFunc("GET /api/chirps", getChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", getChirp)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", deleteChirp)

	mux.HandleFunc("POST /api/users", createUser)
	mux.HandleFunc("PUT /api/users", updateUser)
	
	mux.HandleFunc("POST /api/login", login)

	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()

}