package command

import (
	"context"

	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type ChangeSlotAvailability struct {
	DoctorID  int
	SlotID    int
	Available bool
}

type ChangeSlotAvailabilityHandler func(context.Context, ChangeSlotAvailability) error

func NewChangeSlotAvailabilityHandler(slotRepo domain.SlotRepository) ChangeSlotAvailabilityHandler {
	return func(ctx context.Context, cmd ChangeSlotAvailability) error {
		// Check if the slot exists.
		slot, err := slotRepo.GetSlotByID(ctx, cmd.SlotID)
		if err != nil {
			return err
		}

		// Check if the slot belongs to the doctor.
		if slot.DoctorID != cmd.DoctorID {
			return domain.ErrAccessDenied
		}

		// Update the slot availability.
		return slotRepo.ChangeSlotAvailability(ctx, cmd.SlotID, cmd.Available)
	}
}
