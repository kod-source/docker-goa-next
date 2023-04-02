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
            <div className='my-5 flex'>
                <ArrowBackIcon
                    className='mr-5 cursor-pointer hover:opacity-60'
                    onClick={() => router.push("/message")}
                />
                <h2>メッセージ一覧</h2>
            </div>
            <div>
                {showRoom.room.isGroup ? (
                    <div className='flex items-center'>
                        <Avatar src={showRoom.room.img} />
                        <p className='mx-5'>{showRoom.room.name}</p>
                        <div className='flex items-center ml-auto border border-neutral-700 rounded-md p-2'>
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
            <div>{/* メッセージの表示 */}</div>
            <div className='flex flex-col rounded-lg my-2 w-3/5 absolute bottom-10'>
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
