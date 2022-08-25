import { Button } from '@mui/material';
import { DateTime } from 'luxon';
import type { NextPage } from 'next';
import Head from 'next/head';
import { useRouter } from 'next/router';
import { FormEvent, useCallback, useContext, useEffect, useState } from 'react';
import styles from '../styles/Home.module.css';
import { AppContext } from './_app';
import Avatar from '@mui/material/Avatar';
import { Post, SelectPost } from '../lib/model/post';
import { User } from '../lib/model/user';
import axios from 'axios';
import { toStringlinefeed } from '../lib/text';
import { isAxiosError, MyAxiosError } from '../lib/axios';
import Image from 'next/image';
import { Loading } from '../lib/components/loading';
import { DetailModal } from '../lib/components/detailModal';
import { ConfirmationModal } from '../lib/components/confirmationModal';
import { PostEditModal } from '../lib/components/postEditModal';
import FavoriteIcon from '@mui/icons-material/Favorite';
import { getToken } from '../lib/token';

const Home: NextPage = () => {
  const { user } = useContext(AppContext);
  const router = useRouter();
  const [postsWithUser, setPostsWithUser] = useState<
    {
      post: Post;
      user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
      countLike: number;
    }[]
  >([]);
  const [post, setPost] = useState<{ title: string; img: string }>({
    title: '',
    img: '',
  });
  const [isShowDetailModal, setIsShowDetailModal] = useState(false);
  const [widthAndHeightRate, setWidthAndHeightRate] = useState({
    width: '',
    height: '',
  });
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [selectPostID, setSelectPostID] = useState(0);
  const [isMyPost, setIsMyPost] = useState(false);
  const [showPostEditModal, setShowPostEditModal] = useState(false);
  const [selectPost, setSelectPost] = useState<SelectPost>({
    id: 0,
    title: '',
    img: '',
  });
  const [againFetch, setAgainFetch] = useState(true);
  const [nextToken, setNextToken] = useState<string | null>(
    'http://localhost:3000/posts'
  );
  const [isLoading, setIsLoading] = useState(false);
  const [myLikePostIds, setMyLikePostIds] = useState<number[]>([]);

  const logout = () => {
    localStorage.removeItem('token');
    router.push('/login');
  };

  const fetchPostData = async () => {
    if (!nextToken) return;
    setIsLoading(true);
    const res = await axios.get(nextToken, {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    });
    const postsWithUser: {
      post: Post;
      user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
      countLike: number;
    }[] = [];
    res.data.show_posts.forEach((d: any) => {
      const post = new Post(
        d.post.id,
        d.post.user_id,
        d.post.title,
        new Date(d.post.created_at),
        new Date(d.post.updated_at),
        d.post.img
      );
      postsWithUser.push({
        post: post,
        user: { name: d.user_name, avatar: d.avatar },
        countLike: d.count_like,
      });
    });
    const nextT: string | null = res.data.next_token;
    setNextToken(nextT);
    setPostsWithUser((old) => {
      if (nextToken === 'http://localhost:3000/posts') {
        return postsWithUser;
      }
      return [...old, ...postsWithUser];
    });
    setIsLoading(false);
  };

  const fetchData = async () => {
    const res = await axios.get('http://localhost:3000/likes', {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    });
    setMyLikePostIds(res.data);
  };

  useEffect(() => {
    if (againFetch) {
      fetchPostData();
    }
    window.addEventListener('scroll', changeBottom);
    return () => window.removeEventListener('scroll', changeBottom);
  }, [againFetch]);

  useEffect(() => {
    fetchData();
  }, []);

  const changeBottom = useCallback(() => {
    const bottomPosition =
      document.body.offsetHeight - (window.scrollY + window.innerHeight);
    if (bottomPosition < 0) {
      setAgainFetch(true);
      return;
    }
    setAgainFetch(false);
  }, []);

  const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e: ProgressEvent<FileReader>) => {
        if (!e.target) return null;
        setPost((old) => ({
          title: old.title,
          img: e.target?.result as string,
        }));
      };
      reader.readAsDataURL(file);
    }
  };

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const res = await axios.post(
        'http://localhost:3000/posts',
        {
          title: post.title,
          img: post.img,
        },
        {
          headers: {
            Authorization: `Bearer ${getToken()}`,
          },
        }
      );
      setPost({ title: '', img: '' });
      setPostsWithUser((old) => {
        const d = res.data;
        const post = new Post(
          d.post.id,
          d.post.user_id,
          d.post.title,
          new Date(d.post.created_at),
          new Date(d.post.updated_at),
          d.post.img
        );
        return [
          {
            post: post,
            user: { name: d.user_name, avatar: d.avatar },
            countLike: 0,
          },
          ...old,
        ];
      });
    } catch (e) {
      if (isAxiosError(e)) {
        const myAxiosError = e.response?.data as MyAxiosError;
        if (!myAxiosError.message) {
          return alert(e.message);
        }
        return alert(myAxiosError.message);
      }
    }
  };

  const onDelete = async () => {
    try {
      await axios.delete(`http://localhost:3000/posts/${selectPostID}`, {
        headers: {
          Authorization: `Bearer ${getToken()}`,
        },
      });
      setShowConfirmModal(false);
      setPostsWithUser(
        postsWithUser?.filter((p) => p.post.id !== selectPostID)
      );
    } catch (e) {
      if (e instanceof Error) {
        alert(e.message);
      }
    }
  };

  const clickLikeButton = async (postId: number) => {
    try {
      if (myLikePostIds.includes(postId)) {
        await axios.delete('http://localhost:3000/likes', {
          headers: {
            Authorization: `Bearer ${getToken()}`,
          },
          data: {
            post_id: postId,
          },
        });
        setMyLikePostIds((old) => {
          return old.filter((i) => i !== postId);
        });
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postId) {
              return {
                post: o.post,
                user: o.user,
                countLike: o.countLike - 1,
              };
            }
            return o;
          });
          return newPosts;
        });
      } else {
        const res = await axios.post(
          'http://localhost:3000/likes',
          {
            post_id: postId,
          },
          {
            headers: {
              Authorization: `Bearer ${getToken()}`,
            },
          }
        );
        if (res.data.post_id !== postId) {
          throw new Error('post_id unknow');
        }
        setMyLikePostIds((old) => [...old, postId]);
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postId) {
              return {
                post: o.post,
                user: o.user,
                countLike: o.countLike + 1,
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

  if (!user) {
    return <Loading />;
  }
  return (
    <div className={styles.container}>
      <Head>
        <title>Top Page</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>
      <div className='text-right m-2'>
        <Button onClick={() => logout()}>ログアウト</Button>
      </div>
      <div className='text-center'>
        <form onSubmit={onSubmit}>
          <div className='flex justify-center'>
            <Avatar
              sx={{ width: 80, height: 80 }}
              className='mx-5'
              alt='投稿者'
              src={user.avatar ? user.avatar : '/avatar.png'}
            />
            <label className='cursor-pointer'>
              <textarea
                className='w-96'
                autoFocus
                required
                value={post.title}
                onChange={(e) =>
                  setPost((old) => ({ title: e.target.value, img: old.img }))
                }
              />
            </label>
          </div>
          <div className='my-2'>
            {!!post.img && (
              <div className='relative'>
                <Image
                  src={post.img}
                  width={500}
                  height={500}
                  alt={'post picture'}
                />
                <div className='absolute left-[35%] bottom-[90%]'>
                  <Button
                    onClick={() =>
                      setPost((old) => ({ title: old.title, img: '' }))
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
              <Button type='submit'>投稿する</Button>
            </div>
          </div>
        </form>
      </div>
      <div>
        {postsWithUser.map((p) => (
          <div key={p.post.id} className='my-5 mx-auto w-3/5'>
            <div className='flex justify-center'>
              <Avatar
                sx={{ width: 80, height: 80 }}
                alt='投稿者'
                src={p.user.avatar ? p.user.avatar : '/avatar.png'}
              />
              <div className='pt-5 mx-3'>
                <p>{p.user.name}</p>
                <div className='flex'>
                  <p>
                    投稿日：
                    {DateTime.fromJSDate(p.post.createdAt).toFormat(
                      'yyyy年MM月dd日'
                    )}
                  </p>
                  {p.post.createdAt.getTime() !==
                    p.post.updatedAt.getTime() && (
                    <p className='mx-5'>
                      更新日：
                      {DateTime.fromJSDate(p.post.updatedAt).toFormat(
                        'yyyy年MM月dd日'
                      )}
                    </p>
                  )}
                </div>
              </div>
              <div className='ml-auto'>
                <Button
                  className='text-white'
                  onClick={(e) => {
                    const currentWidth = e.clientX;
                    const currentHeight = e.clientY;
                    setWidthAndHeightRate({
                      width:
                        String((currentWidth / window.innerWidth) * 100) + '%',
                      height:
                        String((currentHeight / window.innerHeight) * 100) +
                        '%',
                    });
                    setSelectPostID(p.post.id);
                    setIsShowDetailModal(true);
                    setIsMyPost(p.post.userId === user.id);
                    setSelectPost({
                      id: p.post.id,
                      title: p.post.title,
                      img: p.post.img,
                    });
                  }}
                >
                  :
                </Button>
              </div>
            </div>
            <div>
              <p>{toStringlinefeed(p.post.title)}</p>
              {!!p.post.img && (
                <Image
                  src={p.post.img}
                  width={500}
                  height={500}
                  alt={p.post.title + 'picture'}
                />
              )}
            </div>
            <div>
              <div
                className='cursor-pointer hover:opacity-60'
                onClick={() => clickLikeButton(p.post.id)}
              >
                <FavoriteIcon
                  className='mr-3'
                  color={
                    myLikePostIds.includes(p.post.id) ? 'error' : 'inherit'
                  }
                />
                {p.countLike}
              </div>
            </div>
          </div>
        ))}
        {isLoading && (
          <div className='my-10'>
            <Loading />
          </div>
        )}
      </div>
      {isShowDetailModal && (
        <DetailModal
          open={isShowDetailModal}
          handleClose={() => setIsShowDetailModal(false)}
          widthRate={widthAndHeightRate.width}
          heightRate={widthAndHeightRate.height}
          onDeleteClick={() => {
            setIsShowDetailModal(false);
            setShowConfirmModal(true);
          }}
          onUpdateClick={() => {
            setIsShowDetailModal(false);
            setShowPostEditModal(true);
          }}
          isMyPost={isMyPost}
        />
      )}
      {showConfirmModal && (
        <ConfirmationModal
          open={showConfirmModal}
          handleClose={() => setShowConfirmModal(false)}
          text='削除してもよろしいですか？'
          confirmInvoke={() => onDelete()}
        />
      )}
      {showPostEditModal && (
        <PostEditModal
          open={showPostEditModal}
          handleClose={() => setShowPostEditModal(false)}
          post={selectPost}
          setPost={setSelectPost}
          postWithUser={postsWithUser}
          setPostWithUser={setPostsWithUser}
        />
      )}
    </div>
  );
};

export default Home;
