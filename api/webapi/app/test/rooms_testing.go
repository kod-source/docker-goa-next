// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": rooms TestHelpers
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
	"strconv"
	"strings"
	"testing"
)

// CreateRoomRoomsBadRequest runs the method CreateRoom of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateRoomRoomsBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, payload *app.CreateRoomRoomsPayload) (http.ResponseWriter, *app.ServiceVerror) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/rooms"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	createRoomCtx, _err := app.NewCreateRoomRoomsContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	createRoomCtx.Payload = payload

	// Perform action
	_err = ctrl.CreateRoom(createRoomCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.ServiceVerror
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.ServiceVerror)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ServiceVerror", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// CreateRoomRoomsCreated runs the method CreateRoom of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateRoomRoomsCreated(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, payload *app.CreateRoomRoomsPayload) (http.ResponseWriter, *app.RoomUser) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/rooms"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	createRoomCtx, _err := app.NewCreateRoomRoomsContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	createRoomCtx.Payload = payload

	// Perform action
	_err = ctrl.CreateRoom(createRoomCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.RoomUser
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.RoomUser)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.RoomUser", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// CreateRoomRoomsInternalServerError runs the method CreateRoom of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateRoomRoomsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, payload *app.CreateRoomRoomsPayload) http.ResponseWriter {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/rooms"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	createRoomCtx, _err := app.NewCreateRoomRoomsContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil
	}
	createRoomCtx.Payload = payload

	// Perform action
	_err = ctrl.CreateRoom(createRoomCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ExistsRoomsInternalServerError runs the method Exists of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ExistsRoomsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, userID int) http.ResponseWriter {
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
	{
		sliceVal := []string{strconv.Itoa(userID)}
		query["user_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/rooms/exists"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(userID)}
		prms["user_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	existsCtx, err := app.NewExistsRoomsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Exists(existsCtx)

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

// ExistsRoomsNotFound runs the method Exists of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ExistsRoomsNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, userID int) http.ResponseWriter {
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
	{
		sliceVal := []string{strconv.Itoa(userID)}
		query["user_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/rooms/exists"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(userID)}
		prms["user_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	existsCtx, err := app.NewExistsRoomsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	err = ctrl.Exists(existsCtx)

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

// ExistsRoomsOK runs the method Exists of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ExistsRoomsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, userID int) (http.ResponseWriter, *app.Room) {
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
	{
		sliceVal := []string{strconv.Itoa(userID)}
		query["user_id"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/rooms/exists"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(userID)}
		prms["user_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	existsCtx, err := app.NewExistsRoomsContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	err = ctrl.Exists(existsCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Room
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Room)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Room", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// IndexRoomsInternalServerError runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexRoomsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, nextID *int) http.ResponseWriter {
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
		Path:     fmt.Sprintf("/rooms"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexRoomsContext(goaCtx, req, service)
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

// IndexRoomsNotFound runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexRoomsNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, nextID *int) http.ResponseWriter {
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
		Path:     fmt.Sprintf("/rooms"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexRoomsContext(goaCtx, req, service)
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

// IndexRoomsOK runs the method Index of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IndexRoomsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, nextID *int) (http.ResponseWriter, *app.AllRoomUser) {
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
		Path:     fmt.Sprintf("/rooms"),
		RawQuery: query.Encode(),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	if nextID != nil {
		sliceVal := []string{strconv.Itoa(*nextID)}
		prms["next_id"] = sliceVal
	}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	indexCtx, err := app.NewIndexRoomsContext(goaCtx, req, service)
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
	var mt *app.AllRoomUser
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.AllRoomUser)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.AllRoomUser", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowRoomsBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowRoomsBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/rooms/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	showCtx, err := app.NewShowRoomsContext(goaCtx, req, service)
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
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}

	// Return results
	return rw
}

// ShowRoomsInternalServerError runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowRoomsInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/rooms/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	showCtx, err := app.NewShowRoomsContext(goaCtx, req, service)
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

// ShowRoomsNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowRoomsNotFound(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, id int) http.ResponseWriter {
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
		Path: fmt.Sprintf("/rooms/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	showCtx, err := app.NewShowRoomsContext(goaCtx, req, service)
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

// ShowRoomsOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowRoomsOK(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.RoomsController, id int) (http.ResponseWriter, *app.RoomUser) {
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
		Path: fmt.Sprintf("/rooms/%v", id),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "RoomsTest"), rw, req, prms)
	showCtx, err := app.NewShowRoomsContext(goaCtx, req, service)
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
	var mt *app.RoomUser
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.RoomUser)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.RoomUser", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
