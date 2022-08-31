import axios from 'axios';
import { NextPage, GetServerSideProps } from 'next';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { Comment, CommentWithUser } from '../lib/model/comment';
import { Like } from '../lib/model/like';
import { Post, ShowPost } from '../lib/model/post';
import { getEndPoint, getToken } from '../lib/token';

interface Props {
  id: number;
}

const PostShow: NextPage<Props> = ({ id }) => {
  const router = useRouter();
  const [showPost, setShowPost] = useState<ShowPost>();

  const fetchData = async () => {
    const res = await axios.get(`${getEndPoint()}/posts/${id}`, {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    });
    setShowPost(() => {
      const data = res.data;
      const likes: Like[] = data.likes.map((l: any) => {
        return { id: l.id, userId: l.user_id, postId: l.post_id };
      });
      const commentsWithUsers: CommentWithUser[] = data.comments_with_users.map(
        (cu: any) => {
          return {
            comment: new Comment(
              cu.comment.id,
              cu.comment.post_id,
              cu.comment.user_id,
              cu.comment.text,
              cu.comment.created_at,
              cu.comment.updated_at,
              cu.comment.img
            ),
            user: {
              id: cu.user.id,
              name: cu.user.name,
              avatar: cu.user.avatar,
            },
          };
        }
      );
      return {
        post: new Post(
          data.post.id,
          data.post.user_id,
          data.post.title,
          data.post.created_at,
          data.post.updated_at,
          data.post.img
        ),
        user: {
          id: data.user.id,
          name: data.user.name,
          avatar: data.user.avatar,
        },
        likes: likes,
        commentsWithUsers: commentsWithUsers,
      };
    });
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <h1>詳細ページです</h1>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async (content) => {
  const { id } = content.query;
  return {
    props: {
      id: id,
    },
  };
};

export default PostShow;
