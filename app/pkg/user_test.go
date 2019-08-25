package app

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestUserMarshal(t *testing.T) {
	user := &User{
		ID:        "7a687237-77d8-45b7-a8c7-1a7c0d719c8f",
		Name:      "testuser",
		CreatedAt: time.Date(2019, 8, 25, 14, 20, 0, 0, time.UTC),
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	got := string(bytes)
	want := `{"id":"7a687237-77d8-45b7-a8c7-1a7c0d719c8f","name":"testuser","created_at":"2019-08-25T14:20:00Z"}`
	if got != want {
		t.Errorf("\nwant: %#v\ngot : %#v\n", want, got)
	}
}

func TestUserUnmarshal(t *testing.T) {
	want := &User{
		ID:        "7a687237-77d8-45b7-a8c7-1a7c0d719c8f",
		Name:      "testuser",
		CreatedAt: time.Date(2019, 8, 25, 14, 20, 0, 0, time.UTC),
	}
	got := &User{}
	bytes := []byte(`{"id":"7a687237-77d8-45b7-a8c7-1a7c0d719c8f","name":"testuser","created_at":"2019-08-25T14:20:00Z"}`)
	json.Unmarshal(bytes, got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("\nwant: %s\ngot : %s\n", want, got)
	}
}
