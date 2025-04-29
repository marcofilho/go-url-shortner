package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/marcofilho/go-url-shortner/backend/internals/service"
)

type Handler struct {
	Service *service.Service
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	shortID := generateShortID(6)

	err := h.Service.Save(req.URL, shortID)
	if err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	resp := shortenResponse{
		ShortURL: "http://localhost:8080/" + shortID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]

	originalURL, err := h.Service.Get(shortID)
	if err != nil || originalURL == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
