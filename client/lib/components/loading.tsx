import React, { FC } from 'react';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';

export const Loading: FC = () => {
  return (
    <Box
      sx={{ display: 'flex' }}
      className='relative'
      style={{ top: '50%', left: '50%' }}
    >
      <CircularProgress color='inherit' />
    </Box>
  );
};
