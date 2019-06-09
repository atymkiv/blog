package auth_test

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/ribice/gorsk/pkg/utl/model"
	"github.com/stretchr/testify/assert"
	"time"

	"testing"
)

func TestAuthenticate(t *testing.T) {
	type args struct {
		email string
		pass  string
	}
	cases := []struct {
		name string
		args args

		wantErr bool
		udb     *mockdb.User
		jwt     *mock.JWT
		sec     *mock.Secure
	}{
		{
			name:    "Fail on finding user",
			args:    args{email: "test"},
			wantErr: true,
			udb: &mockdb.User{
				FindByEmailFn: func(user string) (*blog.User, error) {
					return nil, blog.ErrGeneric
				},
			},
		},
		{
			name:    "Fail on generate token",
			args:    args{email: "test", pass: "123"},
			wantErr: true,
			udb: &mockdb.User{
				FindByEmailFn: func(user string) (*blog.User, error) {
					return &blog.User{
						Email:    user,
						Password: "123",
					}, nil
				},
			},
			jwt: &mock.JWT{GenerateTokenFn: func(u *blog.User) (string, string, error) {
				return "", "", gorsk.ErrGeneric
			},
			},
		},
		{
			name:    "success",
			args:    args{email: "test@l.com", pass: "123"},
			wantErr: false,
			udb: &mockdb.User{
				FindByEmailFn: func(email string) (*blog.User, error) {
					return &blog.User{
						Email:    email,
						Password: "123",
					}, nil
				},
			},
			jwt: &mock.JWT{GenerateTokenFn: func(u *blog.User) (string, string, error) {
				return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", mock.TestTime(2000).Format(time.RFC3339), nil
			},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := auth.New(tt.udb, tt.jwt, tt.sec)
			_, err := s.Authenticate(nil, tt.args.email, tt.args.pass)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
