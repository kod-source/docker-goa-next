import { ShowRoom } from "lib/model/room";
import { User } from "lib/model/user";

export const getRoomName = (showRoom: ShowRoom, my: User): string => {
    const roomName = showRoom.room.name;
    if (showRoom.room.isGroup) {
        return roomName;
    }
    const segmentsRoomName = roomName.split("/");
    const s = segmentsRoomName.find((s) => s !== my.name);
    if (!s) {
        return "";
    }
    return s;
};
