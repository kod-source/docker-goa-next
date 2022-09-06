// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next';
import { User } from '../../lib/model/user';
import { getEndPoint } from '../../lib/token';

type Data = {
  name: string;
};

export const getUser = async (userId: number): Promise<User> => {
  const response = await fetch(`${getEndPoint()}/users/${userId}`);
  const data = await response.json();
  const user = new User(
    data.id,
    data.name,
    data.email,
    data.created_at,
    data.avatar
  );
  return user;
};

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  res.status(200).json({ name: 'John Doe' });
}
