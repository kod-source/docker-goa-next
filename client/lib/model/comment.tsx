export class Comment {
  constructor(
    public id: number,
    public postId: number,
    public text: string,
    public createdAt: Date,
    public updatedAt: Date,
    public img?: string
  ) {}
}
