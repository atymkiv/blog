package transport

import blog "github.com/atymkiv/echo_frame_learning/blog/model"

// Post create request
// swagger:parameters postCreate
type swaggPostReq struct {
	// in:body
	Body createReq
}

// Post model response
// swagger:response postResp
type swaggPostResponse struct {
	//in:body
	Body struct{
		*blog.Post
	}
}

// Posts feed model response
// swagger:response postsResp
type swaggPostsResponse struct {
	//in:body
	Body struct{
		Posts []blog.Post
	}
}
