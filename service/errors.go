package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrMalformedInput = echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
var ErrDatabaseError = echo.NewHTTPError(http.StatusInternalServerError, "Database issue")
