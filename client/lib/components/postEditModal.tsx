import React, { FC, FormEvent, useState } from 'react';
import Box from '@mui/material/Box';
import { Button, Modal } from '@mui/material';
import Avatar from '@mui/material/Avatar';
import Image from 'next/image';
import { Post, PostWithUser, SelectPost } from '../model/post';
import axios from 'axios';
import { getEndPoint, getToken } from '../token';
import { PostRepository } from '../repository/post';

interface Props {
  open: boolean;
  handleClose: () => void;
  post: SelectPost;
  setPost: React.Dispatch<React.SetStateAction<SelectPost>>;
  setPostsWithUsers: React.Dispatch<React.SetStateAction<PostWithUser[]>>;
}

export const PostEditModal: FC<Props> = (props) => {
  const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };
  const [isUpdating, setIsUpdating] = useState(false);

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsUpdating(true);
    const post = await PostRepository.update(
      props.post.id,
      props.post.title,
      props.post.img
    );
    props.setPostsWithUsers((old) => {
      const newPosts = old.map((p) => {
        if (p.post.id === props.post.id) {
          return {
            post: post,
            user: p.user,
            countLike: p.countLike,
            countComment: p.countComment,
          };
        }
        return p;
      });
      return newPosts;
    });
    setIsUpdating(false);
    props.handleClose();
  };

  const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e: ProgressEvent<FileReader>) => {
        if (!e.target) return null;
        props.setPost((old) => ({
          id: old.id,
          title: old.title,
          img: e.target?.result as string,
        }));
      };
      reader.readAsDataURL(file);
    }
  };
  return (
    <Modal
      open={props.open}
      onClose={props.handleClose}
      aria-labelledby='modal-modal-title'
      aria-describedby='modal-modal-description'
    >
      <Box sx={style} className='text-center'>
        <div className='text-center'>
          <form onSubmit={onSubmit}>
            <div className='flex justify-center'>
              <label className='cursor-pointer'>
                <textarea
                  className='w-96'
                  autoFocus
                  required
                  value={props.post.title}
                  onChange={(e) =>
                    props.setPost((old) => ({
                      id: old.id,
                      title: e.target.value,
                      img: old.img,
                    }))
                  }
                />
              </label>
            </div>
            <div className='my-2'>
              {!!props.post.img && (
                <div className='relative'>
                  <Image
                    src={props.post.img}
                    width={500}
                    height={500}
                    alt={'post picture'}
                  />
                  <div className='absolute left-[35%] bottom-[90%]'>
                    <Button
                      onClick={() =>
                        props.setPost((old) => ({
                          id: old.id,
                          title: old.title,
                          img: '',
                        }))
                      }
                    >
                      ❌
                    </Button>
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
                <Button type='submit' disabled={isUpdating}>
                  {isUpdating ? '更新中' : '更新する'}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </Box>
    </Modal>
  );
};
