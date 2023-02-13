export class User {
    constructor(
        public id: number,
        public name: string,
        public email: string,
        public createdAt: Date,
        public avatar?: string,
    ) {}
}

export const enum UserPostSelection {
    My = "my",
    Media = "media",
    Like = "like",
}
