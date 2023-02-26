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
    const [message, setMessage] = useState("");

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

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setMessage("");
        alert("送信");
    };

    const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (e.key === "Enter" && e.metaKey) {
            onSubmit(e as any);
        }
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
            <div className='flex flex-col bg-gray-200 rounded-lg p-2'>
                <form onSubmit={onSubmit}>
                    <textarea
                        className='resize-none w-full h-auto p-2 border rounded-md'
                        style={{
                            height: `${Math.max(60, message.split("\n").length * 20)}px`,
                            maxHeight: "500px",
                        }}
                        placeholder={`${getRoomName(showRoom, user)}へのメッセージ`}
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        onKeyDown={handleKeyDown}
                    />
                    <button
                        className='bg-blue-500 text-white rounded-lg py-2 mt-2 hover:bg-blue-700'
                        type='submit'
                    >
                        Send
                    </button>
                </form>
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
