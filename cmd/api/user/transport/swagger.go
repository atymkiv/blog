package transport

import blog "github.com/atymkiv/echo_frame_learning/blog/model"

// User create request
// swagger:parameters userCreate
type swaggSignupReq struct {
	// in:body
	Body authReq
}

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	//in:body
	Body struct{
		*blog.User
	}
}