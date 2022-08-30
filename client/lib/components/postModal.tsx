import React, { FC, FormEvent, useState } from 'react';
import Box from '@mui/material/Box';
import { Avatar, Button, Modal, Typography } from '@mui/material';
import { PostWithUser } from '../model/post';
import { Comment } from '../model/comment';
import Image from 'next/image';
import { toStringlinefeed } from '../text';
import { DateTime } from 'luxon';
import axios from 'axios';
import { getToken } from '../token';

interface Props {
  open: boolean;
  handleClose: () => void;
  postWithUser: PostWithUser;
  setPostsWithUser: React.Dispatch<React.SetStateAction<PostWithUser[]>>;
  comments: Comment[];
  setComments: React.Dispatch<React.SetStateAction<Comment[]>>;
}

export const PostModal: FC<Props> = ({
  open,
  handleClose,
  postWithUser,
  setPostsWithUser,
  comments,
  setComments,
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

  const [comment, setComment] = useState('');
  const [imagePath, setImagePath] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit = async (e: FormEvent<HTMLFormElement>, postId: number) => {
    setIsLoading(true);
    e.preventDefault();
    const res = await axios.post(
      'http://localhost:3000/comments',
      {
        post_id: postId,
        text: comment,
        img: imagePath,
      },
      {
        headers: {
          Authorization: `Bearer ${getToken()}`,
        },
      }
    );
    setComment('');
    setImagePath('');
    setComments((old) => {
      const newComment = new Comment(
        res.data.id,
        res.data.post_id,
        res.data.text,
        res.data.created_at,
        res.data.updated_at,
        res.data.img
      );
      return [...old, newComment];
    });
    setPostsWithUser((old) => {
      const newPosts = old.map((p) => {
        if (p.post.id === postId) {
          return {
            post: p.post,
            user: p.user,
            countLike: p.countLike,
            countComment: p.countComment + 1,
          };
        }
        return p;
      });
      return newPosts;
    });
    setIsLoading(false);
  };

  const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e: ProgressEvent<FileReader>) => {
        if (!e.target) return null;
        setImagePath(e.target.result as string);
      };
      reader.readAsDataURL(file);
    }
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
        <div className='text-center'>
          <form onSubmit={(e) => onSubmit(e, postWithUser.post.id)}>
            <div className='flex justify-center'>
              <label className='cursor-pointer'>
                <textarea
                  className='w-96'
                  autoFocus
                  required
                  value={comment}
                  onChange={(e) => setComment(e.target.value)}
                />
              </label>
            </div>
            <div className='my-2'>
              {!!imagePath && (
                <div className='relative'>
                  <Image
                    src={imagePath}
                    width={500}
                    height={500}
                    alt={'post picture'}
                  />
                  <div className='absolute left-[35%] bottom-[90%]'>
                    <Button onClick={() => setImagePath('')}>❌</Button>
                  </div>
                </div>
              )}
            </div>
            <div className='mr-36'>
              <label className='cursor-pointer'>
                <Avatar src='/add_photo.jpg' className='m-auto' />
                <input
                  type='file'
                  className='hidden'
                  accept='image/*'
                  onChange={onChangeInputFile}
                />
              </label>
              <div className='text-right'>
                <Button type='submit' disabled={isLoading}>
                  {isLoading ? 'アップロード中' : '返信'}
                </Button>
              </div>
            </div>
          </form>
        </div>
        <div>
          {comments.map((c) => (
            <div
              key={c.id}
              className='my-2 border border-slate-600 p-5 rounded-md'
            >
              <p>{toStringlinefeed(c.text)}</p>
              {c.img && (
                <Image
                  src={c.img}
                  width={500}
                  height={500}
                  alt={postWithUser.post.title + 'picture'}
                />
              )}
            </div>
          ))}
        </div>
      </Box>
    </Modal>
  );
};
