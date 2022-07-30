// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": operands Resource Client
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=$(GOPATH)/src/app/webapi
// --version=v1.5.13

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AddOperandsPath computes a request path to the add action of operands.
func AddOperandsPath(left int, right int) string {
	param0 := strconv.Itoa(left)
	param1 := strconv.Itoa(right)
	return fmt.Sprintf("/add/%s/%s", param0, param1)
}

// add returns the sum of the left and right parameters in the response body
func (c *Client) AddOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddOperandsRequest create the request corresponding to the add action endpoint of the operands resource.
func (c *Client) NewAddOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
