import React, { FC } from 'react';
import Box from '@mui/material/Box';
import { Avatar, Button, Modal, Typography } from '@mui/material';
import { PostWithUser } from '../model/post';
import { Comment } from '../model/comment';
import Image from 'next/image';
import { toStringlinefeed } from '../text';
import { DateTime } from 'luxon';

interface Props {
  open: boolean;
  handleClose: () => void;
  postWithUser: PostWithUser;
  comments: Comment[];
}

export const PostModal: FC<Props> = ({
  open,
  handleClose,
  postWithUser,
  comments,
}) => {
  const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 800,
    bgcolor: '#222222',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };
  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby='modal-modal-title'
      aria-describedby='modal-modal-description'
    >
      <Box sx={style}>
        <div className='my-2 border border-slate-600 p-5 rounded-md'>
          <div className='flex'>
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
          </div>
          <div>
            <p>{toStringlinefeed(postWithUser.post.title)}</p>
            {/* {!!postWithUser.post.img && (
              <Image
                src={postWithUser.post.img}
                width={500}
                height={500}
                alt={postWithUser.post.title + 'picture'}
              />
            )} */}
          </div>
        </div>
      </Box>
    </Modal>
  );
};
