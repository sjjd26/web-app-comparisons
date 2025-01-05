package models

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/big"
	"time"
)

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Salt         string `json:"salt"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func generateSalt() (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 16)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", errors.New("random number generation failed: " + err.Error())
		}
		b[i] = letters[n.Int64()]
	}
	return string(b), nil
}

func createPasswordHash(password string, salt string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password + salt))
	if err != nil {
		return "", errors.New("failed to hash password: " + err.Error())
	}
	return string(hash.Sum(nil)), nil
}

func NewUser(email string, password string) (User, error) {
	salt, err := generateSalt()
	if err != nil {
		return User{}, err
	}
	passwordHash, err := createPasswordHash(password, salt)
	if err != nil {
		return User{}, err
	}

	newUser := User{
		Email:        email,
		PasswordHash: passwordHash,
		Salt:         salt,
		CreatedAt:    time.Now().Format(time.RFC3339),
		UpdatedAt:    time.Now().Format(time.RFC3339),
	}

	return newUser, nil
}

func (user User) GetPasswordHash() string {
	// TODO: Implement this function
	return ""
}
