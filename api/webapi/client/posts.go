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
	"strconv"
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

// DeletePostsPath computes a request path to the delete action of posts.
func DeletePostsPath(id int) string {
	param0 := strconv.Itoa(id)
	return fmt.Sprintf("/posts/%s", param0)
}

// 投稿を削除する
func (c *Client) DeletePosts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePostsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePostsRequest create the request corresponding to the delete action endpoint of the posts resource.
func (c *Client) NewDeletePostsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
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

// IndexPostsPath computes a request path to the index action of posts.
func IndexPostsPath() string {
	return fmt.Sprintf("/posts")
}

// 全部の登録を取得する
func (c *Client) IndexPosts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIndexPostsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexPostsRequest create the request corresponding to the index action endpoint of the posts resource.
func (c *Client) NewIndexPostsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowPostsPath computes a request path to the show action of posts.
func ShowPostsPath(id int) string {
	param0 := strconv.Itoa(id)
	return fmt.Sprintf("/posts/%s", param0)
}

// 一つの投稿を取得する
func (c *Client) ShowPosts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPostsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPostsRequest create the request corresponding to the show action endpoint of the posts resource.
func (c *Client) NewShowPostsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdatePostsPayload is the posts update action payload.
type UpdatePostsPayload struct {
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// UpdatePostsPath computes a request path to the update action of posts.
func UpdatePostsPath(id int) string {
	param0 := strconv.Itoa(id)
	return fmt.Sprintf("/posts/%s", param0)
}

// 投稿を更新する
func (c *Client) UpdatePosts(ctx context.Context, path string, payload *UpdatePostsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdatePostsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdatePostsRequest create the request corresponding to the update action endpoint of the posts resource.
func (c *Client) NewUpdatePostsRequest(ctx context.Context, path string, payload *UpdatePostsPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), &body)
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
