package post_test

import (
	"context"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/post"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestCreate(t *testing.T) {

	cases := []struct {
		name       string
		req        blog.Post
		wantErr    bool
		pdb        *mockdb.Post
		grpcClient *mock.GrpcClient
	}{
		{
			name:    "Fail on creating error",
			req:     blog.Post{From: "test@l.com"},
			wantErr: true,
			pdb: &mockdb.Post{
				CreateFn: func(pst blog.Post) (*blog.Post, error) {
					return nil, blog.ErrGeneric
				},
			},
		},

		{
			name:    "success",
			req:     blog.Post{From: "test@l.com", Message: "hi there"},
			wantErr: false,
			pdb: &mockdb.Post{
				CreateFn: func(pst blog.Post) (*blog.Post, error) {
					return &blog.Post{
						From:    "test@l.com",
						Message: "hi there",
					}, nil
				},
			},
			grpcClient: &mock.GrpcClient{
				CreatePostFn: func(ctx context.Context, in *routeguide.Post, opts ...grpc.CallOption) (*routeguide.Result, error) {
					return &pb.Result{Code: 0}, nil
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := post.New(tt.pdb, nil, tt.grpcClient)
			_, err := s.Create(nil, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestViewAll(t *testing.T) {
	cases := []struct {
		name string
		req  blog.Post

		wantErr bool
		pdb     *mockdb.Post
	}{
		{
			name:    "Fail on viewAll error",
			req:     blog.Post{From: "test@l.com"},
			wantErr: true,
			pdb: &mockdb.Post{
				ViewAllFn: func() (*[]blog.Post, error) {
					return nil, blog.ErrGeneric
				},
			},
		},
		{
			name:    "success",
			req:     blog.Post{From: "test@l.com", Message: "hi there"},
			wantErr: false,
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
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := post.New(tt.pdb, nil, nil)
			_, err := s.ViewAll(nil)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUserEmailFromToken(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		wantResp string
		jwt      *mock.JWT
	}{
		{
			name:    "Fail on returning token",
			wantErr: true,
			jwt: &mock.JWT{
				ParseTokenFn: func(c echo.Context) (*jwt.Token, error) {
					return nil, blog.ErrGeneric
				},
			},
		},
		{
			name:    "success",
			wantErr: false,
			jwt: &mock.JWT{
				ParseTokenFn: func(c echo.Context) (*jwt.Token, error) {
					token, _ := jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlIjoidGVzdEBsLmNvbSIsImlkIjoyNX0.k21UnYNmZEpV-uzuZpJsTYYNd7m7VX6m4nj2HQHbEe4", func(token *jwt.Token) (interface{}, error) {
						return []byte("foobar"), nil
					})
					return token, nil
				},
			},
			wantResp: "test@l.com",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := post.New(nil, tt.jwt, nil)
			response, err := s.UserEmailFromToken(nil)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantResp != "" {
				assert.Equal(t, tt.wantResp, response)
			}
		})

	}
}
