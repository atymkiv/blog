package mockdb

import (
	"github.com/atymkiv/echo_frame_learning/blog/model"
)

// User database mock
type Post struct {
	CreateFn  func(blog.Post) (*blog.Post, error)
	ViewAllFn func() (*[]blog.Post, error)
}

// Create mock
func (p *Post) Create(pst blog.Post) (*blog.Post, error) {
	return p.CreateFn(pst)
}

// ViewAll mock
func (p *Post) ViewAll() (*[]blog.Post, error) {
	return p.ViewAllFn()
}
