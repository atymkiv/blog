package transport_test

import (
	"bytes"
	"encoding/json"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth/transport"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *blog.AuthToken
		udb        *mockdb.User
		jwt        *mock.JWT
		sec        *mock.Secure
	}{
		{
			name:       "Invalid request",
			req:        `{"email":"juzernejm"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Fail on FindByEmail",
			req:        `{"email":"juzernejm","password":"hunter123"}`,
			wantStatus: http.StatusInternalServerError,
			udb: &mockdb.User{
				FindByEmailFn: func(string) (*blog.User, error) {
					return nil, blog.ErrGeneric
				},
			},
		},
		{
			name:       "Success",
			req:        `{"email":"juzernejm","password":"hunter123"}`,
			wantStatus: http.StatusOK,
			udb: &mockdb.User{
				FindByEmailFn: func(string) (*blog.User, error) {
					return &blog.User{
						Password: "hunter123",
					}, nil
				},
			},
			jwt: &mock.JWT{
				GenerateTokenFn: func(*blog.User) (string, string, error) {
					return "jwttokenstring", mock.TestTime(2018).Format(time.RFC3339), nil
				},
			},
			sec: &mock.Secure{
				TokenFn: func(string) string {
					return "refreshtoken"
				},
			},
			wantResp: &blog.AuthToken{Token: "jwttokenstring", Expires: mock.TestTime(2018).Format(time.RFC3339), RefreshToken: "refreshtoken"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			transport.NewHTTP(auth.New(tt.udb, tt.jwt, tt.sec), r)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/login"
			res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(blog.AuthToken)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				tt.wantResp.RefreshToken = response.RefreshToken
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}
