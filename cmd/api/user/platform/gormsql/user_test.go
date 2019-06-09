package gormsql_test

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user/platform/gormsql"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	db := mockdb.NewFakeDbOrFatal()
	defer db.Close()
	cases := []struct {
		name    string
		req     *blog.User
		wantErr bool
	}{
		{
			name:    "success",
			req:     &blog.User{Email: "a@l.com", Password: "123"},
			wantErr: false,
		},
		{
			name:    "Failed on pushing same email twice",
			req:     &blog.User{Email: "a@l.com", Password: "123"},
			wantErr: true,
		},
	}
	userDB := gormsql.NewUser(db)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userDB.Signup(*tt.req)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
