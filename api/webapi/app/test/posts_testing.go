// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": posts TestHelpers
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=$(GOPATH)/src/app/webapi
// --version=v1.5.13

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
	"strconv"
	"strings"
	"testing"
)

// CreatePostPostsBadRequest runs the method CreatePost of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreatePostPostsBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, payload *app.CreatePostPostsPayload) (http.ResponseWriter, *app.ServiceVerror) {
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
		Path: fmt.Sprintf("/posts"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	createPostCtx, err := app.NewCreatePostPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}
	createPostCtx.Payload = payload

	// Perform action
	err = ctrl.CreatePost(createPostCtx)

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

// CreatePostPostsCreated runs the method CreatePost of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreatePostPostsCreated(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, payload *app.CreatePostPostsPayload) (http.ResponseWriter, *app.IndexPostJSON) {
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
		Path: fmt.Sprintf("/posts"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	createPostCtx, err := app.NewCreatePostPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}
	createPostCtx.Payload = payload

	// Perform action
	err = ctrl.CreatePost(createPostCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.IndexPostJSON
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.IndexPostJSON)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.IndexPostJSON", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// CreatePostPostsInternalServerError runs the method CreatePost of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreatePostPostsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, payload *app.CreatePostPostsPayload) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	createPostCtx, err := app.NewCreatePostPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	createPostCtx.Payload = payload

	// Perform action
	err = ctrl.CreatePost(createPostCtx)

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

// DeletePostsInternalServerError runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeletePostsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("DELETE", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	deleteCtx, err := app.NewDeletePostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

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

// DeletePostsOK runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeletePostsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("DELETE", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	deleteCtx, err := app.NewDeletePostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

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

// IndexPostsInternalServerError runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexPostsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, nextID *int) http.ResponseWriter {
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
	query := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		query["next_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/posts"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Index(indexCtx)

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

// IndexPostsNotFound runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexPostsNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, nextID *int) http.ResponseWriter {
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
	query := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		query["next_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/posts"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Index(indexCtx)

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

// IndexPostsOK runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexPostsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, nextID *int) (http.ResponseWriter, *app.PostAllLimit) {
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
	query := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		query["next_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/posts"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	err = ctrl.Index(indexCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.PostAllLimit
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.PostAllLimit)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.PostAllLimit", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowPostsInternalServerError runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPostsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	showCtx, err := app.NewShowPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Show(showCtx)

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

// ShowPostsNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPostsNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	showCtx, err := app.NewShowPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Show(showCtx)

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

// ShowPostsOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPostsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int) (http.ResponseWriter, *app.ShowPostJSON) {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	showCtx, err := app.NewShowPostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ShowPostJSON
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ShowPostJSON)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ShowPostJSON", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// UpdatePostsBadRequest runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePostsBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int, payload *app.UpdatePostsPayload) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("PUT", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	updateCtx, err := app.NewUpdatePostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

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

// UpdatePostsInternalServerError runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePostsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int, payload *app.UpdatePostsPayload) http.ResponseWriter {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("PUT", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	updateCtx, err := app.NewUpdatePostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

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

// UpdatePostsOK runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePostsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.PostsController, id int, payload *app.UpdatePostsPayload) (http.ResponseWriter, *app.IndexPostJSON) {
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
		Path: fmt.Sprintf("/posts/%v", id),
	}
	req := httptest.NewRequest("PUT", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostsTest"), rw, req, prms)
	updateCtx, err := app.NewUpdatePostsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.IndexPostJSON
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.IndexPostJSON)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.IndexPostJSON", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
