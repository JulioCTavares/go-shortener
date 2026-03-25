package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julioctavares/go-shortener/internal/shortener"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			URL string `json:"url"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		shortURL, err := shortener.CreateShortUrl(body.URL)
		if err != nil {
			http.Error(w, "Failed to create short URL: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
	})

	return mux
}
