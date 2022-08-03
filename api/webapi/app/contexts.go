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
func (ctx *LoginAuthContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LoginAuthContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *LoginAuthContext) InternalServerError() error {
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

// SignUpUsersContext provides the users sign_up action context.
type SignUpUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *SignUpUsersPayload
}

// NewSignUpUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller sign_up action.
func NewSignUpUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*SignUpUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SignUpUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// signUpUsersPayload is the users sign_up action payload.
type signUpUsersPayload struct {
	// メール
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// 名前
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *signUpUsersPayload) Validate() (err error) {
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

// Publicize creates SignUpUsersPayload from signUpUsersPayload
func (payload *signUpUsersPayload) Publicize() *SignUpUsersPayload {
	var pub SignUpUsersPayload
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

// SignUpUsersPayload is the users sign_up action payload.
type SignUpUsersPayload struct {
	// メール
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// 名前
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// パスワード
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *SignUpUsersPayload) Validate() (err error) {

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *SignUpUsersContext) Created(r *Token) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *SignUpUsersContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SignUpUsersContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
