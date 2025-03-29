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

type LoginDoctorHandler struct {
	userRepo domain.UserRepository
}

func (h LoginDoctorHandler) Handle(ctx context.Context, query LoginDoctor) (domain.Doctor, error) {
	d, err := h.userRepo.GetDoctorByUsername(ctx, query.Username)
	if err != nil {
		return domain.InvalidDoctor(), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(query.Password))
	if err != nil {
		return domain.InvalidDoctor(), domain.ErrInvalidPassword
	}

	return d, nil
}
