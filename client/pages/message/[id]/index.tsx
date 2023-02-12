import { ShowRoom } from "lib/model/room";
import { RoomRepository } from "lib/repository/room";
import { NextPage, GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { useContext, useEffect, useState } from "react";

interface Props {
    roomID: number;
}

const ShowMessage: NextPage<Props> = ({ roomID }) => {
    const router = useRouter();
    const [showRoom, setShowRoom] = useState<ShowRoom>();

    const fetchData = async () => {
        const showRoom = await RoomRepository.show(roomID);
        setShowRoom(showRoom);
    };

    useEffect(() => {
        fetchData();
    }, []);

    return (
        <>
            <div>hello</div>
        </>
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
