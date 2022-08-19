// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=$(GOPATH)/src/app/webapi
// --version=v1.5.13

package client

import (
	goa "github.com/shogo82148/goa-v1"
	"net/http"
	"time"
)

// コメント (default view)
//
// Identifier: application/vnd.comment_json; view=default
type CommentJSON struct {
	// 作成日
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// コメントの画像パス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// 投稿ID
	PostID int `form:"post_id" json:"post_id" yaml:"post_id" xml:"post_id"`
	// コメント
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
	// 更新日
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the CommentJSON media type instance.
func (mt *CommentJSON) Validate() (err error) {

	return
}

// DecodeCommentJSON decodes the CommentJSON instance encoded in resp body.
func (c *Client) DecodeCommentJSON(resp *http.Response) (*CommentJSON, error) {
	var decoded CommentJSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Comment_jsonCollection is the media type for an array of Comment_json (default view)
//
// Identifier: application/vnd.comment_json; type=collection; view=default
type CommentJSONCollection []*CommentJSON

// Validate validates the CommentJSONCollection media type instance.
func (mt CommentJSONCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCommentJSONCollection decodes the CommentJSONCollection instance encoded in resp body.
func (c *Client) DecodeCommentJSONCollection(resp *http.Response) (CommentJSONCollection, error) {
	var decoded CommentJSONCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// 投稿 (default view)
//
// Identifier: application/vnd.index_post_json; view=default
type IndexPostJSON struct {
	// ユーザー名
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// post value
	Post *PostJSON `form:"post" json:"post" yaml:"post" xml:"post"`
	// ユーザー名
	UserName string `form:"user_name" json:"user_name" yaml:"user_name" xml:"user_name"`
}

// Validate validates the IndexPostJSON media type instance.
func (mt *IndexPostJSON) Validate() (err error) {
	if mt.Post == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "post"))
	}

	if mt.Post != nil {
		if err2 := mt.Post.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeIndexPostJSON decodes the IndexPostJSON instance encoded in resp body.
func (c *Client) DecodeIndexPostJSON(resp *http.Response) (*IndexPostJSON, error) {
	var decoded IndexPostJSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// 投稿とnext_idに情報 (default view)
//
// Identifier: application/vnd.post_all_limit; view=default
type PostAllLimit struct {
	// http://localhost:3000/posts?next_id=20
	NextToken *string `form:"next_token,omitempty" json:"next_token,omitempty" yaml:"next_token,omitempty" xml:"next_token,omitempty"`
	// post_and_user vbalue
	PostAndUser *IndexPostJSON `form:"post_and_user" json:"post_and_user" yaml:"post_and_user" xml:"post_and_user"`
}

// Validate validates the PostAllLimit media type instance.
func (mt *PostAllLimit) Validate() (err error) {
	if mt.PostAndUser == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "post_and_user"))
	}
	if mt.PostAndUser != nil {
		if err2 := mt.PostAndUser.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodePostAllLimit decodes the PostAllLimit instance encoded in resp body.
func (c *Client) DecodePostAllLimit(resp *http.Response) (*PostAllLimit, error) {
	var decoded PostAllLimit
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Post_all_limitCollection is the media type for an array of Post_all_limit (default view)
//
// Identifier: application/vnd.post_all_limit; type=collection; view=default
type PostAllLimitCollection []*PostAllLimit

// Validate validates the PostAllLimitCollection media type instance.
func (mt PostAllLimitCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePostAllLimitCollection decodes the PostAllLimitCollection instance encoded in resp body.
func (c *Client) DecodePostAllLimitCollection(resp *http.Response) (PostAllLimitCollection, error) {
	var decoded PostAllLimitCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// 投稿とユーザーの情報 (default view)
//
// Identifier: application/vnd.post_and_user_json; view=default
type PostAndUserJSON struct {
	// post value
	Post *PostJSON `form:"post" json:"post" yaml:"post" xml:"post"`
	// user value
	User *User `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the PostAndUserJSON media type instance.
func (mt *PostAndUserJSON) Validate() (err error) {
	if mt.Post == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "post"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Post != nil {
		if err2 := mt.Post.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodePostAndUserJSON decodes the PostAndUserJSON instance encoded in resp body.
func (c *Client) DecodePostAndUserJSON(resp *http.Response) (*PostAndUserJSON, error) {
	var decoded PostAndUserJSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// 投稿 (default view)
//
// Identifier: application/vnd.post_json; view=default
type PostJSON struct {
	// 作成日
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// プロフィール画像のパス
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// タイトル
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// 更新日
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// ユーザーID
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the PostJSON media type instance.
func (mt *PostJSON) Validate() (err error) {

	return
}

// DecodePostJSON decodes the PostJSON instance encoded in resp body.
func (c *Client) DecodePostJSON(resp *http.Response) (*PostJSON, error) {
	var decoded PostJSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// my error (default view)
//
// Identifier: application/vnd.service.verror; view=default
type ServiceVerror struct {
	// Code
	Code int `form:"code" json:"code" yaml:"code" xml:"code"`
	// Details
	Details interface{} `form:"details,omitempty" json:"details,omitempty" yaml:"details,omitempty" xml:"details,omitempty"`
	// エラーメッセージ
	Message string `form:"message" json:"message" yaml:"message" xml:"message"`
	// Status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
}

// Validate validates the ServiceVerror media type instance.
func (mt *ServiceVerror) Validate() (err error) {

	return
}

// DecodeServiceVerror decodes the ServiceVerror instance encoded in resp body.
func (c *Client) DecodeServiceVerror(resp *http.Response) (*ServiceVerror, error) {
	var decoded ServiceVerror
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// 投稿とユーザーとコメントの情報 (default view)
//
// Identifier: application/vnd.show_post_json; view=default
type ShowPostJSON struct {
	// comments value
	Comments CommentJSONCollection `form:"comments" json:"comments" yaml:"comments" xml:"comments"`
	// post value
	Post *PostJSON `form:"post" json:"post" yaml:"post" xml:"post"`
	// user value
	User *User `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the ShowPostJSON media type instance.
func (mt *ShowPostJSON) Validate() (err error) {
	if mt.Post == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "post"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Comments == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "comments"))
	}
	if err2 := mt.Comments.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Post != nil {
		if err2 := mt.Post.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeShowPostJSON decodes the ShowPostJSON instance encoded in resp body.
func (c *Client) DecodeShowPostJSON(resp *http.Response) (*ShowPostJSON, error) {
	var decoded ShowPostJSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// token (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// token value
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
	// user value
	User *User `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the Token media type instance.
func (mt *Token) Validate() (err error) {

	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeToken decodes the Token instance encoded in resp body.
func (c *Client) DecodeToken(resp *http.Response) (*Token, error) {
	var decoded Token
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// プロフィール画像パス
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// 作成日
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// メール
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 名前
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
