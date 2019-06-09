package blog

import "github.com/jinzhu/gorm"

type (
	Post struct {
		gorm.Model
		From    string `json:"from" gorm: NOT NULL"`
		Message string `json:"message" gorm: NOT NULL"`
	}
)
