// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Controllers
//
// Command:
// $ main

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

// CommentsController is the controller interface for the Comments actions.
type CommentsController interface {
	goa.Muxer
	CreateComment(*CreateCommentCommentsContext) error
	DeleteComment(*DeleteCommentCommentsContext) error
	ShowComment(*ShowCommentCommentsContext) error
	UpdateComment(*UpdateCommentCommentsContext) error
}

// MountCommentsController "mounts" a Comments resource controller on the given service.
func MountCommentsController(service *goa.Service, ctrl CommentsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/comments", ctrl.MuxHandler("preflight", handleCommentsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/comments/:id", ctrl.MuxHandler("preflight", handleCommentsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCommentCommentsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCommentCommentsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateComment(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleCommentsOrigin(h)
	service.Mux.Handle("POST", "/comments", ctrl.MuxHandler("create_comment", h, unmarshalCreateCommentCommentsPayload))
	service.LogInfo("mount", "ctrl", "Comments", "action", "CreateComment", "route", "POST /comments", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCommentCommentsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteComment(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleCommentsOrigin(h)
	service.Mux.Handle("DELETE", "/comments/:id", ctrl.MuxHandler("delete_comment", h, nil))
	service.LogInfo("mount", "ctrl", "Comments", "action", "DeleteComment", "route", "DELETE /comments/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCommentCommentsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowComment(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleCommentsOrigin(h)
	service.Mux.Handle("GET", "/comments/:id", ctrl.MuxHandler("show_comment", h, nil))
	service.LogInfo("mount", "ctrl", "Comments", "action", "ShowComment", "route", "GET /comments/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCommentCommentsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateCommentCommentsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateComment(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleCommentsOrigin(h)
	service.Mux.Handle("PUT", "/comments/:id", ctrl.MuxHandler("update_comment", h, unmarshalUpdateCommentCommentsPayload))
	service.LogInfo("mount", "ctrl", "Comments", "action", "UpdateComment", "route", "PUT /comments/:id", "security", "jwt")
}

// handleCommentsOrigin applies the CORS response headers corresponding to the origin.
func handleCommentsOrigin(h goa.Handler) goa.Handler {
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

// unmarshalCreateCommentCommentsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCommentCommentsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCommentCommentsPayload{}
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

// unmarshalUpdateCommentCommentsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCommentCommentsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateCommentCommentsPayload{}
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

// LikesController is the controller interface for the Likes actions.
type LikesController interface {
	goa.Muxer
	Create(*CreateLikesContext) error
	Delete(*DeleteLikesContext) error
	GetLikeByUser(*GetLikeByUserLikesContext) error
	GetMyLike(*GetMyLikeLikesContext) error
}

// MountLikesController "mounts" a Likes resource controller on the given service.
func MountLikesController(service *goa.Service, ctrl LikesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/likes", ctrl.MuxHandler("preflight", handleLikesOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/likes/:user_id", ctrl.MuxHandler("preflight", handleLikesOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateLikesContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateLikesPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleLikesOrigin(h)
	service.Mux.Handle("POST", "/likes", ctrl.MuxHandler("create", h, unmarshalCreateLikesPayload))
	service.LogInfo("mount", "ctrl", "Likes", "action", "Create", "route", "POST /likes", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteLikesContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*DeleteLikesPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleLikesOrigin(h)
	service.Mux.Handle("DELETE", "/likes", ctrl.MuxHandler("delete", h, unmarshalDeleteLikesPayload))
	service.LogInfo("mount", "ctrl", "Likes", "action", "Delete", "route", "DELETE /likes", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetLikeByUserLikesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetLikeByUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleLikesOrigin(h)
	service.Mux.Handle("GET", "/likes/:user_id", ctrl.MuxHandler("get_like_by_user", h, nil))
	service.LogInfo("mount", "ctrl", "Likes", "action", "GetLikeByUser", "route", "GET /likes/:user_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMyLikeLikesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetMyLike(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleLikesOrigin(h)
	service.Mux.Handle("GET", "/likes", ctrl.MuxHandler("get_my_like", h, nil))
	service.LogInfo("mount", "ctrl", "Likes", "action", "GetMyLike", "route", "GET /likes", "security", "jwt")
}

// handleLikesOrigin applies the CORS response headers corresponding to the origin.
func handleLikesOrigin(h goa.Handler) goa.Handler {
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

// unmarshalCreateLikesPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateLikesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createLikesPayload{}
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

// unmarshalDeleteLikesPayload unmarshals the request body into the context request data Payload field.
func unmarshalDeleteLikesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &deleteLikesPayload{}
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

// PostsController is the controller interface for the Posts actions.
type PostsController interface {
	goa.Muxer
	CreatePost(*CreatePostPostsContext) error
	Delete(*DeletePostsContext) error
	Index(*IndexPostsContext) error
	Show(*ShowPostsContext) error
	ShowMyLike(*ShowMyLikePostsContext) error
	ShowPostLike(*ShowPostLikePostsContext) error
	ShowPostMedia(*ShowPostMediaPostsContext) error
	ShowPostMy(*ShowPostMyPostsContext) error
	Update(*UpdatePostsContext) error
}

// MountPostsController "mounts" a Posts resource controller on the given service.
func MountPostsController(service *goa.Service, ctrl PostsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/posts", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/posts/:id", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/posts/my_like", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/posts/likes/:id", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/posts/my_media/:id", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/posts/my_post/:id", ctrl.MuxHandler("preflight", handlePostsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePostPostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreatePostPostsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreatePost(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("POST", "/posts", ctrl.MuxHandler("create_post", h, unmarshalCreatePostPostsPayload))
	service.LogInfo("mount", "ctrl", "Posts", "action", "CreatePost", "route", "POST /posts", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeletePostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("DELETE", "/posts/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "Delete", "route", "DELETE /posts/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewIndexPostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Index(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts", ctrl.MuxHandler("index", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "Index", "route", "GET /posts", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "Show", "route", "GET /posts/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowMyLikePostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowMyLike(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts/my_like", ctrl.MuxHandler("show_my_like", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "ShowMyLike", "route", "GET /posts/my_like", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPostLikePostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowPostLike(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts/likes/:id", ctrl.MuxHandler("show_post_like", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "ShowPostLike", "route", "GET /posts/likes/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPostMediaPostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowPostMedia(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts/my_media/:id", ctrl.MuxHandler("show_post_media", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "ShowPostMedia", "route", "GET /posts/my_media/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPostMyPostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowPostMy(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("GET", "/posts/my_post/:id", ctrl.MuxHandler("show_post_my", h, nil))
	service.LogInfo("mount", "ctrl", "Posts", "action", "ShowPostMy", "route", "GET /posts/my_post/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdatePostsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdatePostsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handlePostsOrigin(h)
	service.Mux.Handle("PUT", "/posts/:id", ctrl.MuxHandler("update", h, unmarshalUpdatePostsPayload))
	service.LogInfo("mount", "ctrl", "Posts", "action", "Update", "route", "PUT /posts/:id", "security", "jwt")
}

// handlePostsOrigin applies the CORS response headers corresponding to the origin.
func handlePostsOrigin(h goa.Handler) goa.Handler {
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

// unmarshalCreatePostPostsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePostPostsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createPostPostsPayload{}
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

// unmarshalUpdatePostsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdatePostsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updatePostsPayload{}
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

// RoomsController is the controller interface for the Rooms actions.
type RoomsController interface {
	goa.Muxer
	CreateRoom(*CreateRoomRoomsContext) error
	Exists(*ExistsRoomsContext) error
	Index(*IndexRoomsContext) error
}

// MountRoomsController "mounts" a Rooms resource controller on the given service.
func MountRoomsController(service *goa.Service, ctrl RoomsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/rooms", ctrl.MuxHandler("preflight", handleRoomsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/rooms/exists", ctrl.MuxHandler("preflight", handleRoomsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateRoomRoomsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateRoomRoomsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateRoom(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleRoomsOrigin(h)
	service.Mux.Handle("POST", "/rooms", ctrl.MuxHandler("create_room", h, unmarshalCreateRoomRoomsPayload))
	service.LogInfo("mount", "ctrl", "Rooms", "action", "CreateRoom", "route", "POST /rooms", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewExistsRoomsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Exists(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleRoomsOrigin(h)
	service.Mux.Handle("GET", "/rooms/exists", ctrl.MuxHandler("exists", h, nil))
	service.LogInfo("mount", "ctrl", "Rooms", "action", "Exists", "route", "GET /rooms/exists", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewIndexRoomsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Index(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleRoomsOrigin(h)
	service.Mux.Handle("GET", "/rooms", ctrl.MuxHandler("index", h, nil))
	service.LogInfo("mount", "ctrl", "Rooms", "action", "Index", "route", "GET /rooms", "security", "jwt")
}

// handleRoomsOrigin applies the CORS response headers corresponding to the origin.
func handleRoomsOrigin(h goa.Handler) goa.Handler {
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

// unmarshalCreateRoomRoomsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRoomRoomsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createRoomRoomsPayload{}
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

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	GetCurrentUser(*GetCurrentUserUsersContext) error
	ShowUser(*ShowUserUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/current_user", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/:id", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

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

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUserUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/users/:id", ctrl.MuxHandler("show_user", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "ShowUser", "route", "GET /users/:id", "security", "jwt")
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
