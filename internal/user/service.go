package user

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

var UsernameAlreadyExists = errors.New("Username already exists")


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


func (s *service) AddUser(u User) (string, error) {
  var count int
  q := "SELECT COUNT(id) FROM users WHERE username=$1"
  s.DB.QueryRow(q, u.Username).Scan(&count)
  if count > 0 {
    return "", fmt.Errorf("failed to insert: %w", UsernameAlreadyExists)
  }

	var id string

  q = "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"


  hashedPassword, err := u.GetPasswordHash()
  if err != nil {
    return "", fmt.Errorf("failed to insert: %w", err)
  }
	err = s.DB.QueryRow(q, u.Username, hashedPassword).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
