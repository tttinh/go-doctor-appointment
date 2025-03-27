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

type SlotService struct {
	repo SlotRepository
}

func NewSlotService(repo SlotRepository) SlotService {
	return SlotService{repo: repo}
}

func (s SlotService) AddSlots(ctx context.Context, doctorID int, slots []time.Time) error {
	return s.repo.CreateSlots(ctx, doctorID, slots)
}

func (s SlotService) ChangeAvailability(ctx context.Context, doctorID int, slotID int, val bool) error {
	return nil
}
