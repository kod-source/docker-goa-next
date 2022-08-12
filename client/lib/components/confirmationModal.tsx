import React, { FC } from 'react';
import Box from '@mui/material/Box';
import { Button, Modal, Typography } from '@mui/material';

interface Props {
  open: boolean;
  handleClose: () => void;
  text: string;
  confirmInvoke: () => void;
  widthRate?: string;
  heightRate?: string;
}

export const ConfirmationModal: FC<Props> = (props) => {
  const style = {
    position: 'absolute' as 'absolute',
    top: props.heightRate ? props.heightRate : '50%',
    left: props.widthRate ? props.widthRate : '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
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
      <Box sx={style} className='text-center'>
        <p style={{ color: 'black' }}>{props.text}</p>
        <div className='mt-5'>
          <Button onClick={props.handleClose}>キャンセル</Button>
          <Button onClick={props.confirmInvoke}>OK</Button>
        </div>
      </Box>
    </Modal>
  );
};
