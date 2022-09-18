import React, { FC, useCallback, useContext, useEffect, useState } from 'react';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import { User, UserPostSelection } from '../model/user';
import { AppContext } from '../../pages/_app';
import { PostRepository } from '../repository/post';
import { PostWithUser } from '../model/post';
import { Loading } from './loading';
import { ShowPost } from './showPost';
import { LikeRepository } from '../repository/like';

interface Props {
  value: UserPostSelection;
  showUser: User;
  myLikePostIds: number[];
  setMyLikePostIds: React.Dispatch<React.SetStateAction<number[]>>;
}

export const ShowPostMy: FC<Props> = ({
  value,
  showUser,
  myLikePostIds,
  setMyLikePostIds,
}) => {
  const { user } = useContext(AppContext);
  const [nextID, setNextID] = useState<number | null>(0);
  const [postsWithUser, setPostsWithUser] = useState<PostWithUser[]>([]);
  const [againFetch, setAgainFetch] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const fetchData = async () => {
    if (nextID == null) return;
    setIsLoading(true);
    const allPostLimit =
      value === UserPostSelection.My
        ? await PostRepository.showPostMy(nextID, showUser.id)
        : value === UserPostSelection.Media
        ? await PostRepository.showPostMedia(nextID, showUser.id)
        : await PostRepository.showPostLike(nextID, showUser.id);
    setNextID(allPostLimit.nextId);
    setPostsWithUser((old) => {
      if (nextID === 0) {
        return allPostLimit.postsWithUsers;
      }
      return [...old, ...allPostLimit.postsWithUsers];
    });
    setIsLoading(false);
  };

  useEffect(() => {
    if (againFetch) {
      fetchData();
    }
    window.addEventListener('scroll', changeBottom);
    return () => window.removeEventListener('scroll', changeBottom);
  }, [value, againFetch]);

  const changeBottom = useCallback(() => {
    const bottomPosition =
      document.body.offsetHeight - (window.scrollY + window.innerHeight);
    if (bottomPosition < 0) {
      setAgainFetch(true);
      return;
    }
    setAgainFetch(false);
  }, []);

  const clickLikeButton = async (postID: number) => {
    try {
      if (myLikePostIds.includes(postID)) {
        await LikeRepository.delete(postID);
        setMyLikePostIds((old) => {
          return old.filter((l) => l !== postID);
        });
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postID) {
              return {
                post: o.post,
                user: o.user,
                countLike: o.countLike - 1,
                countComment: o.countComment,
              };
            }
            return o;
          });
          return newPosts;
        });
      } else {
        const like = await LikeRepository.create(postID);
        if (like.postId !== postID) {
          throw new Error('post_id unknow');
        }
        setMyLikePostIds((old) => [...old, postID]);
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postID) {
              return {
                post: o.post,
                user: o.user,
                countLike: o.countLike + 1,
                countComment: o.countComment,
              };
            }
            return o;
          });
          return newPosts;
        });
      }
    } catch (e) {
      if (e instanceof Error) {
        alert(e.message);
      }
    }
  };

  const onClickDetail = () => {};

  return (
    <div>
      <div>
        {postsWithUser.map((p) => (
          <ShowPost
            key={p.post.id}
            postWithUser={p}
            setPostsWithUser={setPostsWithUser}
            myLikePostIds={myLikePostIds}
            clickLikeButton={clickLikeButton}
            onClickDetail={onClickDetail}
          />
        ))}
        {isLoading && (
          <div className='my-10'>
            <Loading />
          </div>
        )}
      </div>
    </div>
  );
};
