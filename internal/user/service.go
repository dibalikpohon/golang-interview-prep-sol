package user

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
  "golang.org/x/crypto/bcrypt"
)

var UsernameAlreadyExists = errors.New("Username already exists")

const HASH_COST = 18

type service struct {
	dbUser     string
	dbPassword string
}

func NewService(dbUser, dbPassword string) (*service, error) {
	if dbUser == "" {
		return nil, errors.New("dbUser was empty")
	}
	return &service{dbUser: dbUser, dbPassword: dbPassword}, nil
}

type User struct {
	Username     string
	Password string
}

func (u *User) GetPasswordHash() (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), HASH_COST);
  return string(bytes), err
}

func (s *service) AddUser(u User) (string, error) {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost/test_repo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

  var count int
  q := "SELECT COUNT(id) FROM users WHERE username=?"
  db.QueryRow(q, u.Username).Scan(&count)
  if count > 0 {
    return "", fmt.Errorf("failed to insert: %w", UsernameAlreadyExists)
  }

	var id string
  q = "INSERT INTO users (username, password) VALUES (?, ?) RETURNING id"

  hashedPassword, err := u.GetPasswordHash()
  if err != nil {
    return "", fmt.Errorf("failed to insert: %w", err)
  }
	err = db.QueryRow(q, u.Username, hashedPassword).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
