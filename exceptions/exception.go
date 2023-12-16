package exceptions

import (
	"net/http"

	"github.com/BerdanAkbulut/task-app-backend/pkg"
	"github.com/labstack/echo/v4"
)

func ThrowHttpError(message string, code int) error {
	return &pkg.HttpError{
		Message: message,
		Code: code,
	}
}

func GlobalExceptionHandler(e *echo.Echo) {
	e.HTTPErrorHandler = customHTTPErrorHandler
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		code                = http.StatusInternalServerError
		message interface{} = http.StatusText(code)
	)

	if ex, ok := err.(*pkg.HttpError); ok {
		code = ex.Code
		message = ex.Message
	} else if ex, ok := err.(*echo.HTTPError); ok {
		code = ex.Code
		message = ex.Message
	}

	c.JSON(code, map[string]interface{}{
		"error_message": message,
	})
}
