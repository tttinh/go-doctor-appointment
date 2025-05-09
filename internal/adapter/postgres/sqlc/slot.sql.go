// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: slot.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const fetchSlotByID = `-- name: FetchSlotByID :one
SELECT
  id,
  doctor_id,
  start_time,
  available
FROM
  slot
WHERE
  id = $1
`

type FetchSlotByIDRow struct {
	ID        int32
	DoctorID  pgtype.Int4
	StartTime pgtype.Timestamp
	Available pgtype.Bool
}

func (q *Queries) FetchSlotByID(ctx context.Context, id int32) (FetchSlotByIDRow, error) {
	row := q.db.QueryRow(ctx, fetchSlotByID, id)
	var i FetchSlotByIDRow
	err := row.Scan(
		&i.ID,
		&i.DoctorID,
		&i.StartTime,
		&i.Available,
	)
	return i, err
}

const fetchSlots = `-- name: FetchSlots :many
SELECT
  id,
  doctor_id,
  start_time,
  available
FROM
  slot
`

type FetchSlotsRow struct {
	ID        int32
	DoctorID  pgtype.Int4
	StartTime pgtype.Timestamp
	Available pgtype.Bool
}

func (q *Queries) FetchSlots(ctx context.Context) ([]FetchSlotsRow, error) {
	rows, err := q.db.Query(ctx, fetchSlots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchSlotsRow
	for rows.Next() {
		var i FetchSlotsRow
		if err := rows.Scan(
			&i.ID,
			&i.DoctorID,
			&i.StartTime,
			&i.Available,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type InsertSlotsParams struct {
	DoctorID  pgtype.Int4
	StartTime pgtype.Timestamp
}

const updateSlotAvailability = `-- name: UpdateSlotAvailability :exec
UPDATE slot
SET
  available = $2
WHERE
  id = $1
`

type UpdateSlotAvailabilityParams struct {
	ID        int32
	Available pgtype.Bool
}

func (q *Queries) UpdateSlotAvailability(ctx context.Context, arg UpdateSlotAvailabilityParams) error {
	_, err := q.db.Exec(ctx, updateSlotAvailability, arg.ID, arg.Available)
	return err
}

const updateSlotAvailable = `-- name: UpdateSlotAvailable :exec
UPDATE slot
SET
  available = $2
WHERE
  id = $1
`

type UpdateSlotAvailableParams struct {
	ID        int32
	Available pgtype.Bool
}

func (q *Queries) UpdateSlotAvailable(ctx context.Context, arg UpdateSlotAvailableParams) error {
	_, err := q.db.Exec(ctx, updateSlotAvailable, arg.ID, arg.Available)
	return err
}
