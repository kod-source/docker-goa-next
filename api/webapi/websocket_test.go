package webapi

import (
	"context"
	"testing"

	"github.com/kod-source/docker-goa-next/app/model"
)

func TestAppendRoomConn(t *testing.T) {
	ctx := context.Background()
	l := &WsConnections{
		connections: make(map[int][]Comm),
		ctx:         ctx,
	}
	roomID := 1
	comm := make(Comm)
	l.appendRoomConn(model.RoomID(roomID), comm)

	// test if the connection is added to the list
	if len(l.connections[roomID]) != 1 {
		t.Errorf("Expected 1 connection, but got %v", len(l.connections[roomID]))
	}
	if l.connections[roomID][0] != comm {
		t.Errorf("Expected %v, but got %v", comm, l.connections[roomID][0])
	}
}

func TestRemoveRoomConn(t *testing.T) {
	ctx := context.Background()
	conn := WsConnections{
		connections: make(map[int][]Comm),
		ctx:         ctx,
	}

	roomID := 1
	comm1 := make(Comm)
	comm2 := make(Comm)
	comm3 := make(Comm)

	conn.appendRoomConn(model.RoomID(roomID), comm1)
	conn.appendRoomConn(model.RoomID(roomID), comm2)
	conn.appendRoomConn(model.RoomID(roomID), comm3)

	if len(conn.connections[roomID]) != 3 {
		t.Errorf("expected 3 connections, but got %d", len(conn.connections[roomID]))
	}

	conn.removeRoomConn(model.RoomID(roomID), comm2)

	if len(conn.connections[roomID]) != 2 {
		t.Errorf("expected 2 connections, but got %d", len(conn.connections[roomID]))
	}
}
