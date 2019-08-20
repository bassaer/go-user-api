package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler := NewHandler()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error("Status code should be 200")
	}

	got := rec.Body.String()
	want := "OK\n"
	if got != want {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}
