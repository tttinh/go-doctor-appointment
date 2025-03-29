package query

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginDoctor struct {
	Username string
	Password string
}

type LoginDoctorHandler func(context.Context, LoginDoctor) (domain.Doctor, error)

func NewLoginDoctorHandler(userRepo domain.UserRepository) LoginDoctorHandler {
	return func(ctx context.Context, query LoginDoctor) (domain.Doctor, error) {
		d, err := userRepo.GetDoctorByUsername(ctx, query.Username)
		if err != nil {
			return domain.InvalidDoctor(), err
		}

		err = bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(query.Password))
		if err != nil {
			return domain.InvalidDoctor(), domain.ErrInvalidPassword
		}

		return d, nil
	}
}
