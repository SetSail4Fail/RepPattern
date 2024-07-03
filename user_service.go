package main

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (s *Service) getUserByID(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error) {
	return s.UserRepository.getUser(db, inputId)
}

func (s *Service) updateUserByID(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error) {
	return s.UserRepository.updateUser(db, inputId)
}

func (s *Service) deleteUserByID(db *sqlx.DB, inputId uuid.UUID) (GreetingData, error) {
	return s.UserRepository.deleteUser(db, inputId)
}

func (s *Service) userContainsByID(db *sqlx.DB, inputId uuid.UUID) (*GreetingData, error) {
	return s.UserRepository.userContains(db, inputId)
}