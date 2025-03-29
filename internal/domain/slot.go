package domain

import "time"

// type SlotStatus struct {
// 	value string
// }

// var (
// 	SlotAvailable   SlotStatus = SlotStatus{"available"}
// 	SlotUnavailable SlotStatus = SlotStatus{"unavailable"}
// 	SlotBooked      SlotStatus = SlotStatus{"booked"}
// )

type Slot struct {
	ID        int
	DotorID   int
	Hour      time.Time
	Available bool
	// Status    SlotStatus
}
