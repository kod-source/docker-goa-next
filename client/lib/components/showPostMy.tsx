import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";
import { useRouter } from "next/router";
import React, { FC, useCallback, useContext, useEffect, useState } from "react";

import { AppContext } from "../../pages/_app";
import { PostWithUser, SelectPost } from "../model/post";
import { User, UserPostSelection } from "../model/user";
import { LikeRepository } from "../repository/like";
import { PostRepository } from "../repository/post";
import { ConfirmationModal } from "./confirmationModal";
import { DetailModal } from "./detailModal";
import { Loading } from "./loading";
import { PostEditModal } from "./postEditModal";
import { ShowPost } from "./showPost";

interface Props {
  value: UserPostSelection;
  setValue: React.Dispatch<React.SetStateAction<UserPostSelection>>;
  showUser: User;
  setShowUser: React.Dispatch<React.SetStateAction<User | undefined>>;
  myLikePostIds: number[];
  setMyLikePostIds: React.Dispatch<React.SetStateAction<number[]>>;
}

export const ShowPostMy: FC<Props> = ({
  value,
  setValue,
  showUser,
  setShowUser,
  myLikePostIds,
  setMyLikePostIds,
}) => {
  const { user } = useContext(AppContext);
  const router = useRouter();
  const [nextID, setNextID] = useState<number | null>(0);
  const [postsWithUser, setPostsWithUser] = useState<PostWithUser[]>([]);
  const [againFetch, setAgainFetch] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [isShowDetailModal, setIsShowDetailModal] = useState(false);
  const [widthAndHeightRate, setWidthAndHeightRate] = useState({
    width: "",
    height: "",
  });
  const [isMyPost, setIsMyPost] = useState(false);
  const [selectPost, setSelectPost] = useState<SelectPost>({
    id: 0,
    title: "",
    img: "",
  });
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [showPostEditModal, setShowPostEditModal] = useState(false);

  const fetchData = async () => {
    if (nextID == null) return;
    setIsLoading(true);
    const allPostLimit =
      value === UserPostSelection.My
        ? await PostRepository.showPostMy(nextID, showUser.id)
        : value === UserPostSelection.Media
        ? await PostRepository.showPostMedia(nextID, showUser.id)
        : await PostRepository.showPostLike(nextID, showUser.id);
    setNextID(allPostLimit.nextId);
    setPostsWithUser((old) => {
      if (nextID === 0) {
        return allPostLimit.postsWithUsers;
      }
      return [...old, ...allPostLimit.postsWithUsers];
    });
    setIsLoading(false);
  };

  useEffect(() => {
    if (againFetch) {
      fetchData();
    }
    window.addEventListener("scroll", changeBottom);
    return () => window.removeEventListener("scroll", changeBottom);
  }, [value, againFetch, showUser, nextID]);

  const changeBottom = useCallback(() => {
    const bottomPosition = document.body.offsetHeight - (window.scrollY + window.innerHeight);
    if (bottomPosition < 0) {
      setAgainFetch(true);
      return;
    }
    setAgainFetch(false);
  }, []);

  const clickLikeButton = async (postID: number) => {
    try {
      if (myLikePostIds.includes(postID)) {
        await LikeRepository.delete(postID);
        setMyLikePostIds((old) => {
          return old.filter((l) => l !== postID);
        });
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postID) {
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
        const like = await LikeRepository.create(postID);
        if (like.postId !== postID) {
          throw new Error("post_id unknow");
        }
        setMyLikePostIds((old) => [...old, postID]);
        setPostsWithUser((old) => {
          const newPosts = old.map((o) => {
            if (o.post.id === postID) {
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

  return (
    <div>
      <div>
        {postsWithUser.map((p) => (
          <ShowPost
            key={p.post.id}
            postWithUser={p}
            setPostsWithUser={setPostsWithUser}
            myLikePostIds={myLikePostIds}
            clickLikeButton={clickLikeButton}
            onClickDetail={onClickDetail}
            onRouter={() => {
              if (showUser.id === p.post.userId) {
                return;
              }
              setValue(UserPostSelection.My);
              setNextID(0);
              setShowUser((old) => {
                if (!old) return;
                return new User(
                  p.post.userId,
                  p.user.name,
                  old.email,
                  old.createdAt,
                  p.user.avatar,
                );
              });
              router.push(`${p.post.userId}`);
            }}
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
