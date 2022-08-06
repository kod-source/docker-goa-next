import React, { FC } from 'react';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';

export const Loading: FC = () => {
  return (
    <Box sx={{ display: 'flex' }} className='absolute top-1/2 left-1/2'>
      <CircularProgress color='inherit' />
    </Box>
  );
};
