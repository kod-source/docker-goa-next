import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import CommentIcon from "@mui/icons-material/Comment";
import FavoriteIcon from "@mui/icons-material/Favorite";
import ShareIcon from "@mui/icons-material/Share";
import { Avatar, Button } from "@mui/material";
import { DateTime } from "luxon";
import { NextPage, GetServerSideProps } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import { FormEvent, useContext, useEffect, useState } from "react";
import { CommentEditModal } from "../lib/components/commentEditModal";
import { ConfirmationModal } from "../lib/components/confirmationModal";
import { DetailModal } from "../lib/components/detailModal";
import { Loading } from "../lib/components/loading";
import { PostEditModal } from "../lib/components/postEditModal";
import { toStringlinefeed } from "../lib/components/text";
import { Comment } from "../lib/model/comment";
import { ShowPost } from "../lib/model/post";
import { CommentRepository } from "../lib/repository/comment";
import { LikeRepository } from "../lib/repository/like";
import { PostRepository } from "../lib/repository/post";
import { AppContext } from "./_app";

interface Props {
  id: number;
}

const PostShow: NextPage<Props> = ({ id }) => {
  const router = useRouter();
  const { user } = useContext(AppContext);
  const [showPost, setShowPost] = useState<ShowPost>();
  const [text, setText] = useState("");
  const [imagePath, setImagePath] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [widthAndHeightRate, setWidthAndHeightRate] = useState({
    width: "",
    height: "",
  });
  const [isShowDetailModal, setIsShowDetailModal] = useState(false);
  const [isShowConfirmModal, setIsShowConfirmModal] = useState(false);
  const [isShowUpdateModal, setIsShowUpdateModal] = useState(false);
  const [selectComment, setSelectComment] = useState<Comment>();
  const [isMine, setIsMine] = useState(false);

  const fetchData = async () => {
    const showPost = await PostRepository.show(id);
    setShowPost(showPost);
  };

  useEffect(() => {
    fetchData();
  }, []);

  const onSubmit = async (e: FormEvent<HTMLFormElement>, postId: number) => {
    setIsLoading(true);
    e.preventDefault();
    const commentWithUser = await CommentRepository.create(postId, text, imagePath);
    setShowPost((old) => {
      if (!old) return;
      return {
        post: old.post,
        user: old.user,
        likes: old.likes,
        commentsWithUsers: [commentWithUser, ...old.commentsWithUsers],
      };
    });
    setText("");
    setImagePath("");
    setIsLoading(false);
  };

  const clickLikeButton = async (postId: number) => {
    try {
      if (showPost?.likes.some((l) => l.userId === user?.id)) {
        await LikeRepository.delete(postId);
        setShowPost((old) => {
          if (!old) return;
          const filterLikes = old.likes.filter(
            (l) => !(l.userId === user?.id && l.postId === postId),
          );
          return {
            post: old.post,
            user: old.user,
            likes: filterLikes,
            commentsWithUsers: old.commentsWithUsers,
          };
        });
      } else {
        const like = await LikeRepository.create(postId);
        if (like.postId !== postId) {
          throw new Error("post_id unknow");
        }
        setShowPost((old) => {
          if (!old) return;
          return {
            post: old.post,
            user: old.user,
            likes: [...old.likes, like],
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

  const onClickDetail = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    const currentWidth = e.clientX;
    const currentHeight = e.clientY;
    setWidthAndHeightRate({
      width: String((currentWidth / window.innerWidth) * 100) + "%",
      height: String((currentHeight / window.innerHeight) * 100) + "%",
    });
    setIsShowDetailModal(true);
  };

  const onDelete = async () => {
    if (selectComment) {
      await CommentRepository.delete(selectComment.id);
      setShowPost((old) => {
        if (!old) return;
        const newCommentsWithUser = old.commentsWithUsers.filter(
          (cu) => cu.comment.id !== selectComment.id,
        );
        return {
          post: old.post,
          user: old.user,
          likes: old.likes,
          commentsWithUsers: newCommentsWithUser,
        };
      });
    } else {
      await PostRepository.delete(id);
      router.push("/");
    }
    setIsShowConfirmModal(false);
  };

  if (!showPost) return <Loading />;
  return (
    <>
      <div className='mx-auto w-3/5 '>
        <div className='my-5 flex'>
          <ArrowBackIcon
            className='mr-5 cursor-pointer hover:opacity-60'
            onClick={() => router.push("/")}
          />
          <h2>投稿</h2>
        </div>
        <div className='my-5 border border-slate-600 p-5 rounded-md cursor-pointer'>
          <div className='flex justify-center'>
            <div onClick={() => router.push(`users/${showPost.user.id}`)}>
              <Avatar
                sx={{ width: 80, height: 80 }}
                alt='投稿者'
                src={showPost.user.avatar ? showPost.user.avatar : "/avatar.png"}
              />
            </div>
            <div className='pt-5 mx-3'>
              <p>{showPost.user.name}</p>
              <div className='flex'>
                <p>
                  投稿日：
                  {DateTime.fromJSDate(showPost.post.createdAt).toFormat("yyyy年MM月dd日")}
                </p>
                {showPost.post.createdAt.getTime() !== showPost.post.updatedAt.getTime() && (
                  <p className='mx-5'>
                    更新日：
                    {DateTime.fromJSDate(showPost.post.updatedAt).toFormat("yyyy年MM月dd日")}
                  </p>
                )}
              </div>
            </div>
            <div className='ml-auto'>
              <Button
                className='text-white'
                onClick={(e) => {
                  setIsMine(showPost.user.id === user?.id);
                  onClickDetail(e);
                }}
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
                alt={showPost.post.title + "picture"}
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
                color={showPost.likes.some((l) => l.userId === user?.id) ? "error" : "inherit"}
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
                  <Image src={imagePath} width={500} height={500} alt={"post picture"} />
                  <div className='absolute left-[35%] bottom-[90%]'>
                    <Button onClick={() => setImagePath("")}>❌</Button>
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
                  {isLoading ? "アップロード中" : "返信"}
                </Button>
              </div>
            </div>
          </form>
        </div>
        <div>
          {showPost.commentsWithUsers.map((cu) => (
            <div key={cu.comment.id} className='my-2 border border-slate-600 p-5 rounded-md'>
              <div className='flex justify-center'>
                <Avatar
                  sx={{ width: 80, height: 80 }}
                  alt='投稿者'
                  src={cu.user.avatar ? cu.user.avatar : "/avatar.png"}
                />
                <div className='pt-5 mx-3'>
                  <p>{cu.user.name}</p>
                  <div className='flex'>
                    <p>
                      投稿日：
                      {DateTime.fromJSDate(cu.comment.createdAt).toFormat("yyyy年MM月dd日")}
                    </p>
                    {cu.comment.createdAt.getTime() !== cu.comment.updatedAt.getTime() && (
                      <p className='mx-5'>
                        更新日：
                        {DateTime.fromJSDate(cu.comment.updatedAt).toFormat("yyyy年MM月dd日")}
                      </p>
                    )}
                  </div>
                </div>
                <div className='ml-auto'>
                  <Button
                    className='text-white'
                    onClick={(e) => {
                      setIsMine(cu.user.id === user?.id);
                      setSelectComment(cu.comment);
                      onClickDetail(e);
                    }}
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
                    alt={cu.comment.text + "picture"}
                  />
                )}
              </div>
            </div>
          ))}
        </div>
      </div>
      {isShowDetailModal && (
        <DetailModal
          open={isShowDetailModal}
          handleClose={() => {
            setSelectComment(undefined);
            setIsShowDetailModal(false);
          }}
          widthRate={widthAndHeightRate.width}
          heightRate={widthAndHeightRate.height}
          onUpdateClick={() => {
            setIsShowDetailModal(false);
            setIsShowUpdateModal(true);
          }}
          onDeleteClick={() => {
            setIsShowDetailModal(false);
            setIsShowConfirmModal(true);
          }}
          isMyPost={isMine}
        />
      )}
      {isShowConfirmModal && (
        <ConfirmationModal
          open={isShowConfirmModal}
          handleClose={() => setIsShowConfirmModal(false)}
          text='削除してもよろしいでしょうか？'
          confirmInvoke={() => onDelete()}
        />
      )}
      {isShowUpdateModal && (
        <CommentEditModal
          open={isShowUpdateModal}
          handleClose={() => {
            setSelectComment(undefined);
            setIsShowUpdateModal(false);
          }}
          comment={selectComment}
          setComment={setSelectComment}
          showPost={showPost}
          setShowPost={setShowPost}
        />
      )}
    </>
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
