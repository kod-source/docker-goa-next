import { Avatar, Box, Button } from '@mui/material';
import { DateTime } from 'luxon';
import { NextPage, GetServerSideProps } from 'next';
import Image from 'next/image';
import { useRouter } from 'next/router';
import { FormEvent, useContext, useEffect, useState } from 'react';
import { User, UserPostSelection } from '../../../lib/model/user';
import { UserRepostiory } from '../../../lib/repository/user';
import { AppContext } from '../../_app';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { Loading } from '../../../lib/components/loading';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import { LikeRepository } from '../../../lib/repository/like';
import { ShowPostMy } from '../../../lib/components/showPostMy';

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
  }, []);

  const handleChange = (
    e: React.SyntheticEvent,
    newValue: UserPostSelection
  ) => {
    setValue(newValue);
  };

  if (!showUser) return <Loading />;
  return (
    <>
      <div className='mx-auto w-3/5'>
        <div className='my-5 flex'>
          <ArrowBackIcon
            className='mr-5 cursor-pointer hover:opacity-60'
            onClick={() => router.push('/')}
          />
          <h2>{showUser.name}</h2>
        </div>
        <div>
          <Avatar
            sx={{ width: 300, height: 300 }}
            alt='投稿者'
            className='mx-auto'
            src={showUser.avatar ? showUser.avatar : '/avatar.png'}
          />
        </div>
        <Box sx={{ width: '100%', typography: 'body1' }}>
          <TabContext value={value}>
            <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
              <TabList
                onChange={handleChange}
                aria-label='lab API tabs example'
              >
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
                showUser={showUser}
                myLikePostIds={myLikePostIds}
                setMyLikePostIds={setMyLikePostIds}
              />
            </TabPanel>
            <TabPanel value={UserPostSelection.Media}>
              <ShowPostMy
                value={value}
                showUser={showUser}
                myLikePostIds={myLikePostIds}
                setMyLikePostIds={setMyLikePostIds}
              />
            </TabPanel>
            <TabPanel value={UserPostSelection.Like}>
              <ShowPostMy
                value={value}
                showUser={showUser}
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
