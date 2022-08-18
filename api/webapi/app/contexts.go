// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=$(GOPATH)/src/app/webapi
// --version=v1.5.13

package app

import (
	"context"
	goa "github.com/shogo82148/goa-v1"
	"net/http"
	"strconv"
)

// LoginAuthContext provides the auth login action context.
type LoginAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *LoginAuthPayload
}

// NewLoginAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller login action.
func NewLoginAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// loginAuthPayload is the auth login action payload.
type loginAuthPayload struct {
	// メール
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *loginAuthPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates LoginAuthPayload from loginAuthPayload
func (payload *loginAuthPayload) Publicize() *LoginAuthPayload {
	var pub LoginAuthPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// LoginAuthPayload is the auth login action payload.
type LoginAuthPayload struct {
	// メール
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// パスワード
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *LoginAuthPayload) Validate() (err error) {

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginAuthContext) OK(r *Token) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginAuthContext) BadRequest(r *ServiceVerror) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.service.verror")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LoginAuthContext) NotFound(r *ServiceVerror) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.service.verror")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *LoginAuthContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// SignUpAuthContext provides the auth sign_up action context.
type SignUpAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *SignUpAuthPayload
}

// NewSignUpAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller sign_up action.
func NewSignUpAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*SignUpAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SignUpAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// signUpAuthPayload is the auth sign_up action payload.
type signUpAuthPayload struct {
	// プロフィール画像のパス
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// メール
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// 名前
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *signUpAuthPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates SignUpAuthPayload from signUpAuthPayload
func (payload *signUpAuthPayload) Publicize() *SignUpAuthPayload {
	var pub SignUpAuthPayload
	if payload.Avatar != nil {
		pub.Avatar = payload.Avatar
	}
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// SignUpAuthPayload is the auth sign_up action payload.
type SignUpAuthPayload struct {
	// プロフィール画像のパス
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// メール
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// 名前
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// パスワード
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *SignUpAuthPayload) Validate() (err error) {

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *SignUpAuthContext) Created(r *Token) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *SignUpAuthContext) BadRequest(r *ServiceVerror) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.service.verror")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SignUpAuthContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CreateCommentCommentsContext provides the comments create_comment action context.
type CreateCommentCommentsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCommentCommentsPayload
}

// NewCreateCommentCommentsContext parses the incoming request URL and body, performs validations and creates the
// context used by the comments controller create_comment action.
func NewCreateCommentCommentsContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateCommentCommentsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateCommentCommentsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCommentCommentsPayload is the comments create_comment action payload.
type createCommentCommentsPayload struct {
	// コメント画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// 投稿ID
	PostID *int `form:"post_id,omitempty" json:"post_id,omitempty" yaml:"post_id,omitempty" xml:"post_id,omitempty"`
	// コメントの内容
	Text *string `form:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty" xml:"text,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createCommentCommentsPayload) Validate() (err error) {
	if payload.PostID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "post_id"))
	}
	if payload.Text == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "text"))
	}
	return
}

// Publicize creates CreateCommentCommentsPayload from createCommentCommentsPayload
func (payload *createCommentCommentsPayload) Publicize() *CreateCommentCommentsPayload {
	var pub CreateCommentCommentsPayload
	if payload.Img != nil {
		pub.Img = payload.Img
	}
	if payload.PostID != nil {
		pub.PostID = *payload.PostID
	}
	if payload.Text != nil {
		pub.Text = *payload.Text
	}
	return &pub
}

// CreateCommentCommentsPayload is the comments create_comment action payload.
type CreateCommentCommentsPayload struct {
	// コメント画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// 投稿ID
	PostID int `form:"post_id" json:"post_id" yaml:"post_id" xml:"post_id"`
	// コメントの内容
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateCommentCommentsPayload) Validate() (err error) {

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateCommentCommentsContext) Created(r *CommentJSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.comment_json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCommentCommentsContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateCommentCommentsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteCommentCommentsContext provides the comments delete_comment action context.
type DeleteCommentCommentsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeleteCommentCommentsContext parses the incoming request URL and body, performs validations and creates the
// context used by the comments controller delete_comment action.
func NewDeleteCommentCommentsContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteCommentCommentsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteCommentCommentsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteCommentCommentsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteCommentCommentsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCommentCommentsContext provides the comments show_comment action context.
type ShowCommentCommentsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PostID int
}

// NewShowCommentCommentsContext parses the incoming request URL and body, performs validations and creates the
// context used by the comments controller show_comment action.
func NewShowCommentCommentsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCommentCommentsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCommentCommentsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPostID := req.Params["post_id"]
	if len(paramPostID) > 0 {
		rawPostID := paramPostID[0]
		if postID, err2 := strconv.Atoi(rawPostID); err2 == nil {
			rctx.PostID = postID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("post_id", rawPostID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCommentCommentsContext) OK(r CommentJSONCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.comment_json; type=collection")
	}
	if r == nil {
		r = CommentJSONCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCommentCommentsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCommentCommentsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// UpdateCommentCommentsContext provides the comments update_comment action context.
type UpdateCommentCommentsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *UpdateCommentCommentsPayload
}

// NewUpdateCommentCommentsContext parses the incoming request URL and body, performs validations and creates the
// context used by the comments controller update_comment action.
func NewUpdateCommentCommentsContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateCommentCommentsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateCommentCommentsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// updateCommentCommentsPayload is the comments update_comment action payload.
type updateCommentCommentsPayload struct {
	// コメント画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// コメントの内容
	Text *string `form:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty" xml:"text,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateCommentCommentsPayload) Validate() (err error) {
	if payload.Text == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "text"))
	}
	if payload.Img == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "img"))
	}
	return
}

