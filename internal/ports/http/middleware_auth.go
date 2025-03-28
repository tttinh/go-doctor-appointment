package httpport

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/common/auth"
)

const userKey = "user_1743122841"

func userFromContext(c *gin.Context) (auth.User, bool) {
	i, ok := c.Get(userKey)
	if !ok {
		return auth.User{}, false
	}

	u, ok := i.(auth.User)
	return u, ok
}

func authMiddleware(jwt auth.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		token := strings.Split(bearerToken, " ")[1]
		u, err := jwt.Validate(token)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set(userKey, u)
		c.Next()
	}
}
