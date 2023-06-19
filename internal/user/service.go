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
  DB *sql.DB
}

func NewService(dbUser, dbPassword string) (*service, error) {
  dsn := fmt.Sprintf("postgres://%s:%s@localhost/test_repo?sslmode=disable", dbUser, dbPassword)
  db, err := sql.Open("postgres", dsn)
  if err != nil {
    panic(err)
  }
  return &service{DB: db}, nil
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
  var count int
  q := "SELECT COUNT(id) FROM users WHERE username=$1"
  if count > 0 {
    return "", fmt.Errorf("failed to insert: %w", UsernameAlreadyExists)
  }

	var id string

  q = "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"


  hashedPassword, err := u.GetPasswordHash()
  if err != nil {
    return "", fmt.Errorf("failed to insert: %w", err)
  }
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
