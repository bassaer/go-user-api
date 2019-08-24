package app

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type mockRepository struct{}

func (m *mockRepository) Get(reqID string) (*User, error) {
	id, _ := strconv.Atoi(reqID)
	return &User{
		ID:        id,
		Name:      "testname",
		CreatedAt: time.Time{},
	}, nil

}

func (m *mockRepository) Set(user User) error {
	return nil
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/?id=100", nil)
	rec := httptest.NewRecorder()

	handler := NewHandler(&mockRepository{})
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code should be 200, but %d\n", rec.Code)
	}

	got := rec.Body.String()
	want := `{"id":100,"name":"testname","created_at":"0001-01-01T00:00:00Z"}`
	if got != want {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}
