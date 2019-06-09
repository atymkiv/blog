package blog

import (
	"errors"
	"github.com/labstack/echo"
)

var (
	// ErrGeneric is used for testing purposes and for errors handled later in the callstack
	ErrGeneric = errors.New("generic error")

	// ErrBadRequest (400) is returned for bad request (validation)
	ErrBadRequest = echo.NewHTTPError(400)

	// ErrUnauthorized (401) is returned when user is not authorized
	ErrUnauthorized = echo.ErrUnauthorized

	ErrNoUserWithID = errors.New("no user with this id")

	ErrNoUserWithEmail = errors.New("no user with this email")

	ErrNoUserWithToken = errors.New("no user with this token")
)
