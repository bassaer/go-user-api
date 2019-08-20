package main

import (
	"encoding/json"
	"net/http"
)

// Handler is
type Handler struct {
	repo Repository
}

// NewHandler is
func NewHandler() *Handler {
	return &Handler{
		repo: &UserRepository{},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if user, err := h.repo.Get(id); err == nil {
			bytes, _ := json.Marshal(user)
			w.Write(bytes)
			return
		}
	} else if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
