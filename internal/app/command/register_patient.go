package command

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPatient struct {
	Username string
	Email    string
	Password string
}

type RegisterPatientHandler func(context.Context, RegisterPatient) error

func NewRegisterPatientHandler(userRepo domain.UserRepository) RegisterPatientHandler {
	return func(ctx context.Context, cmd RegisterPatient) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		p := domain.Patient{
			Username: cmd.Username,
			Email:    cmd.Email,
			Password: string(hashedPassword),
		}

		_, err = userRepo.CreatePatient(ctx, p)
		if err != nil {
			return err
		}

		return nil
	}
}
