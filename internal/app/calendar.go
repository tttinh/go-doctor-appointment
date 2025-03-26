package app

import (
	"context"
	"time"
)

// type SlotStatus struct {
// 	value string
// }

// var (
// 	SlotAvailable   SlotStatus = SlotStatus{"available"}
// 	SlotUnavailable SlotStatus = SlotStatus{"unavailable"}
// 	SlotBooked      SlotStatus = SlotStatus{"booked"}
// )

// type Slot struct {
// 	ID        int
// 	DotorID   int
// 	Hour      time.Time
// 	Status    SlotStatus
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type CalendarService struct {
	repo CalendarRepository
}

func (s CalendarService) Get(ctx context.Context, doctorID int) error {
	return nil
}

func (s CalendarService) AddSlots(ctx context.Context, doctorID int, slots []time.Time) error {
	return s.repo.CreateSlots(ctx, doctorID, slots)
}

func (s CalendarService) ChangeAvailability(ctx context.Context, doctorID int, slotID int, val bool) error {
	return nil
}

func (s CalendarService) GetDoctorAppointments(ctx context.Context, doctorID int) error {
	return nil
}

func (s CalendarService) GetPatientAppointments(ctx context.Context, doctorID int) error {
	return nil
}

func (s CalendarService) Book(ctx context.Context, patientID int, slotID int) error {
	return nil
}

func (s CalendarService) Cancel(ctx context.Context, patientID int, slotID int) error {
	return nil
}
