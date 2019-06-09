package mockdb

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
)

// User database mock
type User struct {
	SignupFn      func(blog.User) (*blog.User, error)
	FindByEmailFn func(string) (*blog.User, error)
}

// Create mock
func (u *User) Signup(usr blog.User) (*blog.User, error) {
	return u.SignupFn(usr)
}

// FindByUsername mock
func (u *User) FindByEmail(umeil string) (*blog.User, error) {
	return u.FindByEmailFn(umeil)
}
