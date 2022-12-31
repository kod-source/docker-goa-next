import { asyncApiClient } from '../axios';
import { IndexRoom, Room, ShowRoom } from '../model/room';
import { User } from '../model/user';

export const RoomRepository = {
  create: async (): Promise<ShowRoom> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.post('rooms');

    const users: Omit<User, 'email' | 'password' | 'createdAt'>[] =
      res.data.users.map((u: any) => {
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
        new Date(res.data.updated_at)
      ),
      users: users,
    };
    return indexRoom;
  },
};
