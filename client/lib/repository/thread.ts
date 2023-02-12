import { AllThread, IndexThread, Thread } from "lib/model/thread";
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
};
