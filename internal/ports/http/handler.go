package httpport

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
	"github.com/tinhtt/go-doctor-appointment/internal/common/auth"
)

func NewHandler(
	l *slog.Logger,
	u app.UserService,
	s app.SlotService,
) http.Handler {
	jwt := auth.NewJWT("abc", 24*time.Hour)
	h := handler{
		jwt:  jwt,
		user: u,
		slot: s,
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
	}

	private := public.Use(authMiddleware(jwt))
	{
		// Slot
		private.POST("/slot", h.addSlots)
		private.PUT("/slot/:id/availability", h.updateSlotAvailability)

		// Appointment
		private.GET("/appointment", h.listAppointments)
		private.POST("/appointment", h.createApointment)
		private.DELETE("/appointment/:id", h.removeAppointment)
	}

	return router
}

type handler struct {
	jwt  auth.JWT
	user app.UserService
	slot app.SlotService
}
