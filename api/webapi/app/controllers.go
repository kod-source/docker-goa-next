// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Controllers
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
	"github.com/shogo82148/goa-v1/cors"
	"net/http"
	"regexp"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Login(*LoginAuthContext) error
	SignUp(*SignUpAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/login", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/sign_up", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LoginAuthPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/login", ctrl.MuxHandler("login", h, unmarshalLoginAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Login", "route", "POST /login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSignUpAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SignUpAuthPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.SignUp(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/sign_up", ctrl.MuxHandler("sign_up", h, unmarshalSignUpAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "SignUp", "route", "POST /sign_up")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLoginAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &loginAuthPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalSignUpAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalSignUpAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &signUpAuthPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Add(*AddOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/add/:left/:right", ctrl.MuxHandler("preflight", handleOperandsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddOperandsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	h = handleOperandsOrigin(h)
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("add", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Add", "route", "GET /add/:left/:right")
}

// handleOperandsOrigin applies the CORS response headers corresponding to the origin.
func handleOperandsOrigin(h goa.Handler) goa.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	GetCurrentUser(*GetCurrentUserUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/current_user", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetCurrentUserUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetCurrentUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/current_user", ctrl.MuxHandler("get_current_user", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "GetCurrentUser", "route", "GET /current_user", "security", "jwt")
}

// handleUsersOrigin applies the CORS response headers corresponding to the origin.
func handleUsersOrigin(h goa.Handler) goa.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
