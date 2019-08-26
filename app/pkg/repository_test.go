package app

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	//mockID := &MockID{}
	user := &User{
		ID:        "7a687237-77d8-45b7-a8c7-1a7c0d719c8f",
		Name:      "testuser",
		CreatedAt: time.Date(2019, 8, 25, 14, 20, 0, 0, time.UTC),
	}

	anyArg := sqlmock.AnyArg()
	rows := sqlmock.NewRows([]string{"id", "name", "created_at"}).AddRow(user.ID, user.Name, user.CreatedAt)
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users VALUES (?, ?, ?)"))
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, name, created_at FROM users WHERE id = ?"))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users")).WithArgs(anyArg, user.Name, anyArg).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, created_at FROM users WHERE id = ?")).WillReturnRows(rows)

	repo, err := NewUserRepository(db)
	if err != nil {
		t.Error(err)
	}
	if err = repo.Set(user); err != nil {
		t.Error(err)
	}
	if _, err = repo.Get(user.ID); err != nil {
		t.Error(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
