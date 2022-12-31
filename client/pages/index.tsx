import { Button } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import { useRouter } from "next/router";
import { FormEvent, useCallback, useContext, useEffect, useState } from "react";
import { isAxiosError, MyAxiosError } from "../lib/axios";
import { PostWithUser, SelectPost } from "../lib/model/post";
import { getEndPoint, getToken } from "../lib/token";
import { AppContext } from "./_app";
import { ConfirmationModal } from "lib/components/confirmationModal";
import { DetailModal } from "lib/components/detailModal";
import { Loading } from "lib/components/loading";
import { PostEditModal } from "lib/components/postEditModal";
import { ShowPost } from "lib/components/showPost";
import { LikeRepository } from "lib/repository/like";
import { PostRepository } from "lib/repository/post";
import styles from "styles/Home.module.css";

const Home: NextPage = () => {
  const { user } = useContext(AppContext);
  const router = useRouter();
  const [postsWithUser, setPostsWithUser] = useState<PostWithUser[]>([]);
  const [post, setPost] = useState<{ title: string; img: string }>({
    title: "",
    img: "",
  });
  const [isShowDetailModal, setIsShowDetailModal] = useState(false);
  const [widthAndHeightRate, setWidthAndHeightRate] = useState({
    width: "",
    height: "",
  });
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [isMyPost, setIsMyPost] = useState(false);
  const [showPostEditModal, setShowPostEditModal] = useState(false);
  const [selectPost, setSelectPost] = useState<SelectPost>({
    id: 0,
    title: "",
    img: "",
  });
  const [againFetch, setAgainFetch] = useState(true);
  const [nextID, setNextID] = useState<number | null>(0);
  const [isLoading, setIsLoading] = useState(false);
  const [myLikePostIds, setMyLikePostIds] = useState<number[]>([]);

  const logout = () => {
    localStorage.removeItem("token");
    router.push("/login");
  };

  const fetchPostData = async () => {
    if (nextID == null) return;
    setIsLoading(true);
    const postAllLimit = await PostRepository.index(nextID);
    setNextID(postAllLimit.nextId);
    setPostsWithUser((old) => {
      if (nextID === 0) {
        return postAllLimit.postsWithUsers;
      }
      return [...old, ...postAllLimit.postsWithUsers];
    });
    setIsLoading(false);
  };

  const fetchData = async () => {
    const myLikePostIds = await LikeRepository.getMyLike();
    setMyLikePostIds(myLikePostIds);
  };

  useEffect(() => {
    if (againFetch) {
      fetchPostData();
    }
    window.addEventListener("scroll", changeBottom);
    return () => window.removeEventListener("scroll", changeBottom);
  }, [againFetch]);

  useEffect(() => {
    fetchData();
  }, []);

  const changeBottom = useCallback(() => {
    const bottomPosition = document.body.offsetHeight - (window.scrollY + window.innerHeight);
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
      const postWithUser = await PostRepository.create(post.title, post.img);
      setPost({ title: "", img: "" });
      setPostsWithUser((old) => [postWithUser, ...old]);
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
      await PostRepository.delete(selectPost.id);
      setShowConfirmModal(false);
      setPostsWithUser(postsWithUser?.filter((p) => p.post.id !== selectPost.id));
    } catch (e) {
      if (e instanceof Error) {
        alert(e.message);
      }
    }
  };

  const clickLikeButton = async (postId: number) => {
    try {
      if (myLikePostIds.includes(postId)) {
        await LikeRepository.delete(postId);
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
                countComment: o.countComment,
              };
            }
            return o;
          });
          return newPosts;
        });
      } else {
        const like = await LikeRepository.create(postId);
        if (like.postId !== postId) {
          throw new Error("post_id unknow");
        }
        setMyLikePostIds((old) => [...old, postId]);
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postId) {
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

  const onClickDetail = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>, p: PostWithUser) => {
    const currentWidth = e.clientX;
    const currentHeight = e.clientY;
    setWidthAndHeightRate({
      width: String((currentWidth / window.innerWidth) * 100) + "%",
      height: String((currentHeight / window.innerHeight) * 100) + "%",
    });
    setIsShowDetailModal(true);
    setIsMyPost(p.post.userId === user?.id);
    setSelectPost({
      id: p.post.id,
      title: p.post.title,
      img: p.post.img,
    });
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
            <div
              className='cursor-pointer hover:opacity-60'
              onClick={() => router.push(`users/${user.id}`)}
            >
              <Avatar
                sx={{ width: 80, height: 80 }}
                className='mx-5'
                alt='投稿者'
                src={user.avatar ? user.avatar : "/avatar.png"}
              />
            </div>
            <label className='cursor-pointer'>
              <textarea
                className='w-96'
                autoFocus
                required
                value={post.title}
                onChange={(e) => setPost((old) => ({ title: e.target.value, img: old.img }))}
              />
            </label>
          </div>
          <div className='my-2'>
            {!!post.img && (
              <div className='relative'>
                <Image src={post.img} width={500} height={500} alt={"post picture"} />
                <div className='absolute left-[35%] bottom-[90%]'>
                  <Button onClick={() => setPost((old) => ({ title: old.title, img: "" }))}>
                    ❌
                  </Button>
                </div>
              </div>
            )}
          </div>
          <div className='mr-36'>
            <label className='cursor-pointer'>
              <Avatar src='/add_photo.jpg' className='m-auto' />
              <input type='file' className='hidden' accept='image/*' onChange={onChangeInputFile} />
            </label>
            <div className='text-right'>
              <Button type='submit'>投稿する</Button>
            </div>
          </div>
        </form>
      </div>
      <div className='mx-auto w-3/5'>
        {postsWithUser.map((p) => (
          <ShowPost
            key={p.post.id}
            postWithUser={p}
            setPostsWithUser={setPostsWithUser}
            myLikePostIds={myLikePostIds}
            clickLikeButton={clickLikeButton}
            onClickDetail={onClickDetail}
            onRouter={() => router.push(`users/${p.post.userId}`)}
          />
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
          setPostsWithUsers={setPostsWithUser}
        />
      )}
    </div>
  );
};

export default Home;
