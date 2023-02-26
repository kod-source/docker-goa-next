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
    const [againFetch, setAgainFetch] = useState(false);

    const fetchData = async () => {
        await RoomRepository.show(roomID).then((value) => setShowRoom(value));
    };

    const fetchThreadData = async () => {
        if (nextID == null) return;
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

    const onMessageSubmit = (message: string, imgData: string): void => {
        console.log("message", message);
        console.log("img", imgData);
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
                hello
                {/* ルーム情報 */}
            </div>
            <div>{/* メッセージの表示 */}</div>
            <div className='flex flex-col rounded-lg my-2'>
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
