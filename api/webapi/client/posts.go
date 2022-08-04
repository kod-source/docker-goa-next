// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": posts Resource Client
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=$(GOPATH)/src/app/webapi
// --version=v1.5.13

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreatePostPostsPayload is the posts create_post action payload.
type CreatePostPostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// CreatePostPostsPath computes a request path to the create_post action of posts.
func CreatePostPostsPath() string {
	return fmt.Sprintf("/posts")
}

// 投稿を作成する
func (c *Client) CreatePostPosts(ctx context.Context, path string, payload *CreatePostPostsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePostPostsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePostPostsRequest create the request corresponding to the create_post action endpoint of the posts resource.
func (c *Client) NewCreatePostPostsRequest(ctx context.Context, path string, payload *CreatePostPostsPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
