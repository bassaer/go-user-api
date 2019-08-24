package app

import (
	"database/sql"
	"strconv"

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
	user := &User{}
	id, err := strconv.Atoi(reqID)
	if err != nil {
		return nil, err
	}
	err = u.stmtOut.QueryRow(id).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
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
