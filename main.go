package main

import (
	"net/http"
)

func main() {
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

	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()

}