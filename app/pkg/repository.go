package app

import (
	"database/sql"

	"github.com/google/uuid"
)

const (
	queryIns = "INSERT INTO users VALUES (?, ?, ?)"
	queryOut = "SELECT id, name, created_at FROM users WHERE id = ?"
)

// Repository is
type Repository interface {
	Get(id string) (*User, error)
	Set(user *User) error
}

// UserRepository is
type UserRepository struct {
	db      *sql.DB
	stmtIns *sql.Stmt
	stmtOut *sql.Stmt
}

// NewUserRepository is
func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	ins, err := db.Prepare(queryIns)
	if err != nil {
		return nil, err
	}

	out, err := db.Prepare(queryOut)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db:      db,
		stmtIns: ins,
		stmtOut: out,
	}, nil
}

// Get is
func (u *UserRepository) Get(id string) (*User, error) {
	user := &User{}
	if err := u.stmtOut.QueryRow(id).Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

// Set is
func (u *UserRepository) Set(user *User) error {
	// get version 4 UUID
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	//now := time.Now()
	if _, err := u.stmtIns.Exec(id, user.Name, user.CreatedAt); err != nil {
		return err
	}
	user.ID = id.String()
	//user.CreatedAt = now
	return nil
}

// Close is
func (u *UserRepository) Close() {
	u.db.Close()
	u.stmtIns.Close()
	u.stmtOut.Close()
}
