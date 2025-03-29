package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tinhtt/go-doctor-appointment/internal/adapter/postgres/sqlc"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
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

func (s Slots) CreateSlots(ctx context.Context, doctorID int, slots []time.Time) error {
	params := []sqlc.InsertSlotsParams{}
	for _, s := range slots {
		params = append(params, sqlc.InsertSlotsParams{
			DoctorID:  pgtype.Int4{Int32: int32(doctorID), Valid: true},
			StartTime: pgtype.Timestamp{Time: s, Valid: true},
		})
	}
	_, err := s.InsertSlots(ctx, params)
	return err
}

func (s Slots) ListSlots(ctx context.Context, doctorID int) ([]domain.Slot, error) {
	slots := []domain.Slot{}
	rows, err := s.FetchSlots(ctx)
	if err != nil {
		return slots, err
	}

	for _, r := range rows {
		slots = append(slots, domain.Slot{
			ID:        int(r.ID),
			DotorID:   int(r.DoctorID.Int32),
			Hour:      r.StartTime.Time,
			Available: r.Available.Bool,
		})
	}

	return slots, nil
}
