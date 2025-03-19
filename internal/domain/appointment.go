package domain

type Appointment struct {
	Base
	SlotID    int
	DoctorID  int
	PatientID int
	Note      string
	Status    string // booked, cancelled, rescheduled
}
