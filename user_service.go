package main

import (
	"github.com/google/uuid"
)

type UserService struct {
	Repo UserRepository
}

func (s *UserService) getUserByID(inputId uuid.UUID) (*GreetingData, error) {
	return s.Repo.getUser(inputId)
}

func (s *UserService) updateUserByID(inputId uuid.UUID) (*GreetingData, error) {
	return s.Repo.updateUser(inputId)
}

func (s *UserService) deleteUserByID(inputId uuid.UUID) (GreetingData, error) {
	return s.Repo.deleteUser(inputId)
}

func (s *UserService) userContainsByID(inputId uuid.UUID) (*GreetingData, error) {	
	return s.Repo.userContains(inputId)
}

