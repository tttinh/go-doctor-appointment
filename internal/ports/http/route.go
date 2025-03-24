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
	api.GET("/doctor/:id/calendar", h.todo)

	// Private resources.
	// Doctors operations
	api.POST("/doctor/signup", h.todo)
	api.POST("/doctor/signin", h.todo)
	api.POST("/doctor/slot", h.todo)
	api.PUT("/doctor/slot/:id", h.todo)
	api.GET("/doctor/appointment", h.todo)

	// Patients operations
	api.POST("/patient/signup", h.todo)
	api.POST("/patient/signin", h.todo)
	api.GET("/patient/appointment", h.todo)
	api.POST("/patient/appointment", h.todo)

	return router
}

type handler struct {
}

func (h *handler) todo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO"})
}
