package main

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	getUser(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error)
	updateUser(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error)
	deleteUser(db *sqlx.DB, inputId uuid.UUID) (GreetingData, error)
	userContains(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error)
}

type Service struct {
	UserRepository UserRepository
}
