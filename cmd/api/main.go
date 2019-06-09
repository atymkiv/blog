// This documentation describes example APIs found under https://github.com/ribice/golang-swaggerui-example
//
//     Schemes: http
//     Version: 0.0.1
//     Contact: Andriy Tymkiv <a.tymkiv99@gmail.com>
//     Host: localhost/goswagg
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta

package main

import (
	"crypto/sha1"
	"flag"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth"
	al "github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth/logging"
	gormmsqlA "github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth/platform/gormsql"
	at "github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth/transport"
	ps "github.com/atymkiv/echo_frame_learning/blog/cmd/api/post"
	gormsqlP "github.com/atymkiv/echo_frame_learning/blog/cmd/api/post/platform/gormsql"
	ptr "github.com/atymkiv/echo_frame_learning/blog/cmd/api/post/transport"
	us "github.com/atymkiv/echo_frame_learning/blog/cmd/api/user"
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user/platform/gormsql"
	ut "github.com/atymkiv/echo_frame_learning/blog/cmd/api/user/transport"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/middleware/jwt"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/gorm"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/grpc"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/messages"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/nats"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/secure"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/server"
)

func main() {

	cfgPath := flag.String("p", "./cmd/api/config.json", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(Start(cfg))
}

func Start(cfg *config.Configuration) error {
	db, err := gorm.New(&cfg.DB)
	if err != nil {
		return err
	}
	sec := secure.New(sha1.New())
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)

	e := server.New()

	authDB := gormmsqlA.NewUser(db)
	at.NewHTTP(al.New(auth.New(authDB, jwt, sec)), e)

	natsClient, err := nats.New(cfg.Nats)
	checkErr(err)
	messageService := messages.Create(natsClient)
	userDB := gormsql.NewUser(db)
	ut.NewHTTP(us.New(userDB, messageService), e)

	v2 := e.Group("/post")
	v2.Use(jwt.MWFunc())
	postDB := gormsqlP.NewPost(db)
	grpcClient, err := grpc.New(cfg.GRPC)
	checkErr(err)
	ptr.NewHTTP(ps.New(postDB, jwt, grpcClient), v2)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
