package mysql

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	models "github.com/whoiswentz/snippetbox/pkg"
	"golang.org/x/crypto/bcrypt"
	"strings"
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
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uniq_email") {
				return models.ErrDuplicatedEmail
			}
		}
		return err
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
