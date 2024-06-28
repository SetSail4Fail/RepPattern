package main

import "github.com/google/uuid"

type UserRepository interface {
	getUser(inputId uuid.UUID) (*GreetingData, error)
	updateUser(inputId uuid.UUID) (*GreetingData, error)
	deleteUser(inputId uuid.UUID) (GreetingData, error)
	userContains(inputId uuid.UUID) (*GreetingData, error)
}

