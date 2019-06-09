package user_test

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignup(t *testing.T) {
	cases := []struct {
		name    string
		req     blog.User
		wantErr bool
		udb     *mockdb.User
	}{
		{
			name:    "Fail on creating error",
			req:     blog.User{Email: "a@l.com", Password: "123"},
			wantErr: true,
			udb: &mockdb.User{
				SignupFn: func(usr blog.User) (*blog.User, error) {
					return nil, blog.ErrGeneric
				},
			},
		},
		{
			name:    "success",
			req:     blog.User{Email: "a@l.com", Password: "123"},
			wantErr: false,
			udb: &mockdb.User{
				SignupFn: func(usr blog.User) (*blog.User, error) {
					return &blog.User{
						Email:    "a@l.com",
						Password: "123",
					}, nil
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(tt.udb)
			_, err := s.Signup(nil, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
