package app

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type normalRepository struct{}

func (n *normalRepository) Get(id string) (*User, error) {
	return &User{
		ID:        id,
		Name:      "testuser",
		CreatedAt: time.Date(2019, 8, 25, 14, 20, 0, 0, time.UTC),
	}, nil

}

func (n *normalRepository) Set(user *User) error {
	user.ID = "a8133582-a700-4ebf-9df2-cd1ad7c44547"
	user.CreatedAt = time.Date(2019, 9, 25, 14, 20, 0, 0, time.UTC)
	return nil
}

type errRepository struct{}

func (e *errRepository) Get(id string) (*User, error) {
	return nil, errors.New("DB error")
}

func (e *errRepository) Set(user *User) error {
	return errors.New("DB error")
}

var testcases = []struct {
	repo   Repository
	method string
	query  string
	body   io.Reader
	code   int
	resp   string
}{
	{
		repo:   &normalRepository{},
		method: http.MethodGet,
		query:  "/?id=7a687237-77d8-45b7-a8c7-1a7c0d719c8f",
		body:   nil,
		code:   http.StatusOK,
		resp:   `{"id":"7a687237-77d8-45b7-a8c7-1a7c0d719c8f","name":"testuser","created_at":"2019-08-25T14:20:00Z"}`,
	},
	{
		repo:   &normalRepository{},
		method: http.MethodPost,
		query:  "/",
		body:   strings.NewReader(`{"name":"newuser"}`),
		code:   http.StatusOK,
		resp:   `{"id":"a8133582-a700-4ebf-9df2-cd1ad7c44547","name":"newuser","created_at":"2019-09-25T14:20:00Z"}`,
	},
	{
		repo:   &errRepository{},
		method: http.MethodGet,
		query:  "/?id=7a687237-77d8-45b7-a8c7-1a7c0d719c8f",
		body:   strings.NewReader(`{"name":"newuser"}`),
		code:   http.StatusMethodNotAllowed,
		resp:   "",
	},
	{
		repo:   &errRepository{},
		method: http.MethodPost,
		query:  "/",
		body:   strings.NewReader(`{"name":"newuser"}`),
		code:   http.StatusMethodNotAllowed,
		resp:   "",
	},
}

func TestHandler(t *testing.T) {
	for _, tc := range testcases {
		req := httptest.NewRequest(tc.method, tc.query, tc.body)
		rec := httptest.NewRecorder()

		handler := NewHandler(tc.repo)
		handler.ServeHTTP(rec, req)

		if tc.code != rec.Code {
			t.Errorf("Status code should be %d, but %d\n", tc.code, rec.Code)
		}

		got := rec.Body.String()
		if tc.resp != got {
			t.Errorf("want: %#v, got: %#v", tc.resp, got)
		}
	}
}
