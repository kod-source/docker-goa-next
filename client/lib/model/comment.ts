import { User } from "./user";

export class Comment {
  constructor(
    public id: number,
    public postId: number,
    public userId: number,
    public text: string,
    public createdAt: Date,
    public updatedAt: Date,
    public img?: string,
  ) {}
}

export interface CommentWithUser {
  comment: Comment;
  user: Omit<User, "email" | "email" | "password" | "createdAt">;
}
