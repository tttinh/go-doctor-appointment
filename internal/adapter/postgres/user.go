package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tinhtt/go-doctor-appointment/internal/adapter/postgres/sqlc"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type Users struct {
	*sqlc.Queries
	// db *pgxpool.Pool
}

func NewUsers(db *pgxpool.Pool) Users {
	return Users{
		Queries: sqlc.New(db),
	}
}

func (u Users) GetDoctorByUsername(ctx context.Context, username string) (domain.Doctor, error) {
	r, err := u.FetchDoctorByUsername(ctx, username)
	if err != nil {
		return domain.InvalidDoctor(), err
	}

	return domain.Doctor{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) GetPatientByUsername(ctx context.Context, username string) (domain.Patient, error) {
	r, err := u.FetchPatientByUsername(ctx, username)
	if err != nil {
		return domain.InvalidPatient(), err
	}

	return domain.Patient{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) CreateDoctor(ctx context.Context, d domain.Doctor) (domain.Doctor, error) {
	r, err := u.InsertDoctor(ctx, sqlc.InsertDoctorParams{
		Username:       d.Username,
		Email:          d.Email,
		HashedPassword: d.Password,
	})
	if err != nil {
		return domain.InvalidDoctor(), nil
	}

	return domain.Doctor{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) CreatePatient(ctx context.Context, p domain.Patient) (domain.Patient, error) {
	r, err := u.InsertPatient(ctx, sqlc.InsertPatientParams{
		Username:       p.Username,
		Email:          p.Email,
		HashedPassword: p.Password,
	})
	if err != nil {
		return domain.InvalidPatient(), nil
	}

	return domain.Patient{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}
