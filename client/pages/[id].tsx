import axios from 'axios';
import { NextPage, GetServerSideProps } from 'next';
import { useRouter } from 'next/router';
import { useState } from 'react';
import { ShowPost } from '../lib/model/post';

interface Props {
  id: number;
}

const PostShow: NextPage<Props> = ({ id }) => {
  const router = useRouter();
  const [showPost, setShowPost] = useState<ShowPost>()

  const fetchData = async () => {
    const res = await axios.get("https")
  }

  return (
    <div>
      <h1>詳細ページです</h1>
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

export default PostShow;
