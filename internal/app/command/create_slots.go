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

type CreateSlotsHandler struct {
	slotRepo domain.SlotRepository
}

func (h CreateSlotsHandler) Handle(ctx context.Context, cmd CreateSlots) error {
	return h.slotRepo.CreateSlots(ctx, cmd.DoctorID, cmd.Slots)
}
