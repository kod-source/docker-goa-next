import axios from "axios";
import { asyncApiClient } from "../axios";
import { Comment, CommentWithUser } from "../model/comment";
import { Like } from "../model/like";
import { Post, PostAllLimit, PostWithUser, ShowPost } from "../model/post";
import { getToken } from "../token";

export const PostRepository = {
  index: async (nextID: number): Promise<PostAllLimit> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts?next_id=${nextID}`);
    const postsWithUser: PostWithUser[] = res.data.show_posts.map((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img,
      );
      return {
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
        countComment: d.count_comment,
      };
    });

    return {
      postsWithUsers: postsWithUser,
      nextId: res.data.next_id ? res.data.next_id : null,
    };
  },

  create: async (title: string, img?: string): Promise<PostWithUser> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.post("posts", {
      title: title,
      img: img,
    });
    const d = res.data;
    const post = new Post(
      d.post.id,
      d.post.user_id,
      d.post.title,
      new Date(d.post.created_at),
      new Date(d.post.updated_at),
      d.post.img,
    );

    return {
      post: post,
      user: { name: d.user_name, avatar: d.avatar },
      countLike: 0,
      countComment: 0,
    };
  },

  delete: async (postId: number): Promise<void> => {
    const apiClient = await asyncApiClient.create();
    await apiClient.delete(`posts/${postId}`);
  },

  update: async (id: number, title: string, img?: string): Promise<Post> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.put(`posts/${id}`, {
      title: title,
      img: img,
    });
    const postRes = res.data.post;

    return new Post(
      postRes.id,
      postRes.user_id,
      postRes.title,
      new Date(postRes.created_at),
      new Date(postRes.updated_at),
      postRes.img,
    );
  },

  show: async (id: number): Promise<ShowPost> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts/${id}`);
    const data = res.data;
    const likes: Like[] = data.likes.map((l: any) => {
      return { id: l.id, userId: l.user_id, postId: l.post_id };
    });
    const commentsWithUsers: CommentWithUser[] = data.comments_with_users.map((cu: any) => {
      return {
        comment: new Comment(
          cu.comment.id,
          cu.comment.post_id,
          cu.comment.user_id,
          cu.comment.text,
          new Date(cu.comment.created_at),
          new Date(cu.comment.updated_at),
          cu.comment.img,
        ),
        user: {
          id: cu.user.id,
          name: cu.user.name,
          avatar: cu.user.avatar,
        },
      };
    });

    return {
      post: new Post(
        data.post.id,
        data.post.user_id,
        data.post.title,
        new Date(data.post.created_at),
        new Date(data.post.updated_at),
        data.post.img,
      ),
      user: {
        id: data.user.id,
        name: data.user.name,
        avatar: data.user.avatar,
      },
      likes: likes,
      commentsWithUsers: commentsWithUsers,
    };
  },

  // showMyLike 自分がいいねした投稿を取得する
  showMyLike: async (nextID: number): Promise<PostAllLimit> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts/my_like?next_id=${nextID}`);
    const postsWithUser: PostWithUser[] = res.data.show_posts.map((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img,
      );
      return {
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
        countComment: d.count_comment,
      };
    });

    return {
      postsWithUsers: postsWithUser,
      nextId: res.data.next_id ? res.data.next_id : null,
    };
  },

  // showPostLike 指定したユーザーのいいねした投稿を取得する
  showPostLike: async (nextID: number, userID: number): Promise<PostAllLimit> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts/likes/${userID}?next_id=${nextID}`);
    const postsWithUser: PostWithUser[] = res.data.show_posts.map((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img,
      );
      return {
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
        countComment: d.count_comment,
      };
    });

    return {
      postsWithUsers: postsWithUser,
      nextId: res.data.next_id ? res.data.next_id : null,
    };
  },

  // showPostMy 指定したユーザー自身が投稿したものを取得する
  showPostMy: async (nextID: number, userID: number): Promise<PostAllLimit> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts/my_post/${userID}?next_id=${nextID}`);
    const postsWithUser: PostWithUser[] = res.data.show_posts.map((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img,
      );
      return {
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
        countComment: d.count_comment,
      };
    });

    return {
      postsWithUsers: postsWithUser,
      nextId: res.data.next_id ? res.data.next_id : null,
    };
  },

  // showPostMedia 指定したユーザーの画像あり投稿を取得する
  showPostMedia: async (nextID: number, userID: number): Promise<PostAllLimit> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`posts/my_media/${userID}?next_id=${nextID}`);
    const postsWithUser: PostWithUser[] = res.data.show_posts.map((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img,
      );
      return {
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
        countComment: d.count_comment,
      };
    });

    return {
      postsWithUsers: postsWithUser,
      nextId: res.data.next_id ? res.data.next_id : null,
    };
  },
};
