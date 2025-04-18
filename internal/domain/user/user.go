package domain

import (
	"AvitoFlats/internal/utils"
	"time"
)

type User struct {
	Id               string
	Email            string
	Username         string
	Password         string
	TimeRegistration time.Time
	TimeUpdatedAt    time.Time
}

func NewUser(email string, username string, password string) (*User, error) {
	return &User{
		Id:               utils.GenerateUUID(),
		Email:            email,
		Username:         username,
		Password:         utils.HashPassword(password),
		TimeRegistration: time.Now(),
		TimeUpdatedAt:    time.Now(),
	}, nil
}
