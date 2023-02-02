package webapi

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	goa "github.com/shogo82148/goa-v1"
)

type Comm chan struct{}

// WsConnections is connection pool websocket conn
type WsConnections struct {
	connections map[int][]Comm
	ctx         context.Context
}

// newConnections create WsConnections
func newConnections(ctx context.Context) *WsConnections {
	return &WsConnections{
		connections: map[int][]Comm{},
		ctx:         ctx,
	}
}

func (l *WsConnections) appendRoomConn(roomID model.RoomID, comm Comm) {
	list := l.connections[int(roomID)]
	if list == nil {
		list = []Comm{}
	}
	list = append(list, comm)
	l.connections[int(roomID)] = list
	goa.LogInfo(l.ctx, "appendRoomConn", "list", list)
}

func (l *WsConnections) removeRoomConn(roomID model.RoomID, comm Comm) {
	list := l.connections[int(roomID)]
	if list == nil {
		list = []Comm{}
	}
	newList := []Comm{}
	for _, c := range list {
		if c == comm {
			continue
		}
		newList = append(newList, c)
	}

	l.connections[int(roomID)] = newList
	goa.LogInfo(l.ctx, "removeRoomConn", "list", newList)
}

func (l *WsConnections) updateRoom(roomID model.RoomID) {
	goa.LogInfo(l.ctx, "updateRoom", "roomID", roomID)
	comms, ok := l.connections[int(roomID)]
	if !ok {
		return
	}
	goa.LogInfo(l.ctx, "updateRoom", "comms", comms)
	for _, comm := range comms {
		comm <- struct{}{}
	}
}
