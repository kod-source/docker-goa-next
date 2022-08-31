import React, { FC, FormEvent, useEffect, useState } from 'react';
import Box from '@mui/material/Box';
import { Avatar, Button, Modal, Typography } from '@mui/material';
import { PostWithUser } from '../model/post';
import { Comment, CommentWithUser } from '../model/comment';
import Image from 'next/image';
import { toStringlinefeed } from '../text';
import { DateTime } from 'luxon';
import axios from 'axios';
import { getEndPoint, getToken } from '../token';
import { Loading } from './loading';
import { isAxiosError } from '../axios';

interface Props {
  open: boolean;
  handleClose: () => void;
  postWithUser: PostWithUser;
  setPostsWithUser: React.Dispatch<React.SetStateAction<PostWithUser[]>>;
}

export const PostModal: FC<Props> = ({
  open,
  handleClose,
  postWithUser,
  setPostsWithUser,
}) => {
  const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 800,
    height: 1000,
    bgcolor: '#222222',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };

  const [commentsWithUser, setCommentsWithUser] = useState<CommentWithUser[]>();
  const [text, setText] = useState('');
  const [imagePath, setImagePath] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const fetchData = async () => {
    try {
      const res = await axios.get(
        `${getEndPoint}/comments/${postWithUser.post.id}`,
        {
          headers: {
            Authorization: `Bearer ${getToken()}`,
          },
        }
      );
      setCommentsWithUser(() => {
        const newCommentsWithUserr = res.data.map((d: any) => {
          const comment = new Comment(
            d.comment.id,
            d.comment.post_id,
            d.comment.user_id,
            d.comment.text,
            new Date(d.comment.created_at),
            new Date(d.comment.updated_at),
            d.comment.img
          );
          return {
            comment: comment,
            user: { id: d.user.id, name: d.user.name, avatar: d.user.avatar },
          };
        });
        return [...newCommentsWithUserr];
      });
    } catch (e) {
      if (isAxiosError(e)) {
        const myAxiosError = e.response;
        if (myAxiosError?.status === 404) {
          setCommentsWithUser([]);
          return;
        }
        return alert(myAxiosError?.statusText);
      }
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const onSubmit = async (e: FormEvent<HTMLFormElement>, postId: number) => {
    setIsLoading(true);
    e.preventDefault();
    const res = await axios.post(
      `${getEndPoint()}/comments`,
      {
        post_id: postId,
        text: text,
        img: imagePath,
      },
      {
        headers: {
          Authorization: `Bearer ${getToken()}`,
        },
      }
    );
    setText('');
    setImagePath('');
    setCommentsWithUser((old) => {
      const newCommentWithUserr: CommentWithUser = {
        comment: new Comment(
          res.data.comment.id,
          res.data.comment.post_id,
          res.data.comment.user_id,
          res.data.comment.text,
          new Date(res.data.comment.created_at),
          new Date(res.data.comment.updated_at),
          res.data.comment.img
        ),
        user: {
          id: res.data.user.id,
          name: res.data.user.name,
          avatar: res.data.user.avatar,
        },
      };
      if (!old) return [newCommentWithUserr];
      return [...old, newCommentWithUserr];
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
      <Box sx={style} className='overflow-scroll'>
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
                  value={text}
                  onChange={(e) => setText(e.target.value)}
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
        {commentsWithUser ? (
          <div>
            {commentsWithUser.map((cu) => (
              <div
                key={cu.comment.id}
                className='my-2 border border-slate-600 p-5 rounded-md'
              >
                <div className='flex justify-center'>
                  <Avatar
                    sx={{ width: 80, height: 80 }}
                    alt='投稿者'
                    src={cu.user.avatar ? cu.user.avatar : '/avatar.png'}
                  />
                  <div className='pt-5 mx-3'>
                    <p>{cu.user.name}</p>
                    <div className='flex'>
                      <p>
                        投稿日：
                        {DateTime.fromJSDate(cu.comment.createdAt).toFormat(
                          'yyyy年MM月dd日'
                        )}
                      </p>
                      {cu.comment.createdAt.getTime() !==
                        cu.comment.updatedAt.getTime() && (
                        <p className='mx-5'>
                          更新日：
                          {DateTime.fromJSDate(cu.comment.updatedAt).toFormat(
                            'yyyy年MM月dd日'
                          )}
                        </p>
                      )}
                    </div>
                  </div>
                  <div className='ml-auto'>
                    <Button
                      className='text-white'
                      // onClick={(e) => onClickDetail(e, postWithUser)}
                    >
                      :
                    </Button>
                  </div>
                </div>
                <div>
                  <p>{toStringlinefeed(cu.comment.text)}</p>
                  {cu.comment.img && (
                    <Image
                      src={cu.comment.img}
                      width={500}
                      height={500}
                      alt={postWithUser.post.title + 'picture'}
                    />
                  )}
                </div>
              </div>
            ))}
          </div>
        ) : (
          <Loading />
        )}
      </Box>
    </Modal>
  );
};
