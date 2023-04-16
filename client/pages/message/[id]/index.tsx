import { Loading } from "lib/components/loading";
import { getRoomName } from "lib/function/getRoomName";
import { ShowRoom } from "lib/model/room";
import { IndexThread } from "lib/model/thread";
import { RoomRepository } from "lib/repository/room";
import { ThreadRepository } from "lib/repository/thread";
import { NextPage, GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { AppContext } from "pages/_app";
import React, { useCallback, useContext, useEffect, useState } from "react";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { MessageInput } from "lib/components/MessageInput";
import { Avatar } from "@mui/material";
import { isAxiosError } from "lib/axios";
import { DateTime } from "luxon";
import { toStringlinefeed } from "lib/components/text";

interface Props {
    roomID: number;
}

const ShowMessage: NextPage<Props> = ({ roomID }) => {
    const router = useRouter();
    const { user } = useContext(AppContext);
    const [showRoom, setShowRoom] = useState<ShowRoom>();
    const [indexThreads, setIndexThread] = useState<IndexThread[]>([]);
    const [nextID, setNextID] = useState<number | null>(0);
    const [isLoading, setIsLoading] = useState(false);
    const [againFetch, setAgainFetch] = useState(true);

    const fetchData = async () => {
        await RoomRepository.show(roomID).then((value) => setShowRoom(value));
    };

    const fetchThreadData = async () => {
        if (nextID == null) return;
        try {
            setIsLoading(true);
            const allThread = await ThreadRepository.getByRoom(roomID, nextID);
            setNextID(allThread.nextID);
            setIndexThread((old) => {
                if (nextID === 0) {
                    return allThread.indexThreads;
                }
                return [...allThread.indexThreads, ...old];
            });
            setIsLoading(false);
        } catch (e) {
            if (isAxiosError(e)) {
                const myAxiosError = e.response;
                if (myAxiosError?.status === 404) {
                    return;
                }
                return alert(myAxiosError?.statusText);
            }
        }
    };

    useEffect(() => {
        if (againFetch) {
            fetchThreadData();
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

    const onMessageSubmit = async (message: string, imgData: string): Promise<void> => {
        const threadUser = await ThreadRepository.create(roomID, message, imgData);
        setIndexThread((old) => [...old, { threadUser: threadUser, countContent: 0 }]);
    };

    if (!showRoom || !user) return <Loading />;
    return (
        <div className='mx-auto w-3/5'>
            <div className='mt-10 pb-5 flex border border-neutral-700 p-2'>
                <ArrowBackIcon
                    className='mr-5 cursor-pointer hover:opacity-60'
                    onClick={() => router.push("/message")}
                />
                <h2>メッセージ一覧</h2>
            </div>
            <div className='px-5 border border-neutral-700 border-t-0'>
                {showRoom.room.isGroup ? (
                    <div className='flex items-center'>
                        <Avatar src={showRoom.room.img} />
                        <p className='mx-5'>{showRoom.room.name}</p>
                        <div className='flex items-center ml-auto border border-neutral-700 rounded-md p-1 my-3'>
                            {showRoom.users
                                .filter((u) => u.id !== user.id)
                                .map((u) => (
                                    <div key={u.id} className='w-6'>
                                        <Avatar
                                            src={u.avatar}
                                            className={`object-cover rounded-full border-1 border-white transform`}
                                        />
                                    </div>
                                ))}
                            <p className='ml-5 opacity-70'>{showRoom.users.length}</p>
                        </div>
                    </div>
                ) : (
                    <div className='flex items-center'>
                        {showRoom.users.map((u) => {
                            if (u.id === user.id) {
                                return;
                            }
                            return (
                                <>
                                    <div
                                        onClick={() => router.push(`/users/${u.id}`)}
                                        className='cursor-pointer hover:opacity-60'
                                    >
                                        <Avatar src={u.avatar} />
                                    </div>
                                    <p className='mx-5'>{u.name}</p>
                                </>
                            );
                        })}
                    </div>
                )}
            </div>
            <div className='border border-neutral-700 border-t-0 border-b-0 h-[100px]'>
                {indexThreads.map((ih) => (
                    <div key={ih.threadUser.thread.id} className='flex m-5'>
                        <Avatar src={ih.threadUser.user.avatar} />
                        <div className='ml-5'>
                            <p>
                                {ih.threadUser.user.name}
                                <span className='ml-2'>
                                    {DateTime.fromJSDate(ih.threadUser.thread.createdAt).toFormat(
                                        "yyyy年MM月dd日 hh:mm",
                                    )}
                                </span>
                            </p>
                            <p>{toStringlinefeed(ih.threadUser.thread.text)}</p>
                            {ih.countContent && (
                                <div>
                                    <p>{ih.countContent}件の返信</p>
                                </div>
                            )}
                        </div>
                    </div>
                ))}
            </div>
            <div className='fixed bottom-0 w-3/5 mx-auto left-0 right-0 px-5 py-3'>
                <MessageInput
                    onMessageSubmit={onMessageSubmit}
                    placeholderMessage={`${getRoomName(showRoom, user)}へのメッセージ`}
                />
            </div>
        </div>
    );
};

export const getServerSideProps: GetServerSideProps = async (content) => {
    const { id } = content.query;
    return {
        props: {
            roomID: id,
        },
    };
};

export default ShowMessage;
