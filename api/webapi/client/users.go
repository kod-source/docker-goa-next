// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": users Resource Client
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=/Users/horikoudai/Documents/ProgrammingLearning/docker-goa-next/api/webapi
// --version=v1.5.13

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetCurrentUserUsersPath computes a request path to the get_current_user action of users.
func GetCurrentUserUsersPath() string {
	return fmt.Sprintf("/current_user")
}

// ログインしているユーザーの情報を取得する
func (c *Client) GetCurrentUserUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCurrentUserUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCurrentUserUsersRequest create the request corresponding to the get_current_user action endpoint of the users resource.
func (c *Client) NewGetCurrentUserUsersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
