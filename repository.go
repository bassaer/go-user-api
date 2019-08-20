package main

import "time"

// Repository is
type Repository interface {
	Get(id string) (*User, error)
	Set(user User) error
}

// UserRepository is
type UserRepository struct{}

// Get is
func (u *UserRepository) Get(id string) (*User, error) {
	return &User{
		ID:        "testid",
		Name:      "testname",
		CreatedAt: time.Now(),
	}, nil
}

// Set is
func (u *UserRepository) Set(user User) error {
	return nil
}
