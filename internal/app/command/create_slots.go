package command

import (
	"context"
	"time"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type CreateSlots struct {
	DoctorID int
	Slots    []time.Time
}

type CreateSlotsHandler func(context.Context, CreateSlots) error

func NewCreateSlotsHandler(slotRepo domain.SlotRepository) CreateSlotsHandler {
	return func(ctx context.Context, cmd CreateSlots) error {
		return slotRepo.CreateSlots(ctx, cmd.DoctorID, cmd.Slots)
	}
}
