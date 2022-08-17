// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": comments Resource Client
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

// CreateCommentCommentsPayload is the comments create_comment action payload.
type CreateCommentCommentsPayload struct {
	// コメント画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// コメントの内容
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
}

// CreateCommentCommentsPath computes a request path to the create_comment action of comments.
func CreateCommentCommentsPath() string {
	return fmt.Sprintf("/comments")
}

// コメントの作成
func (c *Client) CreateCommentComments(ctx context.Context, path string, payload *CreateCommentCommentsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCommentCommentsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCommentCommentsRequest create the request corresponding to the create_comment action endpoint of the comments resource.
func (c *Client) NewCreateCommentCommentsRequest(ctx context.Context, path string, payload *CreateCommentCommentsPayload, contentType string) (*http.Request, error) {
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
