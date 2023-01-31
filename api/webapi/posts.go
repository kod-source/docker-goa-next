package webapi

import (
	"database/sql"
	"errors"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
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
	p, nextID, err := c.pu.ShowAll(ctx, pointer.IntValue(ctx.NextID))
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(c.toPostAllLimit(p, nextID))
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
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(&app.ShowPostJSON{
		CommentsWithUsers: c.toCommnetJSON(sp.CommenstWithUsers),
		Post: &app.PostJSON{
			ID:        sp.IndexPost.Post.ID,
			UserID:    sp.IndexPost.Post.UserID,
			Title:     sp.IndexPost.Post.Title,
			Img:       sp.IndexPost.Post.Img,
			CreatedAt: &sp.IndexPost.Post.CreatedAt,
			UpdatedAt: &sp.IndexPost.Post.UpdatedAt,
		},
		User: &app.User{
			ID:        int(sp.IndexPost.User.ID),
			Name:      &sp.IndexPost.User.Name,
			Email:     &sp.IndexPost.User.Email,
			Avatar:    sp.IndexPost.User.Avatar,
			CreatedAt: &sp.IndexPost.User.CreatedAt,
		},
		Likes: c.toLikeJSON(sp.Likes),
	})
}

func (c *PostsController) ShowMyLike(ctx *app.ShowMyLikePostsContext) error {
	ips, nextID, err := c.pu.ShowMyLike(ctx, getUserIDCode(ctx), pointer.IntValue(ctx.NextID))
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(c.toPostAllLimit(ips, nextID))
}

func (c *PostsController) ShowPostLike(ctx *app.ShowPostLikePostsContext) error {
	ips, nextID, err := c.pu.ShowMyLike(ctx, ctx.ID, pointer.IntValue(ctx.NextID))
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(c.toPostAllLimit(ips, nextID))
}

func (c *PostsController) ShowPostMy(ctx *app.ShowPostMyPostsContext) error {
	ips, nextID, err := c.pu.ShowPostMy(ctx, ctx.ID, pointer.IntValue(ctx.NextID))
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.OK(c.toPostAllLimit(ips, nextID))
}

func (c *PostsController) ShowPostMedia(ctx *app.ShowPostMediaPostsContext) error {
	ips, nextID, err := c.pu.ShowPostMedia(ctx, ctx.ID, pointer.IntValue(ctx.NextID))
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.OK(c.toPostAllLimit(ips, nextID))
}

func (c *PostsController) toPostAllLimit(indexPosts []*model.IndexPostWithCountLike, nextID *int) *app.PostAllLimit {
	ips := make(app.PostAndUserAndCountLikeJSONCollection, 0, len(indexPosts))
	for _, ip := range indexPosts {
		ips = append(ips, &app.PostAndUserAndCountLikeJSON{
			Post: &app.PostJSON{
				ID:        ip.IndexPost.Post.ID,
				UserID:    ip.IndexPost.Post.UserID,
				Title:     ip.IndexPost.Post.Title,
				Img:       ip.IndexPost.Post.Img,
				CreatedAt: &ip.IndexPost.Post.CreatedAt,
				UpdatedAt: &ip.IndexPost.Post.UpdatedAt,
			},
			UserName:     ip.IndexPost.User.Name,
			Avatar:       ip.IndexPost.User.Avatar,
			CountLike:    ip.CountLike,
			CountComment: ip.CountComment,
		})
	}

	return &app.PostAllLimit{
		ShowPosts: ips,
		NextID:    nextID,
	}
}

func (c *PostsController) toCommnetJSON(commentsWithUsers []*model.ShowCommentWithUser) app.CommentWithUserJSONCollection {
	cs := make(app.CommentWithUserJSONCollection, 0, len(commentsWithUsers))
	for _, cu := range commentsWithUsers {
		if cu.Comment.ID == nil {
			return cs
		}
		cs = append(cs, &app.CommentWithUserJSON{
			Comment: &app.CommentJSON{
				ID:        pointer.IntValue(cu.Comment.ID),
				PostID:    pointer.IntValue(cu.Comment.PostID),
				UserID:    pointer.IntValue(cu.Comment.UserID),
				Text:      pointer.StringValue(cu.Comment.Text),
				Img:       cu.Comment.Img,
				CreatedAt: cu.Comment.CreatedAt,
				UpdatedAt: cu.Comment.UpdatedAt,
			},
			User: &app.User{
				ID:     pointer.IntValue(cu.User.ID),
				Name:   cu.User.Name,
				Avatar: cu.User.Avatar,
			},
		})
	}
	return cs
}

func (c *PostsController) toLikeJSON(likes []*model.Like) app.LikeJSONCollection {
	ls := make(app.LikeJSONCollection, 0, len(likes))
	for _, l := range likes {
		ls = append(ls, &app.LikeJSON{
			ID:     l.ID,
			PostID: l.PostID,
			UserID: l.UserID,
		})
	}
	return ls
}
