import { Button, Modal, selectClasses } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Image from "next/image";
import React, { FC, FormEvent, useState } from "react";
import { Comment } from "../model/comment";
import { Post, ShowPost } from "../model/post";
import { CommentRepository } from "../repository/comment";
import { PostRepository } from "../repository/post";

interface Props {
  open: boolean;
  handleClose: () => void;
  comment?: Comment;
  setComment: React.Dispatch<React.SetStateAction<Comment | undefined>>;
  showPost: ShowPost;
  setShowPost: React.Dispatch<React.SetStateAction<ShowPost | undefined>>;
}

export const CommentEditModal: FC<Props> = (props) => {
  const style = {
    position: "absolute" as "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "background.paper",
    border: "2px solid #000",
    boxShadow: 24,
    p: 4,
  };
  const [isUpdating, setIsUpdating] = useState(false);
  const [selectPost, setSelectPost] = useState(props.showPost.post);

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsUpdating(true);
    if (!props.comment) {
      const post = await PostRepository.update(selectPost.id, selectPost.title, selectPost.img);
      props.setShowPost((old) => {
        if (!old) return;
        return {
          post: post,
          user: old.user,
          likes: old.likes,
          commentsWithUsers: old.commentsWithUsers,
        };
      });
    } else {
      const comment = await CommentRepository.update(
        props.comment.id,
        props.comment.text,
        props.comment.img,
      );
      props.setComment(comment);
      props.setShowPost((old) => {
        if (!old) return;
        const newCommentsWithUser = old.commentsWithUsers.map((cu) => {
          if (cu.comment.id === props.comment?.id) {
            return {
              comment: comment,
              user: cu.user,
            };
          }
          return cu;
        });
        return {
          post: old.post,
          user: old.user,
          likes: old.likes,
          commentsWithUsers: newCommentsWithUser,
        };
      });
    }
    setIsUpdating(false);
    props.handleClose();
  };

  const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e: ProgressEvent<FileReader>) => {
        if (!e.target) return null;
        if (!props.comment) {
          setSelectPost(
            (old) =>
              new Post(
                old.id,
                old.userId,
                old.title,
                old.createdAt,
                old.updatedAt,
                e.target?.result as string,
              ),
          );
        } else {
          props.setComment((old) => {
            if (!old) return;
            return {
              id: old.id,
              postId: old.postId,
              userId: old.userId,
              text: old.text,
              createdAt: old.createdAt,
              updatedAt: old.updatedAt,
              img: e.target?.result as string,
            };
          });
        }
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
                  value={props.comment ? props.comment.text : selectPost.title}
                  onChange={(e) => {
                    if (!props.comment) {
                      return setSelectPost(
                        (old) =>
                          new Post(
                            old.id,
                            old.userId,
                            e.target.value,
                            old.createdAt,
                            old.updatedAt,
                            old.img,
                          ),
                      );
                    }
                    props.setComment((old) => {
                      if (!old) return;
                      return {
                        id: old.id,
                        postId: old.postId,
                        userId: old.userId,
                        text: e.target.value,
                        createdAt: old.createdAt,
                        updatedAt: old.updatedAt,
                        img: old.img,
                      };
                    });
                  }}
                />
              </label>
            </div>
            <div className='my-2'>
              {((props.comment && !!props.comment?.img) ||
                (!!selectPost.img && !props.comment)) && (
                <div className='relative'>
                  <Image
                    src={
                      props.comment
                        ? props.comment.img
                          ? props.comment.img
                          : ""
                        : selectPost.img
                        ? selectPost.img
                        : ""
                    }
                    width={500}
                    height={500}
                    alt={"picture"}
                  />
                  <div className='absolute left-[35%] bottom-[90%]'>
                    <Button
                      onClick={() => {
                        if (!props.comment) {
                          return setSelectPost(
                            (old) =>
                              new Post(
                                old.id,
                                old.userId,
                                old.title,
                                old.createdAt,
                                old.updatedAt,
                                "",
                              ),
                          );
                        }
                        props.setComment((old) => {
                          if (!old) return;
                          return {
                            id: old.id,
                            postId: old.postId,
                            userId: old.userId,
                            text: old.text,
                            createdAt: old.createdAt,
                            updatedAt: old.updatedAt,
                            img: "",
                          };
                        });
                      }}
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
                  {isUpdating ? "更新中" : "更新する"}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </Box>
    </Modal>
  );
};