// Publicize creates UpdateCommentCommentsPayload from updateCommentCommentsPayload
func (payload *updateCommentCommentsPayload) Publicize() *UpdateCommentCommentsPayload {
	var pub UpdateCommentCommentsPayload
	if payload.Img != nil {
		pub.Img = *payload.Img
	}
	if payload.Text != nil {
		pub.Text = *payload.Text
	}
	return &pub
}

// UpdateCommentCommentsPayload is the comments update_comment action payload.
type UpdateCommentCommentsPayload struct {
	// コメント画像のパス
	Img string `form:"img" json:"img" yaml:"img" xml:"img"`
	// コメントの内容
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateCommentCommentsPayload) Validate() (err error) {

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateCommentCommentsContext) OK(r *CommentJSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.comment_json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateCommentCommentsContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateCommentCommentsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CreatePostPostsContext provides the posts create_post action context.
type CreatePostPostsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreatePostPostsPayload
}

// NewCreatePostPostsContext parses the incoming request URL and body, performs validations and creates the
// context used by the posts controller create_post action.
func NewCreatePostPostsContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePostPostsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePostPostsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createPostPostsPayload is the posts create_post action payload.
type createPostPostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createPostPostsPayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	return
}

// Publicize creates CreatePostPostsPayload from createPostPostsPayload
func (payload *createPostPostsPayload) Publicize() *CreatePostPostsPayload {
	var pub CreatePostPostsPayload
	if payload.Img != nil {
		pub.Img = payload.Img
	}
	if payload.Title != nil {
		pub.Title = *payload.Title
	}
	return &pub
}

// CreatePostPostsPayload is the posts create_post action payload.
type CreatePostPostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePostPostsContext) Created(r *IndexPostJSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.index_post_json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreatePostPostsContext) BadRequest(r *ServiceVerror) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.service.verror")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreatePostPostsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeletePostsContext provides the posts delete action context.
type DeletePostsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeletePostsContext parses the incoming request URL and body, performs validations and creates the
// context used by the posts controller delete action.
func NewDeletePostsContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeletePostsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeletePostsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeletePostsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeletePostsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// IndexPostsContext provides the posts index action context.
type IndexPostsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewIndexPostsContext parses the incoming request URL and body, performs validations and creates the
// context used by the posts controller index action.
func NewIndexPostsContext(ctx context.Context, r *http.Request, service *goa.Service) (*IndexPostsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := IndexPostsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *IndexPostsContext) OK(r IndexPostJSONCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.index_post_json; type=collection")
	}
	if r == nil {
		r = IndexPostJSONCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *IndexPostsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *IndexPostsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPostsContext provides the posts show action context.
type ShowPostsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowPostsContext parses the incoming request URL and body, performs validations and creates the
// context used by the posts controller show action.
func NewShowPostsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowPostsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowPostsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPostsContext) OK(r *ShowPostJSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.show_post_json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPostsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowPostsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// UpdatePostsContext provides the posts update action context.
type UpdatePostsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *UpdatePostsPayload
}

// NewUpdatePostsContext parses the incoming request URL and body, performs validations and creates the
// context used by the posts controller update action.
func NewUpdatePostsContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdatePostsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdatePostsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// updatePostsPayload is the posts update action payload.
type updatePostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updatePostsPayload) Validate() (err error) {
	if payload.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "title"))
	}
	return
}

// Publicize creates UpdatePostsPayload from updatePostsPayload
func (payload *updatePostsPayload) Publicize() *UpdatePostsPayload {
	var pub UpdatePostsPayload
	if payload.Img != nil {
		pub.Img = payload.Img
	}
	if payload.Title != nil {
		pub.Title = *payload.Title
	}
	return &pub
}

// UpdatePostsPayload is the posts update action payload.
type UpdatePostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdatePostsContext) OK(r *IndexPostJSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.index_post_json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdatePostsContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdatePostsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// GetCurrentUserUsersContext provides the users get_current_user action context.
type GetCurrentUserUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetCurrentUserUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller get_current_user action.
func NewGetCurrentUserUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetCurrentUserUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetCurrentUserUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetCurrentUserUsersContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetCurrentUserUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetCurrentUserUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
