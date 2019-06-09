package logging

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/labstack/echo"
)

// New creates new auth logging service
func New(svc auth.Service) *LogService {
	return &LogService{
		Service: svc,
	}
}

// LogService represents auth logging service
type LogService struct {
	auth.Service
}

// Authenticate logging
func (ls *LogService) Authenticate(c echo.Context, user, password string) (resp *blog.AuthToken, err error) {
	return ls.Service.Authenticate(c, user, password)
}
