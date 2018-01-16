package net

import (
	"errors"
	"time"
)

// testable struct for things that involve writing to
// web sockets, etc
// implements ocenet.WebsocketEy interface
type testWebSocketConn struct {
	MsgData [][]byte
	Done    bool
}

func (ws *testWebSocketConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func (ws *testWebSocketConn) WriteMessage(messageType int, data []byte) error {
	if !ws.Done {
		ws.MsgData = append(ws.MsgData, data)
		return nil
	} else {
		return errors.New("can't write to done msg")
	}
}

func (ws *testWebSocketConn) Close() error {
	ws.Done = true
	return nil
}

func (ws *testWebSocketConn) GetData() [][]byte {
	return ws.MsgData
}

func (ws *testWebSocketConn) SendIt(data []byte) error {
	err := ws.WriteMessage(0, data)
	return err
}

func (ws *testWebSocketConn) SendError(errorDesc []byte) {

}

func (ws *testWebSocketConn) Finish(chan int) {

}

func NewWebSocketConn() *testWebSocketConn {
	var data [][]byte
	return &testWebSocketConn{
		MsgData: data,
		Done:    false,
	}
}