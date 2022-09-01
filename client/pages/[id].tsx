import { Avatar, Button } from '@mui/material';
import axios from 'axios';
import { DateTime } from 'luxon';
import { NextPage, GetServerSideProps } from 'next';
import Image from 'next/image';
import { useRouter } from 'next/router';
import { FormEvent, useContext, useEffect, useState } from 'react';
import { Loading } from '../lib/components/loading';
import { Comment, CommentWithUser } from '../lib/model/comment';
import { Like } from '../lib/model/like';
import { Post, ShowPost } from '../lib/model/post';
import { toStringlinefeed } from '../lib/text';
import { getEndPoint, getToken } from '../lib/token';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import FavoriteIcon from '@mui/icons-material/Favorite';
import CommentIcon from '@mui/icons-material/Comment';
import ShareIcon from '@mui/icons-material/Share';
import { AppContext } from './_app';

interface Props {
  id: number;
}

const PostShow: NextPage<Props> = ({ id }) => {
  const router = useRouter();
  const { user } = useContext(AppContext);
  const [showPost, setShowPost] = useState<ShowPost>();
  const [text, setText] = useState('');
  const [imagePath, setImagePath] = useState('');
  const [isLoading, setIsLoading] = useState(false);

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
              new Date(cu.comment.created_at),
              new Date(cu.comment.updated_at),
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
          new Date(data.post.created_at),
          new Date(data.post.updated_at),
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
    setShowPost((old) => {
      if (!old) return;
      const newCommentWithUser: CommentWithUser = {
        comment: new Comment(
          res.data.comment.id,
          res.data.comment.post_id,
          res.data.comment.user_id,
          res.data.comment.text,
          new Date(res.data.comment.created_at),
          new Date(res.data.comment.updated_at)
        ),
        user: {
          id: res.data.user.id,
          name: res.data.user.name,
          avatar: res.data.user.avatar,
        },
      };
      return {
        post: old.post,
        user: old.user,
        likes: old.likes,
        commentsWithUsers: [newCommentWithUser, ...old.commentsWithUsers],
      };
    });
    setIsLoading(false);
  };

  const clickLikeButton = async (postId: number) => {
    try {
      if (showPost?.likes.some((l) => l.userId === user?.id)) {
        await axios.delete(`${getEndPoint()}/likes`, {
          headers: {
            Authorization: `Bearer ${getToken()}`,
          },
          data: {
            post_id: postId,
          },
        });
        setShowPost((old) => {
          if (!old) return;
          const filterLikes = old.likes.filter(
            (l) => !(l.userId === user?.id && l.postId === postId)
          );
          return {
            post: old.post,
            user: old.user,
            likes: filterLikes,
            commentsWithUsers: old.commentsWithUsers,
          };
        });
      } else {
        const res = await axios.post(
          `${getEndPoint()}/likes`,
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
        setShowPost((old) => {
          if (!old) return;
          const newLike = new Like(
            res.data.id,
            res.data.user_id,
            res.data.post_id
          );
          return {
            post: old.post,
            user: old.user,
            likes: [...old.likes, newLike],
            commentsWithUsers: old.commentsWithUsers,
          };
        });
      }
    } catch (e) {
      if (e instanceof Error) {
        alert(e.message);
      }
    }
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

  if (!showPost) return <Loading />;
  return (
    <div className='mx-auto w-3/5 '>
      <div className='my-5 flex'>
        <ArrowBackIcon
          className='mr-5 cursor-pointer hover:opacity-60'
          onClick={() => router.push('/')}
        />
        <h2>投稿</h2>
      </div>
      <div className='my-5 border border-slate-600 p-5 rounded-md cursor-pointer'>
        <div className='flex justify-center'>
          <Avatar
            sx={{ width: 80, height: 80 }}
            alt='投稿者'
            src={showPost.user.avatar ? showPost.user.avatar : '/avatar.png'}
          />
          <div className='pt-5 mx-3'>
            <p>{showPost.user.name}</p>
            <div className='flex'>
              <p>
                投稿日：
                {DateTime.fromJSDate(showPost.post.createdAt).toFormat(
                  'yyyy年MM月dd日'
                )}
              </p>
              {showPost.post.createdAt.getTime() !==
                showPost.post.updatedAt.getTime() && (
                <p className='mx-5'>
                  更新日：
                  {DateTime.fromJSDate(showPost.post.updatedAt).toFormat(
                    'yyyy年MM月dd日'
                  )}
                </p>
              )}
            </div>
          </div>
          <div className='ml-auto'>
            <Button
              className='text-white'
              // onClick={(e) => {
              //   e.stopPropagation();
              //   onClickDetail(e, postWithUser);
              // }}
            >
              :
            </Button>
          </div>
        </div>
        <div>
          <p>{toStringlinefeed(showPost.post.title)}</p>
          {!!showPost.post.img && (
            <Image
              src={showPost.post.img}
              width={500}
              height={500}
              alt={showPost.post.title + 'picture'}
            />
          )}
        </div>
        <div className='flex justify-start'>
          <div
            className='cursor-pointer mr-20 hover:opacity-60'
            //   onClick={(e) => {
            //     e.stopPropagation();
            //     setIsShowPostModal(true);
            //   }}
          >
            <CommentIcon className='mr-3' />
            {showPost.commentsWithUsers.length}
          </div>
          <div
            className='cursor-pointer mx-20 hover:opacity-60'
            onClick={() => {
              clickLikeButton(showPost.post.id);
            }}
          >
            <FavoriteIcon
              className='mr-3'
              color={
                showPost.likes.some((l) => l.userId === user?.id)
                  ? 'error'
                  : 'inherit'
              }
            />
            {showPost.likes.length}
          </div>
          <div className='cursor-pointer mx-20 hover:opacity-60'>
            <ShareIcon />
          </div>
        </div>
      </div>
      <div className='text-center'>
        <form onSubmit={(e) => onSubmit(e, showPost.post.id)}>
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
      <div>
        {showPost.commentsWithUsers.map((cu) => (
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
                  alt={cu.comment.text + 'picture'}
                />
              )}
            </div>
          </div>
        ))}
      </div>
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
