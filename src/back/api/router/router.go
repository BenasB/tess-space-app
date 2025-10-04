package router

import (
	"net/http"

	"github.com/BenasB/tess-space-app/back/api/handler"
	"github.com/BenasB/tess-space-app/back/mast"
)

func New(mdc *mast.DownloadClient) http.Handler {
	apiHandler := &handler.ApiHandler{MastClient: mdc}

	mux := http.NewServeMux()
	mux.HandleFunc("/", apiHandler.Greet)

	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}
