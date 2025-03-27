package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tinhtt/go-doctor-appointment/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
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

func (u Users) GetDoctorByUsername(ctx context.Context, username string) (app.Doctor, error) {
	r, err := u.FetchDoctorByUsername(ctx, username)
	if err != nil {
		return app.InvalidDoctor, err
	}

	return app.Doctor{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) GetPatientByUsername(ctx context.Context, username string) (app.Patient, error) {
	r, err := u.FetchPatientByUsername(ctx, username)
	if err != nil {
		return app.InvalidPatient, err
	}

	return app.Patient{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) CreateDoctor(ctx context.Context, d app.Doctor) (app.Doctor, error) {
	r, err := u.InsertDoctor(ctx, sqlc.InsertDoctorParams{
		Username:       d.Username,
		Email:          d.Email,
		HashedPassword: d.Password,
	})
	if err != nil {
		return app.InvalidDoctor, nil
	}

	return app.Doctor{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}

func (u Users) CreatePatient(ctx context.Context, p app.Patient) (app.Patient, error) {
	r, err := u.InsertPatient(ctx, sqlc.InsertPatientParams{
		Username:       p.Username,
		Email:          p.Email,
		HashedPassword: p.Password,
	})
	if err != nil {
		return app.InvalidPatient, nil
	}

	return app.Patient{
		ID:        int(r.ID),
		Username:  r.Username,
		Password:  r.HashedPassword,
		Email:     r.Email,
		CreatedAt: r.CreatedAt.Time,
		UpdatedAt: r.UpdatedAt.Time,
	}, nil
}
