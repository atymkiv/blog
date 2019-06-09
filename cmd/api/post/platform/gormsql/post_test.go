package gormsql_test

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/post/platform/gormsql"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	db := mockdb.NewFakeDbOrFatal()

	cases := []struct {
		name    string
		req     *blog.Post
		wantErr bool
	}{
		{
			name: "Fail on empty message",
			req: &blog.Post{
				From: "test@l.com",
			},
			wantErr: false,
		},
		{
			name:    "success",
			req:     &blog.Post{},
			wantErr: false,
		},
	}
	postDB := gormsql.NewPost(db)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := postDB.Create(*tt.req)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}

func TestViewAll(t *testing.T) {
	db := mockdb.NewFakeDbOrFatal()
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Fail on finding posts",
			wantErr: false,
		},
	}
	postDB := gormsql.NewPost(db)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := postDB.ViewAll()
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
