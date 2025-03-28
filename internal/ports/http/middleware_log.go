package httpport

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func logMiddleware(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer.
		start := time.Now()

		// Process request.
		c.Next()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Stop timer.
		l := l.With(
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"elapsed", time.Since(start),
			"client_ip", c.ClientIP(),
			"response_size", c.Writer.Size(),
		)

		u, ok := userFromContext(c)
		if ok {
			a := slog.Group(
				"user",
				"id", u.ID,
				"role", u.Role,
				"username", u.Username,
			)
			l = l.With(a)
		}

		// No error.
		if len(c.Errors) == 0 {
			l.Info("http")
			return
		}

		// Input errors.
		if c.Writer.Status() < http.StatusInternalServerError {
			l.Warn("http", "err", errors.Unwrap(c.Errors.Last()))
			return
		}

		// Unknown errors.
		l.Error("http", "err", errors.Unwrap(c.Errors.Last()))
	}
}
