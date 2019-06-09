package transport

import blog "github.com/atymkiv/echo_frame_learning/blog/model"

// Login request
// swagger:parameters login
type swaggLoginReq struct {
	// in:body
	Body credentials
}

// Login response
// swagger:response loginResp
type swaggLoginResp struct {
	// in:body
	Body struct {
		*blog.AuthToken
	}
}