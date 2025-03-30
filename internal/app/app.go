package app

import (
	"github.com/tinhtt/go-doctor-appointment/internal/app/command"
	"github.com/tinhtt/go-doctor-appointment/internal/app/query"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type Application struct {
	Command Command
	Query   Query
}

type Command struct {
	CreateSlots            command.CreateSlotsHandler
	ChangeSlotAvailability command.ChangeSlotAvailabilityHandler
	RegisterDoctor         command.RegisterDoctorHandler
	RegisterPatient        command.RegisterPatientHandler
}

type Query struct {
	LoginDoctor  query.LoginDoctorHandler
	LoginPatient query.LoginPatientHandler
	ListSlots    query.ListSlotsHandler
}

func NewApplication(repo domain.Repository) Application {
	return Application{
		Command{
			CreateSlots:            command.NewCreateSlotsHandler(repo.Slot),
			ChangeSlotAvailability: command.NewChangeSlotAvailabilityHandler(repo.Slot),
			RegisterDoctor:         command.NewRegisterDoctorHandler(repo.User),
			RegisterPatient:        command.NewRegisterPatientHandler(repo.User),
		},
		Query{
			LoginDoctor:  query.NewLoginDoctorHandler(repo.User),
			LoginPatient: query.NewLoginPatientHandler(repo.User),
			ListSlots:    query.NewListSlotsHandler(repo.Slot),
		},
	}
}
