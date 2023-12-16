package jwt

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretJwtKey = []byte("supersecretjwtkey")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (*string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &JWTClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretJwtKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretJwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("Invalid token")
		return
	}
	if claims.ExpiresAt.Before(time.Now()) {
		err = errors.New("Token Expired")
		return
	}
	return
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if strings.HasPrefix(c.Request().URL.Path, "/auth") {
			return next(c)
		}
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is required")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		err := ValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}
