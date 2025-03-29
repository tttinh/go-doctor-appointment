package httpport

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
	"github.com/tinhtt/go-doctor-appointment/internal/common/auth"
)

type handler struct {
	jwt auth.JWT
	app app.Application
}

func NewHandler(
	l *slog.Logger,
	app app.Application,
) http.Handler {
	jwt := auth.NewJWT("abc", 24*time.Hour)
	h := handler{
		jwt: jwt,
		app: app,
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logMiddleware(l), gin.Recovery())
	public := router.Group("/api")
	{
		public.GET("/calendar/:id", h.getCalendar)

		// Auth
		public.POST("/doctor/signup", h.signupDoctor)
		public.POST("/doctor/signin", h.signinDoctor)
		public.POST("/patient/signup", h.signupPatient)
		public.POST("/patient/signin", h.signinPatient)

		public.GET("/testing/slots", h.generateSlots)
	}

	private := public.Use(authMiddleware(jwt))
	{
		// Slot
		private.POST("/slot", h.addSlots)
		private.GET("/slot", h.listSlots)
		private.PATCH("/slot/:id/availability", h.changeSlotAvailability)

		// Appointment
		private.GET("/appointment", h.listAppointments)
		private.POST("/appointment", h.createApointment)
		private.DELETE("/appointment/:id", h.removeAppointment)
	}

	return router
}
