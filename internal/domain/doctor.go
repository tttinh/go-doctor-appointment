package domain

import "time"

type Doctor struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InvalidDoctor() Doctor {
	return Doctor{}
}
