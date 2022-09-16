package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrBadRequest   = echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	ErrInternal     = echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	ErrNotFound     = echo.NewHTTPError(http.StatusNotFound, "Not Found")
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	ErrValidation   = echo.NewHTTPError(http.StatusBadRequest, "Validation error")
)
