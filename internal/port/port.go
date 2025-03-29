package port

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/tinhtt/go-doctor-appointment/internal/app"
	httpport "github.com/tinhtt/go-doctor-appointment/internal/port/http"
)

func NewHTTPServer(l *slog.Logger, a app.Application) *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        httpport.NewHandler(l, a),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
