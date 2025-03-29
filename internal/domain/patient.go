package domain

import "time"

type Patient struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InvalidPatient() Patient {
	return Patient{}
}
