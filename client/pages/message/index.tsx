import { NextPage } from "next";
import Head from "next/head";
import { useState, useEffect, useContext, useCallback } from "react";
import { isAxiosError } from "axios";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { useRouter } from "next/router";
import { RoomRepository } from "lib/repository/room";
import { IndexRoom } from "lib/model/room";
import { AppContext } from "pages/_app";
import { Loading } from "lib/components/loading";
import { Avatar, Button } from "@mui/material";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { CreateRoomModal } from "lib/components/createRoomModal";

const Message: NextPage = () => {
  const router = useRouter();
  const { user } = useContext(AppContext);
  const [indexRooms, setIndexRooms] = useState<IndexRoom[] | []>([]);
  const [nextID, setNextID] = useState<number | null>(0);
  const [againFetch, setAgainFetch] = useState(true);
  const [isShowCreateRoomModal, setIsShowCreateRoomModal] = useState<boolean>(false);

  const fetchRoomDate = async () => {
    try {
      if (nextID == null) {
        return;
      }
      const allRooms = await RoomRepository.index(nextID);
      setIndexRooms((old) => {
        if (nextID == 0) {
          return allRooms.indexRooms;
        }
        return [...old, ...allRooms.indexRooms];
      });
      setNextID(allRooms.nextID);
    } catch (e) {
      if (isAxiosError(e)) {
        const myAxiosError = e.response;
        if (myAxiosError?.status === 404) {
          setIndexRooms([]);
          return;
        }
        return alert(myAxiosError?.statusText);
      }
    }
  };

  useEffect(() => {
    if (againFetch) {
      fetchRoomDate();
    }
    window.addEventListener("scroll", changeBottom);
    return () => window.removeEventListener("scroll", changeBottom);
  }, [againFetch]);

  const changeBottom = useCallback(() => {
    const bottomPosition = document.body.offsetHeight - (window.scrollY + window.innerHeight);
    if (bottomPosition < 0) {
      setAgainFetch(true);
      return;
    }
    setAgainFetch(false);
  }, []);

  if (!user) return <Loading />;
  return (
    <>
      <Head>
        <title>Message</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>
      <div className='mx-auto w-3/5 '>
        <div className='my-5 flex'>
          <ArrowBackIcon
            className='mr-5 cursor-pointer hover:opacity-60'
            onClick={() => router.push("/")}
          />
          <h2>ホーム</h2>
        </div>
        <div className='flex justify-between'>
          <h1 className='text-2xl font-bold'>メッセージ</h1>
          <div
            className='hover:cursor-pointer opacity-60'
            onClick={() => setIsShowCreateRoomModal(true)}
          >
            <AddCircleOutlineIcon />
          </div>
        </div>
        <div className='my-5'>
          {indexRooms.map((ir) => (
            <div
              key={ir.room.id}
              className='w-1/2 h-10 my-3 flex hover:opacity-60 cursor-pointer'
              onClick={() => router.push(`/message/${ir.room.id}`)}
            >
              <div>
                <Avatar
                  sx={{ width: 60, height: 60 }}
                  alt='投稿者'
                  src={ir.room.isGroup ? (ir.room.img ? ir.room.img : "/avatar.pgn") : ir.showImg}
                />
              </div>
              <div className='mx-5'>
                <p>
                  {ir.room.isGroup
                    ? `${ir.room.name}(${ir.countUser})`
                    : ir.room.name.split("/").map((name) => {
                        if (name !== user.name) return name;
                      })}
                </p>
                <p className='text-gray-400 opacity-80'>{ir.lastText}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
      <CreateRoomModal
        open={isShowCreateRoomModal}
        handleClose={() => setIsShowCreateRoomModal(false)}
      />
    </>
  );
};

export default Message;
