package user

import "golang.org/x/crypto/bcrypt"

const HASH_COST = 18

type User struct {
	Username     string
	Password string
}

func (u *User) GetPasswordHash() (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), HASH_COST);
  return string(bytes), err
}
