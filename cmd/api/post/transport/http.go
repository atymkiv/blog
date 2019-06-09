package transport

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/post"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/labstack/echo"
	"net/http"
)

type HTTP struct {
	svc post.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc post.Service, er *echo.Group) {
	h := HTTP{svc}

	// swagger:route POST /post/create posts postCreate
	// Creates new post.
	// responses:
	//  200: postResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	er.POST("/create", h.create)
	// swagger:route GET /post/feed posts postsView
	// Shows all posts.
	// responses:
	//  200: postsResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	er.GET("/feed", h.feed)

}

type createReq struct {
	Message string `json:"message" validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}

	from, err := h.svc.UserEmailFromToken(c)
	if err != nil {
		return err
	}

	p, err := h.svc.Create(c, blog.Post{
		From:    from,
		Message: r.Message,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (h *HTTP) feed(c echo.Context) error {
	users, err := h.svc.ViewAll(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
