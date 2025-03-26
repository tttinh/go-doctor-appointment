package app

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

type CalendarRepository interface {
	CreateSlots(
		ctx context.Context,
		doctorID int,
		slots []time.Time,
	) error
}
