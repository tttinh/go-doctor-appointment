package domain

type Review struct {
	Base
	PatientID int
	Content   string
}
