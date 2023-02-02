package webapi

import (
	"context"

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

func (l *WsConnections) apendConn(roomID int, comm Comm) {
	list := l.connections[roomID]
	if list == nil {
		list = []Comm{}
	}
	list = append(list, comm)
	l.connections[roomID] = list
	goa.LogInfo(l.ctx, "apendConn", "list", list)
}

func (l *WsConnections) removeConn(roomID int, comm Comm) {
	list := l.connections[roomID]
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

	l.connections[roomID] = newList
	goa.LogInfo(l.ctx, "removeConn", "list", newList)
}

func (l *WsConnections) updateRoom(roomID int) {
	goa.LogInfo(l.ctx, "updateRoom", "roomID", roomID)
	comms, ok := l.connections[roomID]
	if !ok {
		return
	}
	goa.LogInfo(l.ctx, "updateRoom", "comms", comms)
	for _, comm := range comms {
		comm <- struct{}{}
	}
}
