package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Responses struct {
	Code    int
	Message interface{}
	Data    interface{}
}

func ResponseContext(code int, message interface{}, data interface{}, c echo.Context) error {
	if code == 200 { // Success
		return successContext(message, data, c)
	} else if code == 400 { // Bad Request
		return validationContext(message, data, c)
	} else if code == 404 { // Notfound
		return notFoundContext(message, data, c)
	} else if code == 504 { // Timeout
		return timeoutContext(message, c)
	}
	return errorContext(message, c) // Internal Server Error
}

func successContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"StatusCode": 200,
		"Status":     "success",
		"Message":    message,
		"Data":       data,
	})
}

func errorContext(message interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"StatusCode": 500,
		"Status":     "failed",
		"Message":    message,
		"Data":       nil,
	})
}

func validationContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"StatusCode": 400,
		"Status":     "validation",
		"Message":    message,
		"Data":       data,
	})
}

func timeoutContext(message interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"StatusCode": 504,
		"Status":     "timeout",
		"Message":    message,
		"Data":       nil,
	})
}

func notFoundContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"StatusCode": 404,
		"Status":     "not found",
		"Message":    message,
		"Data":       nil,
	})
}
