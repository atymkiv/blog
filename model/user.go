package blog

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
)

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID    int
	Email string
}

// UpdateLastLogin updates last login field
func (u *User) UpdateLastLogin(token string) {
	u.Token = token
}
