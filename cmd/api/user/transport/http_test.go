package transport_test

import (
	"bytes"
	"encoding/json"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user/transport"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/server"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *blog.User
		udb        *mockdb.User
		sec        *mock.Secure
	}{
		{
			name:       "Fail on validation on password",
			req:        `{"email":"a@l.com"}"`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Fail on validation on email",
			req:        `{"password":"123"}"`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Success",
			req:  `{"email":"a@l.com", "password":"123"}`,

			udb: &mockdb.User{
				SignupFn: func(usr blog.User) (*blog.User, error) {
					return &usr, nil
				},
			},
			wantResp: &blog.User{
				Email:    "a@l.com",
				Password: "123",
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			rg := server.New()
			transport.NewHTTP(user.New(tt.udb), rg)
			ts := httptest.NewServer(rg)
			defer ts.Close()
			path := ts.URL + "/signup"
			res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(blog.User)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}
