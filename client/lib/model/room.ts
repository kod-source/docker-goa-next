import { User } from './user';

export class Room {
  constructor(
    public id: number,
    public name: string,
    public isGroup: boolean,
    public createdAt: Date,
    public updatedAt: Date
  ) {}
}

export interface ShowRoom {
  room: Room;
  users: Omit<User, 'email' | 'password' | 'createdAt'>[];
}

export interface IndexRoom {
  room: Room;
  lastText: string;
  isOpen: boolean;
}
