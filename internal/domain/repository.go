package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	GetDoctorByUsername(ctx context.Context, username string) (Doctor, error)
	GetPatientByUsername(ctx context.Context, username string) (Patient, error)
	CreateDoctor(ctx context.Context, d Doctor) (Doctor, error)
	CreatePatient(ctx context.Context, p Patient) (Patient, error)
}

type SlotRepository interface {
	CreateSlots(
		ctx context.Context,
		doctorID int,
		slots []time.Time,
	) error

	ListSlots(
		ctx context.Context,
		doctorID int,
	) ([]Slot, error)
}
