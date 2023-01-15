import {
  Avatar,
  Button,
  Modal,
  TextField,
  IconButton,
  Box,
  InputLabel,
  Select,
  OutlinedInput,
  Chip,
  MenuItem,
  Theme,
  useTheme,
  Checkbox,
  FormControlLabel,
} from "@mui/material";
import { User } from "lib/model/user";
import { UserRepostiory } from "lib/repository/user";
import React, { FC, FormEvent, useContext, useEffect, useState } from "react";

import { AppContext } from "../../pages/_app";

interface Props {
  open: boolean;
  handleClose: () => void;
}

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
};

export const CreateRoomModal: FC<Props> = ({ open, handleClose }) => {
  const style = {
    position: "absolute" as "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 800,
    height: 1000,
    bgcolor: "#222222",
    border: "2px solid #000",
    boxShadow: 24,
    p: 4,
  };

  function getStyles(id: number, selectIDs: readonly number[], theme: Theme) {
    return {
      fontWeight: selectIDs.includes(id)
        ? theme.typography.fontWeightRegular
        : theme.typography.fontWeightMedium,
    };
  }

  const { user } = useContext(AppContext);
  const theme = useTheme();
  const [selectUsers, setSelectUsers] = useState<User[]>([]);
  const [values, setValues] = useState<{
    name: string;
    isGroup: boolean;
    userIDs: number[];
    img: string | null;
  }>({ name: "", isGroup: true, userIDs: [], img: null });
  const [isLoading, setIsLoading] = useState(false);
  const [isMyJoin, setIsMyJoin] = useState(true);

  const fetchData = async () => {
    const users = await UserRepostiory.index();
    setSelectUsers(users);
  };

  useEffect(() => {
    fetchData();
  }, []);

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {};

  const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();
      reader.onload = (e: ProgressEvent<FileReader>) => {
        if (!e.target) return null;
        setValues((old) => ({ ...old, img: e.target?.result as string }));
      };
      reader.readAsDataURL(file);
    }
  };

  const ShowUser: FC<{ userID: number }> = ({ userID }): JSX.Element => {
    const user = selectUsers.find((u) => u.id === userID);
    if (!user) return <></>;

    return (
      <div className='flex m-auto'>
        <Avatar
          sx={{ width: 40, height: 40 }}
          alt='user img'
          src={user.avatar ? user.avatar : "/avatar.png"}
        />
        <p className='text-white m-auto pl-3 text-base'>{user.name}</p>
      </div>
    );
  };

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby='modal-modal-title'
      aria-describedby='modal-modal-description'
    >
      <Box sx={style} className='overflow-scroll h-3/4'>
        <h1>グループ作成</h1>
        <div className='text-center'>
          <form onSubmit={(e) => onSubmit(e)}>
            <Box textAlign='center'>
              <IconButton>
                <label className='cursor-pointer'>
                  <Avatar
                    sx={{ width: 100, height: 100 }}
                    alt='My Profile Image'
                    src={values.img ? values.img : "/avatar.png"}
                  />
                  <input
                    type='file'
                    className='hidden'
                    accept='image/*'
                    onChange={onChangeInputFile}
                  />
                </label>
              </IconButton>
            </Box>
            <FormControlLabel
              control={
                <Checkbox
                  checked={isMyJoin}
                  onChange={(e) => setIsMyJoin(e.target.checked)}
                  inputProps={{ "aria-label": "controlled" }}
                />
              }
              label='自分も参加する'
            />
            <TextField
              margin='normal'
              required
              fullWidth
              label='グループ名'
              name='name'
              autoComplete='name'
              autoFocus
              value={values.name}
              onChange={(e) => setValues((old) => ({ ...old, name: e.target.value }))}
            />
            <InputLabel id='demo-multiple-chip-label' className='text-white'>
              ユーザーを選択
            </InputLabel>
            <Select
              labelId='demo-multiple-chip-label'
              id='demo-multiple-chip'
              multiple
              fullWidth
              value={values.userIDs}
              onChange={(e) => {
                const values = e.target.value;
                if (typeof values !== "string") {
                  return setValues((old) => ({
                    ...old,
                    userIDs: values,
                  }));
                }
              }}
              input={<OutlinedInput id='select-multiple-chip' label='Chip' />}
              renderValue={(selected) => (
                <Box sx={{ display: "flex", flexWrap: "wrap", gap: 0.5 }}>
                  {selected.map((value) => (
                    <Chip key={value} label={<ShowUser userID={value} />} />
                  ))}
                </Box>
              )}
              MenuProps={MenuProps}
            >
              {selectUsers.map((u) => (
                <MenuItem key={u.id} value={u.id} style={getStyles(u.id, values.userIDs, theme)}>
                  <Avatar
                    sx={{ width: 40, height: 40 }}
                    alt='user img'
                    src={u.avatar ? u.avatar : "/avatar.png"}
                  />
                  {u.name}
                </MenuItem>
              ))}
            </Select>
            <div className='text-right'>
              <Button type='submit' disabled={isLoading}>
                {isLoading ? "作成中" : "作成する"}
              </Button>
            </div>
          </form>
        </div>
      </Box>
    </Modal>
  );
};
