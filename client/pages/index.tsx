import { Button } from '@mui/material';
import { DateTime } from 'luxon';
import type { NextPage } from 'next';
import Head from 'next/head';
import { useRouter } from 'next/router';
import { FormEvent, useContext, useEffect, useState } from 'react';
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
import InfiniteScroll from 'react-infinite-scroller';

const Home: NextPage = () => {
  const { user } = useContext(AppContext);
  const router = useRouter();
  const [postsWithUser, setPostsWithUser] = useState<
    {
      post: Post;
      user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
    }[]
  >([]);
  const [post, setPost] = useState<{ title: string; img: string }>({
    title: '',
    img: '',
  });
  const [nextToken, setNextToken] = useState<string | null>(null);
  const [hasMore, setHasMore] = useState(true);
  const [isShowDetailModal, setIsShowDetailModal] = useState(false);
  const [widthAndHeightRate, setWidthAndHeightRate] = useState({
    width: '',
    height: '',
  });
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [postID, setPostID] = useState(0);
  const [isMyPost, setIsMyPost] = useState(false);
  const [showPostEditModal, setShowPostEditModal] = useState(false);
  const [selectPost, setSelectPost] = useState<SelectPost>({
    id: 0,
    title: '',
    img: '',
  });

  const logout = () => {
    localStorage.removeItem('token');
    router.push('/login');
  };

  const fetchData = async () => {
    const token = localStorage.getItem('token');
    if (!token) return;
    const res = await axios.get('http://localhost:3000/posts', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    const postsWithUser: {
      post: Post;
      user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
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
      });
    });
    setNextToken(res.data.next_token);
    setPostsWithUser(postsWithUser);
  };

  useEffect(() => {
    fetchData();
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
      const token = localStorage.getItem('token');
      if (!token) return;
      const res = await axios.post(
        'http://localhost:3000/posts',
        {
          title: post.title,
          img: post.img,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
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
          { post: post, user: { name: d.user_name, avatar: d.avatar } },
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
      const token = localStorage.getItem('token');
      if (!token) return;
      await axios.delete(`http://localhost:3000/posts/${postID}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setShowConfirmModal(false);
      setPostsWithUser(postsWithUser?.filter((p) => p.post.id !== postID));
    } catch (e) {
      if (e instanceof Error) {
        alert(e.message);
      }
    }
  };

  const loadMore = async () => {
    const token = localStorage.getItem('token');
    if (!token) return;
    const endPoint = nextToken ? nextToken : 'http://localhost:3000/posts';
    const res = await axios.get(endPoint, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    const nextT: string | null = res.data.next_token;
    setNextToken(nextT);
    if (!nextT) {
      setHasMore(false);
    }
    const postsWithUser: {
      post: Post;
      user: Omit<User, 'id' | 'email' | 'email' | 'password' | 'createdAt'>;
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
      });
    });
    setPostsWithUser((old) => {
      if (!nextToken) {
        return [...postsWithUser];
      }
      return [...old, ...postsWithUser];
    });
  };

  if (!user) {
    return <Loading />;
  }
  const loading = <Loading />;
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
      {/* {postsWithUser ? (
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
                          String((currentWidth / window.innerWidth) * 100) +
                          '%',
                        height:
                          String((currentHeight / window.innerHeight) * 100) +
                          '%',
                      });
                      setPostID(p.post.id);
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
            </div>
          ))}
        </div>
      ) : (
        <Loading />
      )} */}
      <InfiniteScroll loadMore={loadMore} hasMore={hasMore} loader={loading}>
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
                    setPostID(p.post.id);
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
          </div>
        ))}
      </InfiniteScroll>
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
