// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/kod-source/docker-goa-next/webapi/design
// --out=/Users/horikoudai/Documents/ProgrammingLearning/docker-goa-next/api/webapi
// --version=v1.5.13

package app

import (
	goa "github.com/shogo82148/goa-v1"
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

// Index_post_jsonCollection is the media type for an array of Index_post_json (default view)
//
// Identifier: application/vnd.index_post_json; type=collection; view=default
type IndexPostJSONCollection []*IndexPostJSON

// Validate validates the IndexPostJSONCollection media type instance.
func (mt IndexPostJSONCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// いいね (default view)
//
// Identifier: application/vnd.like_json; view=default
type LikeJSON struct {
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 投稿ID
	PostID int `form:"post_id" json:"post_id" yaml:"post_id" xml:"post_id"`
	// ユーザーID
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the LikeJSON media type instance.
func (mt *LikeJSON) Validate() (err error) {

	return
}

// Like_jsonCollection is the media type for an array of Like_json (default view)
//
// Identifier: application/vnd.like_json; type=collection; view=default
type LikeJSONCollection []*LikeJSON

// Validate validates the LikeJSONCollection media type instance.
func (mt LikeJSONCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// 投稿とnext_idに情報 (default view)
//
// Identifier: application/vnd.post_all_limit; view=default
type PostAllLimit struct {
	// http://localhost:3000/posts?next_id=20
	NextToken *string `form:"next_token,omitempty" json:"next_token,omitempty" yaml:"next_token,omitempty" xml:"next_token,omitempty"`
	// post_and_user vbalue
	ShowPosts IndexPostJSONCollection `form:"show_posts" json:"show_posts" yaml:"show_posts" xml:"show_posts"`
}

// Validate validates the PostAllLimit media type instance.
func (mt *PostAllLimit) Validate() (err error) {
	if mt.ShowPosts == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "show_posts"))
	}
	if err2 := mt.ShowPosts.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
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

// 投稿とユーザーとコメントの情報 (default view)
//
// Identifier: application/vnd.show_post_json; view=default
type ShowPostJSON struct {
	// comments value
	Comments CommentJSONCollection `form:"comments" json:"comments" yaml:"comments" xml:"comments"`
	// likes value
	Likes LikeJSONCollection `form:"likes" json:"likes" yaml:"likes" xml:"likes"`
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
	if mt.Likes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "likes"))
	}
	if err2 := mt.Comments.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Likes.Validate(); err2 != nil {
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
