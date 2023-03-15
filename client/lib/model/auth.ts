import { User } from "./user";

export class Auth {
    constructor(public token: string, public user: User) {}
}

export interface GoogleRedirect {
    url: string;
}
