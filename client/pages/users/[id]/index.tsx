import { Avatar, Button } from '@mui/material';
import axios from 'axios';
import { DateTime } from 'luxon';
import { NextPage, GetServerSideProps } from 'next';
import Image from 'next/image';
import { useRouter } from 'next/router';
import { FormEvent, useContext, useEffect, useState } from 'react';
import { User } from '../../../lib/model/user';
import { getEndPoint, getToken } from '../../../lib/token';
import { getUser } from '../../api/user';
import { AppContext } from '../../_app';

interface Props {
  id: number;
  showUser: User;
}

const ShowUser: NextPage<Props> = ({ id, showUser }) => {
  console.log(showUser);
  const router = useRouter();
  const { user } = useContext(AppContext);

  return (
    <div>
      <p>ユーザーの詳細ページです</p>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async (content) => {
  const { id } = content.query;
  //   const user = await getUser(Number(id));
  //   console.log(user);
  //   const response = await fetch(
  //     `${process.env.NEXT_PUBLIC_END_POINT}/users/${id}`
  //   );
  const token = await localStorage.getItem('token')
  console.log(token);
  const response = await fetch('http://localhost:3000/users/6', {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  const data = await response.json();
  console.log(data);
  return {
    props: {
      id: id,
      //   showUser: user,
    },
  };
};

export default ShowUser;
