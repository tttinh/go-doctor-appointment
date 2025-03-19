package domain

import "time"

type Slot struct {
	Base
	Hour   time.Time
	Status string // available, unavailable, booked, cancelled,
	Note   string
}
