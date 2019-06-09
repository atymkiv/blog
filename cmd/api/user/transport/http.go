package transport

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/user"
	"github.com/atymkiv/echo_frame_learning/blog/model"
	"github.com/labstack/echo"
	"net/http"
)

type HTTP struct {
	svc user.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc user.Service, e *echo.Echo) {
	h := HTTP{svc}
	// swagger:route POST /signup users userCreate
	// Creates new user account.
	// responses:
	//  200: userResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	e.POST("/signup", h.signup)

}

type authReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) signup(c echo.Context) (err error) {
	r := new(authReq)

	if err := c.Bind(r); err != nil {

		return err
	}

	u, err := h.svc.Signup(c, blog.User{
		Email:    r.Email,
		Password: r.Password,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
