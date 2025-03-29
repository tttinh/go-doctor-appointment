package query

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginPatient struct {
	Username string
	Password string
}

type LoginPatientHandler struct {
	userRepo domain.UserRepository
}

func (h LoginPatientHandler) Handle(ctx context.Context, query LoginPatient) (domain.Patient, error) {
	p, err := h.userRepo.GetPatientByUsername(ctx, query.Username)
	if err != nil {
		return domain.InvalidPatient(), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(query.Password))
	if err != nil {
		return domain.InvalidPatient(), domain.ErrInvalidPassword
	}

	return p, nil
}
