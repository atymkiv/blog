package user

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/messages"
	"github.com/labstack/echo"
	"log"
)

const TOPIC = "users"

// Service represents user application interface
type Service interface {
	Signup(echo.Context, blog.User) (*blog.User, error)
}

// New creates new user application serviceCreateFn
func New(udb UDB, nats *messages.Service) *User {
	return &User{udb: udb, nats: nats}
}

// User represents user application service
type User struct {
	udb  UDB
	nats *messages.Service
}

type UDB interface {
	Signup(blog.User) (*blog.User, error)
}

func (u *User) Signup(c echo.Context, req blog.User) (*blog.User, error) {
	if err := u.nats.PushMessage(req, TOPIC); err != nil {
		log.Printf("failed pushing user into nuts; err: %v", err)
		return nil, err
	}
	user, err := u.udb.Signup(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
