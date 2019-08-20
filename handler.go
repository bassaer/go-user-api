package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler is
type Handler struct {
	repo Repository
}

// NewHandler is
func NewHandler(r Repository) *Handler {
	return &Handler{
		repo: r,
	}
}

func (h *Handler) get(id string) ([]byte, error) {
	user, err := h.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(user)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if bytes, err := h.get(id); err == nil {
			w.Write(bytes)
			return
		}
	} else if r.Method == http.MethodPost {
		fmt.Fprintf(w, "OK")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
