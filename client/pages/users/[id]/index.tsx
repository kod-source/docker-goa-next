import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import EmailIcon from "@mui/icons-material/Email";
import TabContext from "@mui/lab/TabContext";
import TabList from "@mui/lab/TabList";
import TabPanel from "@mui/lab/TabPanel";
import { Avatar, Box, Button } from "@mui/material";
import Tab from "@mui/material/Tab";
import { DateTime } from "luxon";
import { NextPage, GetServerSideProps } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import { FormEvent, useContext, useEffect, useState } from "react";

import { Loading } from "../../../lib/components/loading";
import { ShowPostMy } from "../../../lib/components/showPostMy";
import { User, UserPostSelection } from "../../../lib/model/user";
import { LikeRepository } from "../../../lib/repository/like";
import { UserRepostiory } from "../../../lib/repository/user";
import { AppContext } from "../../_app";
import { isAxiosError } from "lib/axios";
import { RoomRepository } from "lib/repository/room";

interface Props {
    id: number;
}

const ShowUser: NextPage<Props> = ({ id }) => {
    const router = useRouter();
    const { user } = useContext(AppContext);
    const [showUser, setShowUser] = useState<User>();
    const [value, setValue] = useState<UserPostSelection>(UserPostSelection.My);
    const [myLikePostIds, setMyLikePostIds] = useState<number[]>([]);

    const fetchData = async () => {
        const user = await UserRepostiory.get(id);
        setShowUser(user);
        const myLikePostIds = await LikeRepository.getMyLike();
        setMyLikePostIds(myLikePostIds);
    };

    useEffect(() => {
        fetchData();
    }, [id]);

    const handleChange = (e: React.SyntheticEvent, newValue: UserPostSelection) => {
        setValue(newValue);
    };

    const createDirectMessageRoom = async (myUserId: number, id: number) => {
        if (!showUser || !user) return;
        try {
            const room = await RoomRepository.exists(id);
            router.push(`/message/${room.id}`);
        } catch (e) {
            if (isAxiosError(e)) {
                const myAxiosError = e.response;
                if (myAxiosError?.status === 404) {
                    try {
                        const showRoom = await RoomRepository.create(
                            `${user.name}/${showUser.name}`,
                            false,
                            [myUserId, id],
                        );
                        router.push(`message/${showRoom.room.id}`);
                    } catch (e) {
                        if (e instanceof Error) {
                            return alert(e.message);
                        }
                    }
                    return;
                }
                return alert(myAxiosError?.statusText);
            }
        }
    };

    if (!showUser || !user) return <Loading />;
    return (
        <>
            <div className='mx-auto w-3/5'>
                <div className='my-5 flex'>
                    <ArrowBackIcon
                        className='mr-5 cursor-pointer hover:opacity-60'
                        onClick={() => router.push("/")}
                    />
                    <h2>{showUser.name}</h2>
                </div>
                <div>
                    <Avatar
                        sx={{ width: 300, height: 300 }}
                        alt='投稿者'
                        className='mx-auto'
                        src={showUser.avatar ? showUser.avatar : "/avatar.png"}
                    />
                    {user.id != id && (
                        <div className='text-right'>
                            <EmailIcon
                                className='hover:opacity-70 cursor-pointer'
                                onClick={() => createDirectMessageRoom(user.id, showUser.id)}
                            />
                        </div>
                    )}
                </div>
                <Box sx={{ width: "100%", typography: "body1" }}>
                    <TabContext value={value}>
                        <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                            <TabList onChange={handleChange} aria-label='lab API tabs example'>
                                <Tab
                                    label='投稿'
                                    value={UserPostSelection.My}
                                    className='text-gray-500 cursor-pointer hover:opacity-60'
                                />
                                <Tab
                                    label='メディア'
                                    value={UserPostSelection.Media}
                                    className='text-gray-500 cursor-pointer hover:opacity-60'
                                />
                                <Tab
                                    label='いいね'
                                    value={UserPostSelection.Like}
                                    className='text-gray-500 cursor-pointer hover:opacity-60'
                                />
                            </TabList>
                        </Box>
                        <TabPanel value={UserPostSelection.My}>
                            <ShowPostMy
                                value={value}
                                setValue={setValue}
                                showUser={showUser}
                                setShowUser={setShowUser}
                                myLikePostIds={myLikePostIds}
                                setMyLikePostIds={setMyLikePostIds}
                            />
                        </TabPanel>
                        <TabPanel value={UserPostSelection.Media}>
                            <ShowPostMy
                                value={value}
                                setValue={setValue}
                                showUser={showUser}
                                setShowUser={setShowUser}
                                myLikePostIds={myLikePostIds}
                                setMyLikePostIds={setMyLikePostIds}
                            />
                        </TabPanel>
                        <TabPanel value={UserPostSelection.Like}>
                            <ShowPostMy
                                value={value}
                                setValue={setValue}
                                showUser={showUser}
                                setShowUser={setShowUser}
                                myLikePostIds={myLikePostIds}
                                setMyLikePostIds={setMyLikePostIds}
                            />
                        </TabPanel>
                    </TabContext>
                </Box>
            </div>
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

export default ShowUser;
