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
	// CreateSlots creates multiple slots for a given doctor by their ID.
	CreateSlots(
		ctx context.Context,
		doctorID int,
		slots []time.Time,
	) error

	// ListSlots retrieves all available slots for a given doctor by their ID.
	ListSlots(
		ctx context.Context,
		doctorID int,
	) ([]Slot, error)

	// GetSlotByID retrieves a slot by its ID.
	GetSlotByID(
		ctx context.Context,
		slotID int,
	) (Slot, error)

	// ChangeSlotAvailability updates the availability status of a slot.
	ChangeSlotAvailability(
		ctx context.Context,
		slotID int,
		available bool,
	) error
}

type Repository struct {
	User UserRepository
	Slot SlotRepository
}
