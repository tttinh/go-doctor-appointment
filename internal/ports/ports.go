package ports

import (
	"net/http"
	"time"

	httpport "github.com/tinhtt/go-doctor-appointment/internal/ports/http"
)

func NewHTTPServer() *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        httpport.NewHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
