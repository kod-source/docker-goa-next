import React, { FC, FormEvent, useState } from 'react';
import Box from '@mui/material/Box';
import { Button, Modal } from '@mui/material';
import Avatar from '@mui/material/Avatar';
import Image from 'next/image';
import { Post, SelectPost } from '../model/post';
import axios from 'axios';
import { User } from '../model/user';

interface Props {
  open: boolean;
  handleClose: () => void;
  post: SelectPost;
  setPost: React.Dispatch<React.SetStateAction<SelectPost>>;
  postWithUser: {
    post: Post;
    user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
  }[];
  setPostWithUser: React.Dispatch<
    React.SetStateAction<
      | {
          post: Post;
          user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
        }[]
    >
  >;
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
    const token = localStorage.getItem('token');
    if (!token) return;
    const res = await axios.put(
      `http://localhost:3000/posts/${props.post.id}`,
      { title: props.post.title, img: props.post.img },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    const newPostWithUser = props.postWithUser.map((p) => {
      if (p.post.id === res.data.post.id) {
        return {
          post: new Post(
            res.data.post.id,
            res.data.post.user_id,
            res.data.post.title,
            new Date(res.data.post.created_at),
            new Date(res.data.post.updated_at),
            res.data.post.img
          ),
          user: p.user,
        };
      }
      return p;
    });
    props.setPostWithUser(newPostWithUser);
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