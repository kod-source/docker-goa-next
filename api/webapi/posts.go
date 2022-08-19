package main

import (
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// PostsController implements the posts resource.
type PostsController struct {
	*goa.Controller
	pu usecase.PostUseCase
}

// NewPostsController creates a posts controller.
func NewPostsController(service *goa.Service, pu usecase.PostUseCase) *PostsController {
	return &PostsController{Controller: service.NewController("PostsController"), pu: pu}
}

// CreatePost runs the create_post action.
func (c *PostsController) CreatePost(ctx *app.CreatePostPostsContext) error {
	userID := getUserIDCode(ctx)
	ip, err := c.pu.CreatePost(ctx, userID, ctx.Payload.Title, ctx.Payload.Img)
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.Created(&app.IndexPostJSON{
		Post: &app.PostJSON{
			ID:        ip.Post.ID,
			UserID:    ip.Post.UserID,
			Title:     ip.Post.Title,
			Img:       ip.Post.Img,
			CreatedAt: &ip.Post.CreatedAt,
			UpdatedAt: &ip.Post.UpdatedAt,
		},
		Avatar:   ip.User.Avatar,
		UserName: ip.User.Name,
	})
}

func (c *PostsController) Index(ctx *app.IndexPostsContext) error {
	var nextID int
	if ctx.NextID == nil {
		nextID = 0
	} else {
		nextID = *ctx.NextID
	}
	p, nextToken, err := c.pu.ShowAll(ctx, nextID)
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(c.toPostAllLimit(p, nextToken))
}

func (c *PostsController) Delete(ctx *app.DeletePostsContext) error {
	err := c.pu.Delete(ctx, ctx.ID)
	if err != nil {
		return err
	}

	return ctx.OK(nil)
}

func (c *PostsController) Update(ctx *app.UpdatePostsContext) error {
	ip, err := c.pu.Update(ctx, ctx.ID, ctx.Payload.Title, ctx.Payload.Img)
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.OK(&app.IndexPostJSON{
		Post: &app.PostJSON{
			ID:        ip.Post.ID,
			UserID:    ip.Post.UserID,
			Title:     ip.Post.Title,
			Img:       ip.Post.Img,
			CreatedAt: &ip.Post.CreatedAt,
			UpdatedAt: &ip.Post.UpdatedAt,
		},
		UserName: ip.User.Name,
		Avatar:   ip.User.Avatar,
	})
}

func (c *PostsController) Show(ctx *app.ShowPostsContext) error {
	sp, err := c.pu.Show(ctx, ctx.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(&app.ShowPostJSON{
		Comments: c.toCommnetJson(sp.Comments),
		Post: &app.PostJSON{
			ID:        sp.IndexPost.Post.ID,
			UserID:    sp.IndexPost.Post.UserID,
			Title:     sp.IndexPost.Post.Title,
			Img:       sp.IndexPost.Post.Img,
			CreatedAt: &sp.IndexPost.Post.CreatedAt,
			UpdatedAt: &sp.IndexPost.Post.UpdatedAt,
		},
		User: &app.User{
			ID:        sp.IndexPost.User.ID,
			Name:      &sp.IndexPost.User.Name,
			Email:     &sp.IndexPost.User.Email,
			Avatar:    sp.IndexPost.User.Avatar,
			CreatedAt: &sp.IndexPost.User.CreatedAt,
		},
	})
}

func (c *PostsController) toPostAllLimit(indexPosts []*model.IndexPost, nextToken *string) *app.PostAllLimit {
	ips := make(app.IndexPostJSONCollection, 0, len(indexPosts))
	for _, ip := range indexPosts {
		ips = append(ips, &app.IndexPostJSON{
			Post: &app.PostJSON{
				ID:        ip.Post.ID,
				UserID:    ip.Post.UserID,
				Title:     ip.Post.Title,
				Img:       ip.Post.Img,
				CreatedAt: &ip.Post.CreatedAt,
				UpdatedAt: &ip.Post.UpdatedAt,
			},
			UserName: ip.User.Name,
			Avatar:   ip.User.Avatar,
		})
	}

	return &app.PostAllLimit{
		ShowPosts: ips,
		NextToken: nextToken,
	}
}

func (c *PostsController) toCommnetJson(comments []*model.Comment) app.CommentJSONCollection {
	cs := make(app.CommentJSONCollection, 0, len(comments))
	for _, c := range comments {
		if c.ID == nil {
			return cs
		}
		cs = append(cs, &app.CommentJSON{
			ID:        *c.ID,
			PostID:    *c.PostID,
			Text:      *c.Text,
			Img:       c.Img,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return cs
}
