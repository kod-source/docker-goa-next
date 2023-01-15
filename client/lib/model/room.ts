import { User } from "./user";

export class Room {
  constructor(
    public id: number,
    public name: string,
    public isGroup: boolean,
    public createdAt: Date,
    public updatedAt: Date,
    public img?: string,
  ) {}
}

export interface ShowRoom {
  room: Room;
  users: Omit<User, "email" | "password" | "createdAt">[];
}

export interface IndexRoom {
  room: Room;
  isOpen: boolean;
  countUser: number;
  lastText?: string;
  showImg?: string;
}

export interface AllRoom {
  indexRooms: IndexRoom[];
  nextID: number | null;
}
