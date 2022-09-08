import { asyncApiClient } from '../axios';
import { Like } from '../model/like';

export const LikeRepository = {
  getMyLike: async (): Promise<number[]> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get('likes');

    return res.data == null ? [] : res.data;
  },

  create: async (postID: number): Promise<Like> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.post('likes', { post_id: postID });

    return new Like(res.data.id, res.data.user_id, res.data.post_id);
  },

  delete: async (postID: number): Promise<void> => {
    const apiClient = await asyncApiClient.create();
    return await apiClient.delete('likes', { data: { post_id: postID } });
  },

  getLikeByUser: async (userID: number): Promise<number[]> => {
    const apiClient = await asyncApiClient.create();
    const res = await apiClient.get(`likes/${userID}`);

    return res.data == null ? [] : res.data;
  },
};
