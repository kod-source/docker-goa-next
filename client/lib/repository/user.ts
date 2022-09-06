import { asyncApiClient } from '../axios';
import { User } from '../model/user';

export const UserRepostiory = {
  get: async (id: number): Promise<User> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`users/${id}`);
    const user = new User(
      res.data.id,
      res.data.name,
      res.data.email,
      res.data.created_at,
      res.data.avatar
    );
    return user;
  },

  currentUser: async (): Promise<User> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get('current_user');
    const user = new User(
      res.data.id,
      res.data.name,
      res.data.email,
      res.data.created_at,
      res.data.avatar
    );
    return user;
  },
};
