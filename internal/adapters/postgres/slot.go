package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tinhtt/go-doctor-appointment/internal/adapters/postgres/sqlc"
)

type Slots struct {
	*sqlc.Queries
	// db *pgxpool.Pool
}

func NewSlots(db *pgxpool.Pool) Slots {
	return Slots{
		Queries: sqlc.New(db),
	}
}

func (s Slots) CreateSlots(
	ctx context.Context,
	doctorID int,
	slots []time.Time,
) error {
	return nil
}
