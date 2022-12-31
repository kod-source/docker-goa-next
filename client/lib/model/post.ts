import { CommentWithUser } from "./comment";
import { Like } from "./like";
import { User } from "./user";
export class Post {
  constructor(
    public id: number,
    public userId: number,
    public title: string,
    public createdAt: Date,
    public updatedAt: Date,
    public img?: string,
  ) {}
}

export interface SelectPost {
  id: number;
  title: string;
  img?: string;
}

export interface PostWithUser {
  post: Post;
  user: Omit<User, "id" | "email" | "email" | "password" | "createdAt">;
  countLike: number;
  countComment: number;
}

export interface ShowPost {
  post: Post;
  user: Omit<User, "email" | "password" | "createdAt">;
  likes: Like[];
  commentsWithUsers: CommentWithUser[];
}

export interface PostAllLimit {
  postsWithUsers: PostWithUser[];
  nextId: number | null;
}
