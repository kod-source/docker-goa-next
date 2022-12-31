import { User } from "./user";

export class Auth {
  constructor(public token: string, public user: User) {}
}
