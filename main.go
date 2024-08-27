package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apiCfg := apiConfig{}
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	
	mux.HandleFunc("/healthz", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)

	mux.HandleFunc("/reset", apiCfg.handlerReset)


	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()

}