package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
)

func NewHandler(u app.UserService, s app.SlotService) http.Handler {
	h := handler{
		user: u,
		slot: s,
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	api := router.Group("/api")

	api.GET("/calendar/:id", h.getCalendar)

	// Auth
	api.POST("/doctor/signup", h.signupDoctor)
	api.POST("/doctor/signin", h.signinDoctor)
	api.POST("/patient/signup", h.signupPatient)
	api.POST("/patient/signin", h.signinPatient)

	// Slot
	api.POST("/slot", h.addSlots)
	api.PUT("/slot/:id/availability", h.updateSlotAvailability)

	// Appointment
	api.GET("/appointment", h.listAppointments)
	api.POST("/appointment", h.createApointment)
	api.DELETE("/appointment/:id", h.removeAppointment)
	return router
}

type handler struct {
	user app.UserService
	slot app.SlotService
}
