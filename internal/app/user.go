package app

import (
	"context"
	"time"
)

type Doctor struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Patient struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService struct {
	repo UserRepository
}

func (s UserService) SignupDoctor(ctx context.Context, d Doctor) (Doctor, error) {
	return Doctor{}, nil
}

func (s UserService) SignupPatient(ctx context.Context, p Patient) (Patient, error) {
	return Patient{}, nil
}

func (s UserService) SigninDoctor(
	ctx context.Context,
	username string,
	password string,
) (Doctor, error) {
	return Doctor{}, nil
}

func (s UserService) SigninPatient(
	ctx context.Context,
	username string,
	password string,
) (Patient, error) {
	return Patient{}, nil
}
