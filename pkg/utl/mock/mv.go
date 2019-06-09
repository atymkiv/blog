package mock

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*blog.User) (string, string, error)
	ParseTokenFn    func(echo.Context) (*jwt.Token, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *blog.User) (string, string, error) {
	return j.GenerateTokenFn(u)
}

func (j *JWT) ParseToken(c echo.Context) (*jwt.Token, error) {
	return j.ParseTokenFn(c)
}
