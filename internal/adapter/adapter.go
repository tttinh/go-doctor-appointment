package adapter

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tinhtt/go-doctor-appointment/internal/adapter/postgres"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dbURL = "postgres://xyz:xyz@localhost:5432/appointment?sslmode=disable"

func NewPostgresConnection() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dbURL)
}

func NewRepositoryWithPostgres(db *pgxpool.Pool) domain.Repository {
	return domain.Repository{
		User: postgres.NewUsers(db),
		Slot: postgres.NewSlots(db),
	}
}

func Migrate() error {
	m, err := migrate.New(
		"file://db/postgres/migration",
		dbURL,
	)

	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
