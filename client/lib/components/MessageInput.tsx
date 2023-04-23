import { FC, useState } from "react";
import Image from "next/image";
import SendIcon from "@mui/icons-material/Send";
import AddIcon from "@mui/icons-material/Add";
import { Button } from "@mui/material";

interface Props {
    onMessageSubmit: (message: string, imgData: string) => Promise<void>;
    placeholderMessage: string;
}

export const MessageInput: FC<Props> = (props) => {
    const [message, setMessage] = useState("");
    const [imgData, setImgData] = useState("");

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!message && !imgData) {
            return;
        }
        await props.onMessageSubmit(message, imgData);
        setMessage("");
        setImgData("");
    };

    const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (e.key === "Enter" && e.metaKey && (!!message || !!imgData)) {
            props.onMessageSubmit(message, imgData);
            setMessage("");
            setImgData("");
        }
    };

    const onChangeInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            const file = e.target.files[0];
            const reader = new FileReader();
            reader.onload = (e: ProgressEvent<FileReader>) => {
                if (!e.target) return null;
                setImgData(e.target.result as string);
            };
            reader.readAsDataURL(file);
        }
    };

    return (
        <form onSubmit={onSubmit} className='border rounded-lg p-1 bg-[#3B3B3B]'>
            <textarea
                className='resize-none w-full h-auto py-2 rounded-md border-none outline-hidden'
                style={{
                    height: `${Math.max(60, message.split("\n").length * 20)}px`,
                    maxHeight: "500px",
                }}
                placeholder={props.placeholderMessage}
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={handleKeyDown}
            />
            {!!imgData && (
                <div className='relative'>
                    <Image src={imgData} width={100} height={100} alt={"post picture"} />
                    <div className='absolute left-[8%] bottom-[65%]'>
                        <Button onClick={() => setImgData("")}>❌</Button>
                    </div>
                </div>
            )}
            <div className='flex'>
                <label className='cursor-pointer hover:opacity-60 bg-gray-600 rounded-full'>
                    <AddIcon />
                    <input
                        type='file'
                        className='hidden'
                        accept='image/*'
                        onChange={onChangeInputFile}
                    />
                </label>
                <button
                    className={`rounded-lg ml-auto ${
                        !!message || !!imgData
                            ? "bg-green-500 text-white hover:cursor-pointer hover:opacity-60"
                            : "text-gray-500"
                    }`}
                    type='submit'
                    disabled={!message && !imgData}
                >
                    <SendIcon />
                </button>
            </div>
        </form>
    );
};
