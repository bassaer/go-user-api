package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockRepository struct{}

func (m *MockRepository) Get(id string) (*User, error) {
	return &User{
		ID:        id,
		Name:      "testname",
		CreatedAt: time.Time{},
	}, nil

}

func (m *MockRepository) Set(user User) error {
	return nil
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/?id=testid", nil)
	rec := httptest.NewRecorder()

	handler := NewHandler(&MockRepository{})
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error("Status code should be 200")
	}

	got := rec.Body.String()
	want := `{"id":"testid","name":"testname","created_at":"0001-01-01T00:00:00Z"}`
	if got != want {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}
