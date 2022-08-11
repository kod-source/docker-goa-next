package main

import (
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
	post, err := c.pu.CreatePost(ctx, userID, ctx.Payload.Title, ctx.Payload.Img)
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.Created(&app.PostJSON{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Img:       post.Img,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,
	})
}

func (c *PostsController) Index(ctx *app.IndexPostsContext) error {
	p, err := c.pu.ShowAll(ctx)
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(c.toIndexPostJson(p))
}

func (c *PostsController) Delete(ctx *app.DeletePostsContext) error {
	err := c.pu.Delete(ctx, ctx.ID)
	if err != nil {
		return err
	}

	return ctx.OK(nil)
}

func (c *PostsController) toIndexPostJson(indexPosts []*model.IndexPost) app.IndexPostJSONCollection {
	ips := make(app.IndexPostJSONCollection, 0, len(indexPosts))
	for _, ip := range indexPosts {
		ips = append(ips, &app.IndexPostJSON{
			Avatar:    ip.User.Avatar,
			CreatedAt: &ip.Post.CreatedAt,
			Img:       ip.Post.Img,
			PostID:    ip.Post.ID,
			Title:     ip.Post.Title,
			UpdatedAt: &ip.Post.UpdatedAt,
			UserID:    ip.Post.UserID,
			UserName:  ip.User.Name,
		})
	}

	return ips
}
