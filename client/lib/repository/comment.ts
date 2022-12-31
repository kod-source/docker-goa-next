import { asyncApiClient } from "../axios";
import { Comment, CommentWithUser } from "../model/comment";

export const CommentRepository = {
  create: async (postID: number, text: string, img?: string): Promise<CommentWithUser> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.post("comments", {
      post_id: postID,
      text: text,
      img: img,
    });
    const comment = res.data.comment;

    return {
      comment: new Comment(
        comment.id,
        comment.post_id,
        comment.user_id,
        comment.text,
        new Date(comment.created_at),
        new Date(comment.updated_at),
        comment.img,
      ),
      user: {
        id: res.data.user.id,
        name: res.data.user.name,
        avatar: res.data.user.avatar,
      },
    };
  },

  show: async (postID: number): Promise<CommentWithUser[]> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`comments/${postID}`);
    const commentsWithUserr = res.data.map((d: any) => {
      const comment = new Comment(
        d.comment.id,
        d.comment.post_id,
        d.comment.user_id,
        d.comment.text,
        new Date(d.comment.created_at),
        new Date(d.comment.updated_at),
        d.comment.img,
      );
      return {
        comment: comment,
        user: { id: d.user.id, name: d.user.name, avatar: d.user.avatar },
      };
    });

    return commentsWithUserr;
  },

  update: async (id: number, text: string, img?: string): Promise<Comment> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.put(`comments/${id}`, {
      text: text,
      img: img,
    });

    return new Comment(
      res.data.id,
      res.data.post_id,
      res.data.user_id,
      res.data.text,
      new Date(res.data.created_at),
      new Date(res.data.updated_at),
      res.data.img,
    );
  },

  delete: async (id: number): Promise<void> => {
    const apiClient = await asyncApiClient.create();
    return await apiClient.delete(`comments/${id}`);
  },
};
