package gormsql

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/jinzhu/gorm"
)

// NewUser returns a new user database instance
func NewUser(db Database) *User {
	return &User{db: db}
}

// Interface for post database
type Database interface {
	Where(query interface{}, args ...interface{}) *gorm.DB
}

// User represents the client for user table
type User struct {
	db Database
}

// FindByEmail queries for single user by email
func (u *User) FindByEmail(email string) (*blog.User, error) {
	var user = new(blog.User)
	if err := u.db.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
