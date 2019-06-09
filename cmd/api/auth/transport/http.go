package transport

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth"
	"github.com/labstack/echo"
	"net/http"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, e *echo.Echo) {
	h := HTTP{svc}
	// swagger:route POST /login auth login
	// Logs in user by username and password.
	// responses:
	//  200: loginResp
	//  400: errMsg
	//  401: errMsg
	// 	403: err
	//  404: errMsg
	//  500: err
	e.POST("/login", h.login)
}

type credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(c echo.Context) error {

	cred := new(credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}
	r, err := h.svc.Authenticate(c, cred.Email, cred.Password)
	if err != nil {
		return err
	}
	/*return c.Render(http.StatusOK, "login.html", echo.Map{
		"User": r,
	})*/
	return c.JSON(http.StatusOK, r)
}
