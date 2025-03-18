package domain

import "time"

type Base struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Slot struct {
	Base
	Hour   time.Time
	Status string // available, unavailable, booked, cancelled,
	Note   string
}

type Appointment struct {
	Base
	SlotID    int
	DoctorID  int
	PatientID int
	Note      string
	Status    string // created, cancelled, rescheduled
}

type Review struct {
	Base
	PatientID int
	Content   string
}

type Doctor struct {
	Base
	FirstName string
	LastName  string
	Bio       string
	Specs     []string
}

type Patient struct {
	Base
	FirstName string
	LastName  string
	Dob       string
	Gender    bool
}
