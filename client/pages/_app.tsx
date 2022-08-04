import '../styles/globals.css';
import type { AppProps } from 'next/app';
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';
import { User } from '../lib/model/user';
import 'tailwindcss/tailwind.css';

export const AppContext = React.createContext(
  {} as {
    user: User | null;
    setUser: React.Dispatch<React.SetStateAction<User | null>>;
  }
);

function MyApp({ Component, pageProps }: AppProps) {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const path = router.pathname;
  let isFirst = true;

  const fetchData = async () => {
    if (path === '/login' || '/sign_up') {
      return;
    }
    const token = localStorage.getItem('token');
    if (!token) {
      return router.push('/login');
    }
    try {
      const res = await axios.get('http://localhost:3000/current_user', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setUser(
        new User(
          res.data.id,
          res.data.name,
          res.data.email,
          res.data.password,
          new Date(res.data.created_at),
          res.data.user.avatar
        )
      );
    } catch {
      if (isFirst) {
        alert('tokenの認証が切れました。再度ログインしてください。');
        localStorage.removeItem('token');
        router.push('/login');
      }
    }
    isFirst = false;
  };

  useEffect(() => {
    fetchData();
  }, []);
  return (
    <AppContext.Provider value={{ user, setUser }}>
      <Component {...pageProps} />
    </AppContext.Provider>
  );
}

export default MyApp;
