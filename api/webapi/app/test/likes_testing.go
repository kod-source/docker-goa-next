// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": likes TestHelpers
//
// Command:
// $ main

package test

import (
	"context"
	"fmt"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// CreateLikesBadRequest runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateLikesBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.CreateLikesPayload) (http.ResponseWriter, *app.ServiceVerror) {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	createCtx, err := app.NewCreateLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}
	createCtx.Payload = payload

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.ServiceVerror
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ServiceVerror)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ServiceVerror", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CreateLikesCreated runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateLikesCreated(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.CreateLikesPayload) (http.ResponseWriter, *app.LikeJSON) {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	createCtx, err := app.NewCreateLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}
	createCtx.Payload = payload

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.LikeJSON
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.LikeJSON)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.LikeJSON", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CreateLikesInternalServerError runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateLikesInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.CreateLikesPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	createCtx, err := app.NewCreateLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	createCtx.Payload = payload

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// DeleteLikesBadRequest runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteLikesBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.DeleteLikesPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("DELETE", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	deleteCtx.Payload = payload

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}

	// Return results
	return rw
}

// DeleteLikesInternalServerError runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteLikesInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.DeleteLikesPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("DELETE", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	deleteCtx.Payload = payload

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// DeleteLikesOK runs the method Delete of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteLikesOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, payload *app.DeleteLikesPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("DELETE", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	deleteCtx.Payload = payload

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// GetLikeByUserLikesInternalServerError runs the method GetLikeByUser of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetLikeByUserLikesInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, userID int) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes/%v", userID),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["user_id"] = []string{fmt.Sprintf("%v", userID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	getLikeByUserCtx, err := app.NewGetLikeByUserLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.GetLikeByUser(getLikeByUserCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// GetLikeByUserLikesNotFound runs the method GetLikeByUser of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetLikeByUserLikesNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, userID int) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes/%v", userID),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["user_id"] = []string{fmt.Sprintf("%v", userID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	getLikeByUserCtx, err := app.NewGetLikeByUserLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.GetLikeByUser(getLikeByUserCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetLikeByUserLikesOK runs the method GetLikeByUser of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetLikeByUserLikesOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController, userID int) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes/%v", userID),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["user_id"] = []string{fmt.Sprintf("%v", userID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	getLikeByUserCtx, err := app.NewGetLikeByUserLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.GetLikeByUser(getLikeByUserCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// GetMyLikeLikesInternalServerError runs the method GetMyLike of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetMyLikeLikesInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	getMyLikeCtx, err := app.NewGetMyLikeLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.GetMyLike(getMyLikeCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// GetMyLikeLikesOK runs the method GetMyLike of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetMyLikeLikesOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.LikesController) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/likes"),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "LikesTest"), rw, req, prms)
	getMyLikeCtx, err := app.NewGetMyLikeLikesContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.GetMyLike(getMyLikeCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
