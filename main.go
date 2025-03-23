package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lmittmann/tint"
	"github.com/tinhtt/go-doctor-appointment/internal/adapters"
	httpport "github.com/tinhtt/go-doctor-appointment/internal/ports/http"
)

func main() {
	l := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if err, ok := a.Value.Any().(error); ok {
				aErr := tint.Err(err)
				aErr.Key = a.Key
				return aErr
			}
			return a
		},
	}))
	l.Info("start server")

	db, err := adapters.NewPostgresDB()
	if err != nil {
		l.Error("unable to connect database", "err", err)
		os.Exit(1)
	}
	defer db.Close()
	l.Info("connect database successfully!")

	err = adapters.Migrate()
	if err != nil {
		l.Error("unable to run database migration", "err", err)
		os.Exit(1)
	}
	l.Info("run database migration successfully!")

	s := httpport.NewHTTPServer()
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error("unable to run server", "err", err)
			os.Exit(1)
		}
	}()
	l.Info("server is running")

	// Wait for interrupt signal to gracefully shutdown the application
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	l.Info("stop server")
	s.Shutdown(context.Background())
	l.Info("server existed")

}
