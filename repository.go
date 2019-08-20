package main

import "time"

type Repository interface {
	Get(id string) (*User, error)
	Set(user User) error
}

type UserRepository struct{}

func (u *UserRepository) Get(id string) (*User, error) {
	return &User{
		Id:        "testid",
		Name:      "testname",
		CreatedAt: time.Now(),
	}, nil
}

func (u *UserRepository) Set(user User) error {
	return nil
}
