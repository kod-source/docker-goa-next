import { asyncApiClient } from "../axios";
import { AllRoom, IndexRoom, Room, ShowRoom } from "../model/room";
import { User } from "../model/user";

export const RoomRepository = {
  create: async (name: string, isGroup: boolean, userIDs: number[]): Promise<ShowRoom> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.post("rooms", {
      name: name,
      is_group: isGroup,
      user_ids: userIDs,
    });

    const users: Omit<User, "email" | "password" | "createdAt">[] = res.data.users.map((u: any) => {
      return {
        id: u.id,
        name: u.name,
        avatar: u.avatar,
      };
    });
    const indexRoom: ShowRoom = {
      room: new Room(
        res.data.id,
        res.data.name,
        res.data.is_group,
        new Date(res.data.created_at),
        new Date(res.data.updated_at),
        res.data.img ? res.data.img : null,
      ),
      users: users,
    };
    return indexRoom;
  },

  index: async (nextID: number): Promise<AllRoom> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`rooms?next_id=${nextID}`);
    const indexRooms: IndexRoom[] = res.data.index_room.map((d: any) => {
      const room = new Room(
        d.room.id,
        d.room.name,
        d.room.is_group,
        new Date(d.room.created_at),
        new Date(d.room.updated_at),
        d.room.img ? d.room.img : null,
      );
      return {
        room: room,
        isOpen: d.is_open,
        countUser: d.count_user,
        lastText: d.last_text ? d.last_text : null,
        showImg: d.show_img ? d.show_img : null,
      };
    });

    return {
      indexRooms: indexRooms,
      nextID: res.data.next_id ? res.data.next_id : null,
    };
  },

  exists: async (userID: number): Promise<Room> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`rooms/exists?user_id=${userID}`);

    const room = new Room(
      res.data.id,
      res.data.name,
      res.data.is_group,
      new Date(res.data.created_at),
      new Date(res.data.updated_at),
      res.data.img ? res.data.img : null,
    );
    return room;
  },
};
