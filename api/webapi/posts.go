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

func (c *PostsController) toIndexPostJson(indexPosts []*model.IndexPost) app.IndexPostJSONCollection {
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

	return ips
}
