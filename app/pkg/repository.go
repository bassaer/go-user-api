package app

import (
	"database/sql"
	"time"

	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

const auth = "root:test@tcp(db:3306)/userdb?parseTime=true&loc=Asia%2FTokyo"

// Repository is
type Repository interface {
	Get(id string) (*User, error)
	Set(user User) error
}

// UserRepository is
type UserRepository struct {
	db      *sql.DB
	stmtIns *sql.Stmt
	stmtOut *sql.Stmt
}

// NewUserRepository is
func NewUserRepository() (*UserRepository, error) {
	db, err := sql.Open("mysql", auth)
	if err != nil {
		return nil, err
	}

	ins, err := db.Prepare("INSERT INTO users VALUES(?, ?, NOW())")
	if err != nil {
		return nil, err
	}

	out, err := db.Prepare("SELECT id, name, created_at FROM users WHERE id = ?")
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
func (u *UserRepository) Get(reqID string) (*User, error) {
	var id string
	var name string
	var createdAt time.Time
	err := u.stmtOut.QueryRow(reqID).Scan(&id, &name, &createdAt)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}, nil
}

// Set is
func (u *UserRepository) Set(user User) error {
	return nil
}

// Close is
func (u *UserRepository) Close() {
	u.db.Close()
	u.stmtIns.Close()
	u.stmtOut.Close()
}
