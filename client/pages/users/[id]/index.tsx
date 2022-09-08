import { Avatar, Button } from '@mui/material';
import { DateTime } from 'luxon';
import { NextPage, GetServerSideProps } from 'next';
import Image from 'next/image';
import { useRouter } from 'next/router';
import { FormEvent, useContext, useEffect, useState } from 'react';
import { User } from '../../../lib/model/user';
import { UserRepostiory } from '../../../lib/repository/user';
import { AppContext } from '../../_app';

interface Props {
  id: number;
}

const ShowUser: NextPage<Props> = ({ id }) => {
  const router = useRouter();
  const { user } = useContext(AppContext);
  const [showUser, setShowUser] = useState<User>();

  const fetchData = async () => {
    const user = await UserRepostiory.get(id);
    setShowUser(user);
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <p>ユーザーの詳細ページです</p>
    </div>
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
