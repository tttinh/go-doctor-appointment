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

var invalidDoctor = Doctor{}
var invalidPatient = Patient{}

type UserService struct {
	repo UserRepository
}

func (s UserService) SignupDoctor(ctx context.Context, d Doctor) (Doctor, error) {
	d, err := s.repo.CreateDoctor(ctx, d)
	if err != nil {
		return invalidDoctor, nil
	}
	return d, nil
}

func (s UserService) SignupPatient(ctx context.Context, p Patient) (Patient, error) {
	p, err := s.repo.CreatePatient(ctx, p)
	if err != nil {
		return invalidPatient, nil
	}
	return p, nil
}

func (s UserService) SigninDoctor(
	ctx context.Context,
	username string,
	password string,
) (Doctor, error) {
	d, err := s.repo.GetDoctorByUsername(ctx, username)
	if err != nil {
		return invalidDoctor, err
	}

	if password != d.Password {
		return invalidDoctor, ErrInvalidPassword
	}

	return d, nil
}

func (s UserService) SigninPatient(
	ctx context.Context,
	username string,
	password string,
) (Patient, error) {
	p, err := s.repo.GetPatientByUsername(ctx, username)
	if err != nil {
		return invalidPatient, err
	}

	if password != p.Password {
		return invalidPatient, ErrInvalidPassword
	}

	return p, nil
}
