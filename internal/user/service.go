package user

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
  "golang.org/x/crypto/bcrypt"
)

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
	Name     string
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

	var id string
  q := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"

  hashedPassword, err := u.GetPasswordHash()
  if err != nil {
    return "", fmt.Errorf("failed to insert: %w", err)
  }
	err = db.QueryRow(q, u.Name, hashedPassword).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
