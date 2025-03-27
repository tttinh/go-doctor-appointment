package app

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Doctor struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Patient struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var InvalidDoctor = Doctor{}
var InvalidPatient = Patient{}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{repo: repo}
}

func (s UserService) SignupDoctor(ctx context.Context, d Doctor) (Doctor, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return InvalidDoctor, err
	}
	d.Password = string(hashedPassword)

	d, err = s.repo.CreateDoctor(ctx, d)
	if err != nil {
		return InvalidDoctor, nil
	}
	return d, nil
}

func (s UserService) SignupPatient(ctx context.Context, p Patient) (Patient, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return InvalidPatient, err
	}
	p.Password = string(hashedPassword)

	p, err = s.repo.CreatePatient(ctx, p)
	if err != nil {
		return InvalidPatient, nil
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
		return InvalidDoctor, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password))
	if err != nil {
		return InvalidDoctor, ErrInvalidPassword
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
		return InvalidPatient, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return InvalidPatient, ErrInvalidPassword
	}

	return p, nil
}
