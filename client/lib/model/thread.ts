import { User } from "./user";

export class Thread {
    constructor(
        public id: number,
        public roomID: number,
        public userID: number,
        public text: string,
        public createdAt: Date,
        public updatedAt: Date,
        public img?: string,
    ) {}
}

export interface ThreadUser {
    thread: Thread;
    user: Omit<User, "email" | "password">;
}

export interface IndexThread {
    threadUser: ThreadUser;
    countContent: number;
}

export interface AllThread {
    indexThreads: IndexThread[];
    nextID: number | null;
}
