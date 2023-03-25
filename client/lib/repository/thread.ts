import { AllThread, IndexThread, Thread, ThreadUser } from "lib/model/thread";
import { User } from "lib/model/user";
import { asyncApiClient } from "../axios";
import { IndexRoom, Room } from "../model/room";

export const ThreadRepository = {
    getByRoom: async (roomID: number, nextID: number): Promise<AllThread> => {
        const apiClient = await asyncApiClient.create();
        const res = await apiClient.get(`threads/room/${roomID}?next_id=${nextID}`);
        const indexThreads: IndexThread[] = res.data.map((d: any): IndexThread => {
            const thread = d.thread_user.thread;
            const user = d.thread_user.user;
            return {
                threadUser: {
                    thread: new Thread(
                        thread.id,
                        thread.room_id,
                        thread.user_id,
                        thread.text,
                        new Date(thread.created_at),
                        new Date(thread.updated_at),
                    ),
                    user: {
                        id: user.id,
                        name: user.name,
                        createdAt: user.created_at,
                        avatar: user.avatar,
                    },
                },
                countContent: d.count_content,
            };
        });

        return {
            indexThreads: indexThreads,
            nextID: res.data.next_id ? res.data.next_id : null,
        };
    },

    create: async (roomID: number, text: string, img?: string): Promise<ThreadUser> => {
        const apiClient = await asyncApiClient.create();
        const res = await apiClient.post(`threads`, {
            text: text,
            room_id: roomID,
            img: img,
        });
        const thread = res.data.thread;
        const user = res.data.user;

        return {
            thread: {
                id: thread.id,
                roomID: thread.room_id,
                userID: thread.user_id,
                text: thread.text,
                createdAt: new Date(thread.created_at),
                updatedAt: new Date(thread.updated_at),
                img: thread.img,
            },
            user: {
                id: user.id,
                name: user.name,
                createdAt: new Date(user.created_at),
                avatar: user.avatar,
            },
        };
    },
};
