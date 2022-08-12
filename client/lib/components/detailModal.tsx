import React, { FC, useState } from 'react';
import Box from '@mui/material/Box';
import { Button, Modal, Typography } from '@mui/material';

interface Props {
  open: boolean;
  handleClose: () => void;
  widthRate: string;
  heightRate: string;
  onDeleteClick: () => void;
  isMyPost: boolean;
}

export const DetailModal: FC<Props> = (props) => {
  const style = {
    position: 'absolute' as 'absolute',
    top: props.heightRate,
    left: props.widthRate,
    transform: 'translate(-50%, -50%)',
    width: 200,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };

  return (
    <Modal
      open={props.open}
      onClose={props.handleClose}
      aria-labelledby='modal-modal-title'
      aria-describedby='modal-modal-description'
    >
      <Box sx={style} className='text-gray-900'>
        {props.isMyPost && (
          <>
            <div>
              <Button>編集</Button>
            </div>
            <div>
              <Button onClick={props.onDeleteClick}>削除</Button>
            </div>
          </>
        )}
        <div>
          <Button>シェアする</Button>
        </div>
      </Box>
    </Modal>
  );
};
