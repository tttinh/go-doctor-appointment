package app

import (
	"context"
)

type UserRepository interface {
	GetDoctorByUsername(ctx context.Context, username string) (Doctor, error)
	GetPatientByUsername(ctx context.Context, username string) (Patient, error)
	CreateDoctor(ctx context.Context, d Doctor) (Doctor, error)
	CreatePatient(ctx context.Context, p Patient) (Patient, error)
}

// type CalendarRepository interface {
// 	Get(ctx context.Context,
// 		doctorID int,
// 		from time.Time,
// 		to time.Time,
// 	) error

// 	AddAvailability(
// 		ctx context.Context,
// 		doctorID int,
// 		slots []Slot,
// 	) error
// }

// type AppointmentRepository interface {
// 	Book(ctx context.Context, a Appointment) error
// 	Reschedule(ctx context.Context, a Appointment, s Slot) error
// 	Cancel(ctx context.Context, a Appointment) error

// 	Swap(ctx context.Context, a1 Appointment, a2 Appointment) error
// }
