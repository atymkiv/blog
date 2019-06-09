package transport_test

import (
	"bytes"
	"encoding/json"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/post"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/post/transport"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/server"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
		wantResp   *blog.Post
		pdb        *mockdb.Post
		jwt        *mock.JWT
	}{
		{
			name:       "Fail on validation",
			wantStatus: http.StatusBadRequest,
		},

		{
			name: "Fail on getting user email from Token",
			req:  `{"message" : "hi there"}`,

			jwt: &mock.JWT{
				ParseTokenFn: func(c echo.Context) (*jwt.Token, error) {
					return nil, blog.ErrGeneric
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "Fail on creating",
			req:  `{"message" : "hi there"}`,
			jwt: &mock.JWT{
				ParseTokenFn: func(c echo.Context) (*jwt.Token, error) {
					token, _ := jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlIjoidGVzdEBsLmNvbSIsImlkIjoyNX0.k21UnYNmZEpV-uzuZpJsTYYNd7m7VX6m4nj2HQHbEe4", func(token *jwt.Token) (interface{}, error) {
						return []byte("foobar"), nil
					})
					return token, nil
				},
			},
			pdb: &mockdb.Post{
				CreateFn: func(pst blog.Post) (*blog.Post, error) {
					return nil, blog.ErrGeneric
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			req:  `{"message" : "hi there"}`,
			jwt: &mock.JWT{
				ParseTokenFn: func(c echo.Context) (*jwt.Token, error) {
					token, _ := jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlIjoidGVzdEBsLmNvbSIsImlkIjoyNX0.k21UnYNmZEpV-uzuZpJsTYYNd7m7VX6m4nj2HQHbEe4", func(token *jwt.Token) (interface{}, error) {
						return []byte("foobar"), nil
					})
					return token, nil
				},
			},
			pdb: &mockdb.Post{
				CreateFn: func(pst blog.Post) (*blog.Post, error) {
					return &blog.Post{
						From:    "test@lpnu.com",
						Message: "hi there",
					}, nil
				},
			},
			wantResp: &blog.Post{
				From:    "test@lpnu.com",
				Message: "hi there",
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			transport.NewHTTP(post.New(tt.pdb, tt.jwt), rg)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/create"
			res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(blog.Post)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}

func TestFeed(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *blog.User
		pdb        *mockdb.Post
	}{
		{
			name: "Fail on viewing",
			pdb: &mockdb.Post{
				ViewAllFn: func() (*[]blog.Post, error) {
					return nil, blog.ErrGeneric
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			pdb: &mockdb.Post{
				ViewAllFn: func() (*[]blog.Post, error) {
					return &[]blog.Post{
						{
							From:    "test@l.com",
							Message: "hi there",
						},
						{
							From:    "a@l.com",
							Message: "hello",
						},
					}, nil
				},
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			rg := r.Group("")
			transport.NewHTTP(post.New(tt.pdb, nil), rg)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/feed"
			res, err := http.Get(path)
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
