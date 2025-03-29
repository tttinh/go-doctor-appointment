package command

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterDoctor struct {
	Username string
	Email    string
	Password string
}

type RegisterDoctorHandler func(context.Context, RegisterDoctor) error

func NewRegisterDoctorHandler(userRepo domain.UserRepository) RegisterDoctorHandler {
	return func(ctx context.Context, cmd RegisterDoctor) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		d := domain.Doctor{
			Username: cmd.Username,
			Email:    cmd.Email,
			Password: string(hashedPassword),
		}

		_, err = userRepo.CreateDoctor(ctx, d)
		if err != nil {
			return err
		}

		return nil
	}
}
