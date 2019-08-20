package main

import (
	"fmt"
	"net/http"
)

// Handler is
type Handler struct{}

// NewHandler is
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}
