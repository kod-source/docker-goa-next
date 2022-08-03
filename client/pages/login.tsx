import { NextPage } from 'next';
import Head from 'next/head';
import PersonOutlineIcon from '@mui/icons-material/PersonOutline';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { useContext, useState } from 'react';
import axios, { AxiosError } from 'axios';
import { useRouter } from 'next/router';
import { AppContext } from './_app';
import { User } from '../lib/model/user';
import {
  Grid,
  Typography,
  Box,
  Paper,
  Link,
  TextField,
  CssBaseline,
  Button,
  Avatar,
} from '@mui/material';
import { isAxiosError, MyAxiosError } from '../lib/axios';

function Copyright(props: any) {
  return (
    <Typography
      variant='body2'
      color='text.secondary'
      align='center'
      {...props}
    >
      {'Copyright © '}
      <Link color='inherit' href='https://mui.com/'>
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const Login: NextPage = () => {
  const theme = createTheme();
  const { setUser } = useContext(AppContext);
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const res = await axios.post('http://localhost:3000/login', {
        email: email,
        password: password,
      });
      const token: string = res.data.token;
      localStorage.setItem('token', token);
      setUser(
        new User(
          res.data.user.id,
          res.data.user.name,
          res.data.user.email,
          res.data.user.password,
          new Date(res.data.user.created_at)
        )
      );
      router.push('/');
    } catch (e) {
      if (isAxiosError(e)) {
        const myAxiosError = e.response?.data as MyAxiosError;
        if (myAxiosError.message.length == 0) {
          return alert(e.message);
        }
        return alert(myAxiosError.message);
      }
    }
  };
  return (
    <div>
      <Head>
        <title>ログイン</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <ThemeProvider theme={theme}>
        <Grid container component='main' sx={{ height: '100vh' }}>
          <CssBaseline />
          <Grid
            item
            xs={false}
            sm={4}
            md={7}
            sx={{
              backgroundImage: 'url(https://source.unsplash.com/random)',
              backgroundRepeat: 'no-repeat',
              backgroundColor: (t) =>
                t.palette.mode === 'light'
                  ? t.palette.grey[50]
                  : t.palette.grey[900],
              backgroundSize: 'cover',
              backgroundPosition: 'center',
            }}
          />
          <Grid
            item
            xs={12}
            sm={8}
            md={5}
            component={Paper}
            elevation={6}
            square
          >
            <Box
              sx={{
                my: 8,
                mx: 4,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
              }}
            >
              {/* <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                <PersonOutlineIcon />
              </Avatar> */}
              <Typography component='h1' variant='h5'>
                Sign in
              </Typography>
              <form onSubmit={onSubmit}>
                <TextField
                  margin='normal'
                  required
                  fullWidth
                  label='Email Address'
                  name='email'
                  autoComplete='email'
                  autoFocus
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
                <TextField
                  margin='normal'
                  required
                  fullWidth
                  name='password'
                  label='Password'
                  type='password'
                  autoComplete='current-password'
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
                <Button
                  type='submit'
                  fullWidth
                  variant='contained'
                  className='bg-blue-400'
                  sx={{ mt: 3, mb: 2 }}
                >
                  Sign In
                </Button>
                <Button onClick={() => router.push('/sign_up')}>
                  新規アカウント作成はこちら
                </Button>
                <Copyright sx={{ mt: 5 }} />
              </form>
            </Box>
          </Grid>
        </Grid>
      </ThemeProvider>
    </div>
  );
};

export default Login;
