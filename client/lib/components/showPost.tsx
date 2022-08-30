import React, { FC, useState } from 'react';
import { PostWithUser } from '../model/post';
import { Avatar, Button } from '@mui/material';
import { DateTime } from 'luxon';
import { toStringlinefeed } from '../text';
import Image from 'next/image';
import FavoriteIcon from '@mui/icons-material/Favorite';
import CommentIcon from '@mui/icons-material/Comment';
import ShareIcon from '@mui/icons-material/Share';
import axios, { AxiosError } from 'axios';
import { getToken } from '../token';
import { Comment } from '../model/comment';
import { PostModal } from './postModal';
import { isAxiosError, MyAxiosError } from '../axios';

interface Props {
  postWithUser: PostWithUser;
  setPostsWithUser: React.Dispatch<React.SetStateAction<PostWithUser[]>>;
  myLikePostIds: number[];
  clickLikeButton: (postId: number) => Promise<void>;
  onClickDetail: (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
    p: PostWithUser
  ) => void;
}

export const ShowPost: FC<Props> = ({
  postWithUser,
  setPostsWithUser,
  myLikePostIds,
  clickLikeButton,
  onClickDetail,
}) => {
  const [isShowPostModal, setIsShowPostModal] = useState(false);

  return (
    <>
      <div className='my-10 mx-auto w-3/5'>
        <div className='flex justify-center'>
          <Avatar
            sx={{ width: 80, height: 80 }}
            alt='投稿者'
            src={
              postWithUser.user.avatar
                ? postWithUser.user.avatar
                : '/avatar.png'
            }
          />
          <div className='pt-5 mx-3'>
            <p>{postWithUser.user.name}</p>
            <div className='flex'>
              <p>
                投稿日：
                {DateTime.fromJSDate(postWithUser.post.createdAt).toFormat(
                  'yyyy年MM月dd日'
                )}
              </p>
              {postWithUser.post.createdAt.getTime() !==
                postWithUser.post.updatedAt.getTime() && (
                <p className='mx-5'>
                  更新日：
                  {DateTime.fromJSDate(postWithUser.post.updatedAt).toFormat(
                    'yyyy年MM月dd日'
                  )}
                </p>
              )}
            </div>
          </div>
          <div className='ml-auto'>
            <Button
              className='text-white'
              onClick={(e) => onClickDetail(e, postWithUser)}
            >
              :
            </Button>
          </div>
        </div>
        <div>
          <p>{toStringlinefeed(postWithUser.post.title)}</p>
          {!!postWithUser.post.img && (
            <Image
              src={postWithUser.post.img}
              width={500}
              height={500}
              alt={postWithUser.post.title + 'picture'}
            />
          )}
        </div>
        <div className='flex justify-start'>
          <div
            className='cursor-pointer mr-20 hover:opacity-60'
            onClick={() => setIsShowPostModal(true)}
          >
            <CommentIcon className='mr-3' />
            {postWithUser.countComment}
          </div>
          <div
            className='cursor-pointer mx-20 hover:opacity-60'
            onClick={() => clickLikeButton(postWithUser.post.id)}
          >
            <FavoriteIcon
              className='mr-3'
              color={
                myLikePostIds.includes(postWithUser.post.id)
                  ? 'error'
                  : 'inherit'
              }
            />
            {postWithUser.countLike}
          </div>
          <div className='cursor-pointer mx-20 hover:opacity-60'>
            <ShareIcon />
          </div>
        </div>
      </div>
      {isShowPostModal && (
        <PostModal
          open={isShowPostModal}
          handleClose={() => setIsShowPostModal(false)}
          postWithUser={postWithUser}
          setPostsWithUser={setPostsWithUser}
        />
      )}
    </>
  );
};
