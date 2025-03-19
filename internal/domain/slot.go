package domain

import "time"

// Slot represents an available time slot of a doctor.
type Slot struct {
	Base
	Available bool
	Hour      time.Time
}
