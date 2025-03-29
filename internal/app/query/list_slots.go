package query

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type ListSlots struct {
	DoctorID int
}

type ListSlotsHandler func(context.Context, ListSlots) ([]domain.Slot, error)

func NewListSlotsHandler(slotRepo domain.SlotRepository) ListSlotsHandler {
	return func(ctx context.Context, q ListSlots) ([]domain.Slot, error) {
		return slotRepo.ListSlots(ctx, q.DoctorID)
	}
}
