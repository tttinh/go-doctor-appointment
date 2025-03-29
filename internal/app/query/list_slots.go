package query

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type ListSlots struct {
	DoctorID int
}

type ListSlotsHandler struct {
	slotRepo domain.SlotRepository
}

func (h ListSlotsHandler) Handle(ctx context.Context, q ListSlots) ([]domain.Slot, error) {
	return h.slotRepo.ListSlots(ctx, q.DoctorID)
}
