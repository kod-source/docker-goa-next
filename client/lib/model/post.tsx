export class Post {
  constructor(
    public id: number,
    public userId: number,
    public title: string,
    public createdAt: Date,
    public updatedAt: Date,
    public img?: string
  ) {}
}

export interface SelectPost {
  id: number;
  title: string;
  img?: string;
}
