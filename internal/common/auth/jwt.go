package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       int
	Role     string
	Username string
}

type Claims struct {
	UserID   int    `json:"userID"`
	UserRole string `json:"userRole"`
	jwt.RegisteredClaims
}

type JWT struct {
	secret   string
	duration time.Duration
}

func NewJWT(secret string, duration time.Duration) JWT {
	return JWT{
		secret:   secret,
		duration: duration,
	}
}

func (t JWT) Validate(s string) (User, error) {
	var c Claims
	token, err := jwt.ParseWithClaims(
		s,
		&c,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.secret), nil
		},
	)

	if err != nil {
		return User{}, err
	}

	if !token.Valid {
		return User{}, errors.New("invalid token")
	}

	return User{
		ID:       c.UserID,
		Role:     c.UserRole,
		Username: c.Subject,
	}, nil
}

func (t JWT) Generate(u User) (string, error) {
	claims := Claims{
		UserID:   u.ID,
		UserRole: u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.secret))
}
