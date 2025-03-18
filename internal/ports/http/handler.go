package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler() http.Handler {
	h := Handler{}
	router := gin.Default()
	api := router.Group("/api")
	api.GET("", h.do)
	return router
}

type Handler struct{}

func (h *Handler) do(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
