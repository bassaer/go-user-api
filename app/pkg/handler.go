package app

import (
	"encoding/json"
	"log"
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

func (h *Handler) set(user *User) error {
	return h.repo.Set(user)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if bytes, err := h.get(id); err == nil {
			w.Write(bytes)
			return
		}
	} else if r.Method == http.MethodPost {
		user := &User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := h.set(user); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if bytes, err := json.Marshal(user); err == nil {
			w.Write(bytes)
			return
		}
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
