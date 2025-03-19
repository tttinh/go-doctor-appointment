package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

func NewHandler() http.Handler {
	h := Handler{}
	router := gin.Default()
	api := router.Group("/api")
	api.GET("", h.do)
	return router
}

type Handler struct {
	appointments domain.AppointmentRepository
	calendars    domain.CalendarRepository
	doctors      domain.DoctorRepository
}

func (h *Handler) do(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
