// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "docker_goa_next": Application Media Types
//
// Command:
// $ main

package app

import (
	goa "github.com/shogo82148/goa-v1"
	"time"
)

// スレッド一覧とNextID (default view)
//
// Identifier: application/vnd.all_index_threads; view=default
type AllIndexThreads struct {
	// index_threads
	IndexThreads IndexThreadCollection `form:"index_threads" json:"index_threads" yaml:"index_threads" xml:"index_threads"`
	// 次取得するThreadID
	NextID *int `form:"next_id,omitempty" json:"next_id,omitempty" yaml:"next_id,omitempty" xml:"next_id,omitempty"`
}

// Validate validates the AllIndexThreads media type instance.
func (mt *AllIndexThreads) Validate() (err error) {
	if mt.IndexThreads == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "index_threads"))
	}
	if err2 := mt.IndexThreads.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// 全てのルーム (default view)
//
// Identifier: application/vnd.all_room_user; view=default
type AllRoomUser struct {
	// index_rooms
	IndexRoom IndexRoomCollection `form:"index_room" json:"index_room" yaml:"index_room" xml:"index_room"`
	// 次取得するRoomID
	NextID *int `form:"next_id,omitempty" json:"next_id,omitempty" yaml:"next_id,omitempty" xml:"next_id,omitempty"`
}

// Validate validates the AllRoomUser media type instance.
func (mt *AllRoomUser) Validate() (err error) {
	if mt.IndexRoom == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "index_room"))
	}
	if err2 := mt.IndexRoom.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

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
	// ユーザーID
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the CommentJSON media type instance.
func (mt *CommentJSON) Validate() (err error) {

	return
}

