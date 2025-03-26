package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler() http.Handler {
	h := handler{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	api := router.Group("/api")

	// Public resources.
	api.GET("/calendar/:id", h.getCalendar)
	api.POST("/doctor/signup", h.signupDoctor)
	api.POST("/doctor/signin", h.signinDoctor)
	api.POST("/patient/signup", h.signupPatient)
	api.POST("/patient/signin", h.signinPatient)

	// Private resources.
	api.POST("/slot", h.addSlots)
	api.PUT("/slot/:id/availability", h.changeSlotAvailability)
	api.PUT("/slot/:id/book", h.bookApointment)
	api.PUT("/slot/:id/cancellation", h.cancelAppointment)

	api.GET("/appointment", h.listAppointments)
	return router
}

type handler struct {
}
