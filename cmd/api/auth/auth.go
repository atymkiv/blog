package auth

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/labstack/echo"
	"github.com/ribice/gorsk/pkg/utl/model"
	"net/http"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")
)

// New creates new iam service
func New(udb UserDB, j TokenGenerator, sec Securer) *Auth {
	return &Auth{
		udb: udb,
		tg:  j,
		sec: sec,
	}
}

// Service represents auth service interface
type Service interface {
	Authenticate(echo.Context, string, string) (*blog.AuthToken, error)
}

// Auth represents auth application service
type Auth struct {
	udb UserDB
	tg  TokenGenerator
	sec Securer
}

// UserDB represents user repository interface
type UserDB interface {
	FindByEmail(string) (*blog.User, error)
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(*blog.User) (string, string, error)
}

// Securer represents security interface
type Securer interface {
	Token(string) string
}

// Authenticate tries to authenticate the user provided by email and password
func (a *Auth) Authenticate(c echo.Context, user, pass string) (*blog.AuthToken, error) {
	u, err := a.udb.FindByEmail(user)
	if err != nil {
		return nil, err
	}

	token, expire, err := a.tg.GenerateToken(u)
	if err != nil {
		return nil, gorsk.ErrUnauthorized
	}

	return &blog.AuthToken{Token: token, Expires: expire, RefreshToken: u.Token}, nil
}
