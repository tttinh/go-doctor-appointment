package ports

import (
	"net/http"
	"time"

	"github.com/tinhtt/go-doctor-appointment/internal/app"
	httpport "github.com/tinhtt/go-doctor-appointment/internal/ports/http"
)

func NewHTTPServer(u app.UserService, s app.SlotService) *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        httpport.NewHandler(u, s),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
