package gormsql

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/jinzhu/gorm"
)

// NewPost returns a new post database instance
func NewPost(db Database) *Post {
	return &Post{
		db: db,
	}
}

// Post represents the client for post table
type Post struct {
	db Database
}

// Interface for post database
type Database interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
}

// Custom errors

func (p *Post) Create(pst blog.Post) (*blog.Post, error) {

	var post = new(blog.Post)

	post.From = pst.From
	post.Message = pst.Message

	// Save post
	if err := p.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Post) ViewAll() (*[]blog.Post, error) {
	// Retrieve posts from database
	posts := &[]blog.Post{}
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, blog.ErrGeneric
	}
	return posts, nil
}
