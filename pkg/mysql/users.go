package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	models "github.com/whoiswentz/snippetbox/pkg"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserModel struct {
	mu sync.Mutex
	DB *sql.DB
}

func NewUserModel(DB *sql.DB) *UserModel {
	return &UserModel{DB: DB}
}

func (m *UserModel) EmailTaken(email string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	stmt := `SELECT email FROM users WHERE email = ?`
	r := m.DB.QueryRow(stmt, email)

	var e string
	if err := r.Scan(&e); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return true
	}

	return true
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`
	if _, err := m.DB.Exec(stmt, name, email, hashedPassword); err != nil {
		return err
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	stmt := "SELECT id, hashed_password FROM users WHERE email = ? AND active = TRUE"

	var userId int
	var hashedPassword []byte
	row := m.DB.QueryRow(stmt, email)
	if err := row.Scan(&userId, &hashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	return userId, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
