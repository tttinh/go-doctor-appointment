package ports

import (
	"net/http"

	httpport "github.com/tinhtt/go-doctor-appointment/internal/ports/http"
)

func NewHTTPServer() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: httpport.NewHandler(),
	}
}
