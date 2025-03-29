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
	CreateSlots     command.CreateSlotsHandler
	RegisterDoctor  command.RegisterDoctorHandler
	RegisterPatient command.RegisterPatientHandler
}

type Query struct {
	LoginDoctor  query.LoginDoctorHandler
	LoginPatient query.LoginPatientHandler
	ListSlots    query.ListSlotsHandler
}

func NewApplication(userRepo domain.UserRepository, slotRepo domain.SlotRepository) Application {
	return Application{
		Command{
			CreateSlots:     command.NewCreateSlotsHandler(slotRepo),
			RegisterDoctor:  command.NewRegisterDoctorHandler(userRepo),
			RegisterPatient: command.NewRegisterPatientHandler(userRepo),
		},
		Query{
			LoginDoctor:  query.NewLoginDoctorHandler(userRepo),
			LoginPatient: query.NewLoginPatientHandler(userRepo),
			ListSlots:    query.NewListSlotsHandler(slotRepo),
		},
	}
}