// コメントとユーザーの情報 (default view)
//
// Identifier: application/vnd.comment_with_user_json; view=default
type CommentWithUserJSON struct {
	// コメント
	Comment *CommentJSON `form:"comment" json:"comment" yaml:"comment" xml:"comment"`
	// ユーザー
	User *User `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the CommentWithUserJSON media type instance.
func (mt *CommentWithUserJSON) Validate() (err error) {
	if mt.Comment == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "comment"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Comment != nil {
		if err2 := mt.Comment.Validate(); err2 != nil {
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

// Comment_with_user_jsonCollection is the media type for an array of Comment_with_user_json (default view)
//
// Identifier: application/vnd.comment_with_user_json; type=collection; view=default
type CommentWithUserJSONCollection []*CommentWithUserJSON

// Validate validates the CommentWithUserJSONCollection media type instance.
func (mt CommentWithUserJSONCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// コンテント (default view)
//
// Identifier: application/vnd.content; view=default
type Content struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// content id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 画像
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// コンテント内容
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
	// thread id
	ThreadID int `form:"thread_id" json:"thread_id" yaml:"thread_id" xml:"thread_id"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
	// user id
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the Content media type instance.
func (mt *Content) Validate() (err error) {

	return
}

// コンテントとユーザー (default view)
//
// Identifier: application/vnd.content_user; view=default
type ContentUser struct {
	Content *Content  `form:"content" json:"content" yaml:"content" xml:"content"`
	User    *ShowUser `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the ContentUser media type instance.
func (mt *ContentUser) Validate() (err error) {
	if mt.Content == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "content"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Content != nil {
		if err2 := mt.Content.Validate(); err2 != nil {
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

// ルームの表示 (default view)
//
// Identifier: application/vnd.index_room; view=default
type IndexRoom struct {
	// ルームに入っているユーザー数
	CountUser int `form:"count_user" json:"count_user" yaml:"count_user" xml:"count_user"`
	// 開いたどうか
	IsOpen bool `form:"is_open" json:"is_open" yaml:"is_open" xml:"is_open"`
	// 最後の内容
	LastText *string `form:"last_text,omitempty" json:"last_text,omitempty" yaml:"last_text,omitempty" xml:"last_text,omitempty"`
	// room
	Room *Room `form:"room" json:"room" yaml:"room" xml:"room"`
	// DMの際に相手のプロフィール画像のパス
	ShowImg *string `form:"show_img,omitempty" json:"show_img,omitempty" yaml:"show_img,omitempty" xml:"show_img,omitempty"`
}

// Validate validates the IndexRoom media type instance.
func (mt *IndexRoom) Validate() (err error) {
	if mt.Room == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "room"))
	}

	if mt.Room != nil {
		if err2 := mt.Room.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Index_roomCollection is the media type for an array of Index_room (default view)
//
// Identifier: application/vnd.index_room; type=collection; view=default
type IndexRoomCollection []*IndexRoom

// Validate validates the IndexRoomCollection media type instance.
func (mt IndexRoomCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// スレッドの一覧 (default view)
//
// Identifier: application/vnd.index_thread; view=default
type IndexThread struct {
	// スレッドの返信数
	CountContent *int `form:"count_content,omitempty" json:"count_content,omitempty" yaml:"count_content,omitempty" xml:"count_content,omitempty"`
	// スレッドとユーザー
	ThreadUser *ThreadUser `form:"thread_user" json:"thread_user" yaml:"thread_user" xml:"thread_user"`
}

// Validate validates the IndexThread media type instance.
func (mt *IndexThread) Validate() (err error) {
	if mt.ThreadUser == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thread_user"))
	}
	if mt.ThreadUser != nil {
		if err2 := mt.ThreadUser.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Index_threadCollection is the media type for an array of Index_thread (default view)
//
// Identifier: application/vnd.index_thread; type=collection; view=default
type IndexThreadCollection []*IndexThread

// Validate validates the IndexThreadCollection media type instance.
func (mt IndexThreadCollection) Validate() (err error) {
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
	// 次取得するPostID
	NextID *int `form:"next_id,omitempty" json:"next_id,omitempty" yaml:"next_id,omitempty" xml:"next_id,omitempty"`
	// post_and_user vbalue
	ShowPosts PostAndUserAndCountLikeJSONCollection `form:"show_posts" json:"show_posts" yaml:"show_posts" xml:"show_posts"`
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

// 投稿といいね数 (default view)
//
// Identifier: application/vnd.post_and_user_and_count_like_json; view=default
type PostAndUserAndCountLikeJSON struct {
	// ユーザー名
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// コメント数
	CountComment int `form:"count_comment" json:"count_comment" yaml:"count_comment" xml:"count_comment"`
	// いいね数
	CountLike int `form:"count_like" json:"count_like" yaml:"count_like" xml:"count_like"`
	// post value
	Post *PostJSON `form:"post" json:"post" yaml:"post" xml:"post"`
	// ユーザー名
	UserName string `form:"user_name" json:"user_name" yaml:"user_name" xml:"user_name"`
}

// Validate validates the PostAndUserAndCountLikeJSON media type instance.
func (mt *PostAndUserAndCountLikeJSON) Validate() (err error) {
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

// Post_and_user_and_count_like_jsonCollection is the media type for an array of Post_and_user_and_count_like_json (default view)
//
// Identifier: application/vnd.post_and_user_and_count_like_json; type=collection; view=default
type PostAndUserAndCountLikeJSONCollection []*PostAndUserAndCountLikeJSON

// Validate validates the PostAndUserAndCountLikeJSONCollection media type instance.
func (mt PostAndUserAndCountLikeJSONCollection) Validate() (err error) {
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

// ルーム (default view)
//
// Identifier: application/vnd.room; view=default
type Room struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// room id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 画像
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// グループかDMの判定
	IsGroup bool `form:"is_group" json:"is_group" yaml:"is_group" xml:"is_group"`
	// room name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the Room media type instance.
func (mt *Room) Validate() (err error) {

	return
}

// ルーム (default view)
//
// Identifier: application/vnd.room_user; view=default
type RoomUser struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// room id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 画像
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// グループかDMの判定
	IsGroup bool `form:"is_group" json:"is_group" yaml:"is_group" xml:"is_group"`
	// room name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
	// ルームいるユーザー
	Users ShowUserCollection `form:"users" json:"users" yaml:"users" xml:"users"`
}

// Validate validates the RoomUser media type instance.
func (mt *RoomUser) Validate() (err error) {

	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
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
	CommentsWithUsers CommentWithUserJSONCollection `form:"comments_with_users" json:"comments_with_users" yaml:"comments_with_users" xml:"comments_with_users"`
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
	if mt.CommentsWithUsers == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "comments_with_users"))
	}
	if mt.Likes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "likes"))
	}
	if err2 := mt.CommentsWithUsers.Validate(); err2 != nil {
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

// show_user (default view)
//
// Identifier: application/vnd.show_user+json; view=default
type ShowUser struct {
	// プロフィール画像パス
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" yaml:"avatar,omitempty" xml:"avatar,omitempty"`
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the ShowUser media type instance.
func (mt *ShowUser) Validate() (err error) {

	return
}

// Show_userCollection is the media type for an array of Show_user (default view)
//
// Identifier: application/vnd.show_user+json; type=collection; view=default
type ShowUserCollection []*ShowUser

// Validate validates the ShowUserCollection media type instance.
func (mt ShowUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// スレッド (default view)
//
// Identifier: application/vnd.thread; view=default
type Thread struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// room id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 画像
	Img *string `form:"img,omitempty" json:"img,omitempty" yaml:"img,omitempty" xml:"img,omitempty"`
	// room id
	RoomID int `form:"room_id" json:"room_id" yaml:"room_id" xml:"room_id"`
	// スレッド内容
	Text string `form:"text" json:"text" yaml:"text" xml:"text"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
	// user id
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the Thread media type instance.
func (mt *Thread) Validate() (err error) {

	return
}

// スレッドとユーザー情報 (default view)
//
// Identifier: application/vnd.thread_user; view=default
type ThreadUser struct {
	// スレッド
	Thread *Thread `form:"thread" json:"thread" yaml:"thread" xml:"thread"`
	// ユーザー
	User *ShowUser `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the ThreadUser media type instance.
func (mt *ThreadUser) Validate() (err error) {
	if mt.Thread == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thread"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Thread != nil {
		if err2 := mt.Thread.Validate(); err2 != nil {
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

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {

	return
}

// user room (default view)
//
// Identifier: application/vnd.user_room+json; view=default
type UserRoom struct {
	// 作成日
	CreatedAt time.Time `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// 最後に開いた日時
	LastReadAt *time.Time `form:"last_read_at,omitempty" json:"last_read_at,omitempty" yaml:"last_read_at,omitempty" xml:"last_read_at,omitempty"`
	// ルームID
	RoomID int `form:"room_id" json:"room_id" yaml:"room_id" xml:"room_id"`
	// 更新日
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
	// ユーザーID
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the UserRoom media type instance.
func (mt *UserRoom) Validate() (err error) {

	return
}
