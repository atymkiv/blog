package gormsql

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/jinzhu/gorm"
)

// NewUser returns a new user database instance
func NewUser(db Database) *User {
	return &User{
		db: db,
	}
}

// User represents the client for user table
type User struct {
	db Database
}

// Interface for post database
type Database interface {
	Create(value interface{}) *gorm.DB
}

func (u *User) Signup(usr blog.User) (*blog.User, error) {
	var user = new(blog.User)

	user.Email = usr.Email
	user.Password = usr.Password

	// Save user
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
