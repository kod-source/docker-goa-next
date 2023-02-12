import { Avatar, Button, Modal } from "@mui/material";
import Box from "@mui/material/Box";
import { DateTime } from "luxon";
import Image from "next/image";
import React, { FC, FormEvent, useContext, useEffect, useState } from "react";

import { AppContext } from "../../pages/_app";
import { isAxiosError } from "../axios";
import { Comment, CommentWithUser } from "../model/comment";
import { PostWithUser, ShowPost } from "../model/post";
import { User } from "../model/user";
import { CommentRepository } from "../repository/comment";
import { PostRepository } from "../repository/post";
import { ConfirmationModal } from "./confirmationModal";
import { DetailModal } from "./detailModal";
import { Loading } from "./loading";
import { toStringlinefeed } from "./text";

interface Props {
    open: boolean;
    handleClose: () => void;
    postWithUser: PostWithUser;
    setPostsWithUser: React.Dispatch<React.SetStateAction<PostWithUser[]>>;
}

export const PostModal: FC<Props> = ({ open, handleClose, postWithUser, setPostsWithUser }) => {
    const style = {
        position: "absolute" as "absolute",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
        width: 800,
        height: 1000,
        bgcolor: "#222222",
        border: "2px solid #000",
        boxShadow: 24,
        p: 4,
    };

    const { user } = useContext(AppContext);
    const [commentsWithUser, setCommentsWithUser] = useState<CommentWithUser[]>();
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
        try {
            const commentWithUser = await CommentRepository.show(postWithUser.post.id);
            setCommentsWithUser(commentWithUser);
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
        const commentWithUser = await CommentRepository.create(postId, text, imagePath);
        setCommentsWithUser((old) => {
            if (!old) return [commentWithUser];
            return [commentWithUser, ...old];
        });
        setText("");
        setImagePath("");
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
            setCommentsWithUser((old) => {
                if (!old) return;
                const newCommentsWithUser = old.filter((cu) => cu.comment.id !== selectComment.id);
                return newCommentsWithUser;
            });
            setPostsWithUser((old) => {
                const newPosts = old.map((p) => {
                    if (p.post.id === postWithUser.post.id) {
                        return {
                            post: p.post,
                            user: p.user,
                            countLike: p.countLike,
                            countComment: p.countComment - 1,
                        };
                    }
                    return p;
                });
                return newPosts;
            });
        } else {
            await PostRepository.delete(postWithUser.post.id);
            handleClose();
        }
        setIsShowConfirmModal(false);
    };
    return (
        <Modal
            open={open}
            onClose={handleClose}
            aria-labelledby='modal-modal-title'
            aria-describedby='modal-modal-description'
        >
            <Box sx={style} className='overflow-scroll h-3/4'>
                <div className='my-2 border border-slate-600 p-5 rounded-md'>
                    <div className='flex'>
                        <Avatar
                            sx={{ width: 80, height: 80 }}
                            alt='投稿者'
                            src={
                                postWithUser.user.avatar ? postWithUser.user.avatar : "/avatar.png"
                            }
                        />
                        <div className='pt-5 mx-3'>
                            <p>{postWithUser.user.name}</p>
                            <div className='flex'>
                                <p>
                                    投稿日：
                                    {DateTime.fromJSDate(postWithUser.post.createdAt).toFormat(
                                        "yyyy年MM月dd日",
                                    )}
                                </p>
                                {postWithUser.post.createdAt.getTime() !==
                                    postWithUser.post.updatedAt.getTime() && (
                                    <p className='mx-5'>
                                        更新日：
                                        {DateTime.fromJSDate(postWithUser.post.updatedAt).toFormat(
                                            "yyyy年MM月dd日",
                                        )}
                                    </p>
                                )}
                            </div>
                        </div>
                        <div className='ml-auto'>
                            <Button
                                className='text-white'
                                onClick={(e) => {
                                    setIsMine(postWithUser.post.userId === user?.id);
                                    onClickDetail(e);
                                }}
                            >
                                :
                            </Button>
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
                                        alt={"post picture"}
                                    />
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
                                        src={cu.user.avatar ? cu.user.avatar : "/avatar.png"}
                                    />
                                    <div className='pt-5 mx-3'>
                                        <p>{cu.user.name}</p>
                                        <div className='flex'>
                                            <p>
                                                投稿日：
                                                {DateTime.fromJSDate(cu.comment.createdAt).toFormat(
                                                    "yyyy年MM月dd日",
                                                )}
                                            </p>
                                            {cu.comment.createdAt.getTime() !==
                                                cu.comment.updatedAt.getTime() && (
                                                <p className='mx-5'>
                                                    更新日：
                                                    {DateTime.fromJSDate(
                                                        cu.comment.updatedAt,
                                                    ).toFormat("yyyy年MM月dd日")}
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
                                            alt={postWithUser.post.title + "picture"}
                                        />
                                    )}
                                </div>
                            </div>
                        ))}
                    </div>
                ) : (
                    <Loading />
                )}
                {isShowDetailModal && (
                    <DetailModal
                        open={isShowDetailModal}
                        handleClose={() => setIsShowDetailModal(false)}
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
            </Box>
        </Modal>
    );
};
