package post

import (
	"context"
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"log"
)

// Service represents post application interface
type Service interface {
	Create(echo.Context, blog.Post) (*blog.Post, error)
	ViewAll(echo.Context) (*[]blog.Post, error)
	UserEmailFromToken(echo.Context) (string, error)
	//Logout(c echo.Context) error
}

// New creates new user application service
func New(pdb Db, jwt TokenParser, grpcClient pb.RouteGuideClient) *Post {
	return &Post{pdb: pdb, jwt: jwt, grpcClient: grpcClient}
}

// User represents user application service
type Post struct {
	pdb        Db
	jwt        TokenParser
	grpcClient pb.RouteGuideClient
}

type Db interface {
	Create(blog.Post) (*blog.Post, error)
	ViewAll() (*[]blog.Post, error)
}
type TokenParser interface {
	ParseToken(echo.Context) (*jwt.Token, error)
}

func (u *Post) Create(c echo.Context, req blog.Post) (*blog.Post, error) {
	post, err := u.pdb.Create(req)
	if err != nil {
		return nil, err
	}

	//triggering grpc server to push post into redis
	_, err = u.grpcClient.CreatePost(context.Background(), &pb.Post{string(req.ID), req.From, req.Message})
	if err != nil {
		log.Printf("failed pushing post into redis \n err: %v", err)
	}

	return post, nil
}

func (u *Post) ViewAll(c echo.Context) (*[]blog.Post, error) {
	posts, err := u.pdb.ViewAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *Post) UserEmailFromToken(c echo.Context) (string, error) {
	user, err := u.jwt.ParseToken(c)
	if err != nil {
		return "", err
	}
	claims := user.Claims.(jwt.MapClaims)
	return claims["e"].(string), nil

}
