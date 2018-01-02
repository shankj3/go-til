package net

import (
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

// for initializing connections / configurations once, then passing it around for lifetime of application
// Ctx can be anything, just have to set H to be HandleFunc that also takes in the context as first value.
// in handle func, cast Ctx interface{} to your struct that you initialized in startup so you can access fields
// for ex:
// ```
// appctx := &MyContext{config: "config yay wooo the best"}
// muxi.Handle("/ws/builds/{hash}", &ocenet.AppContextHandler{appctx, stream}).Methods("GET")
// ...
//
// func stream(ctx interface{}, w http.ResponseWriter, r *http.Request){
//     a := ctx.(*MyContext)
//     // do stuff
// ...
// ```
type AppContextHandler struct {
	Ctx interface{}
	H func(interface{}, http.ResponseWriter, *http.Request)
}


func (ah *AppContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.H(ah.Ctx, w, r)
}

// testable interface for using websockets
type WebsocketEy interface {
	SetWriteDeadline(t time.Time) error
	WriteMessage(messageType int, data []byte) error
	Close() error
}



// Upgrade returns an OcenetWs object with a websocket connection upgraded by the ResponseWriter and the request
func Upgrade(upgrader websocket.Upgrader, w http.ResponseWriter, r *http.Request) (*OcenetWs, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &OcenetWs{cxn: conn}, nil
}

// fulfills the Streamable interface in ocelot
type OcenetWs struct {
	cxn *websocket.Conn
}

func (ws *OcenetWs) SendIt(data []byte) error {
	if err := ws.cxn.WriteMessage(websocket.TextMessage, data); err != nil {
		ws.cxn.Close()
		return err
	}
	return nil
}

func (ws *OcenetWs) SendError(errorDesc []byte) {
	ws.cxn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	ws.cxn.WriteMessage(websocket.TextMessage, []byte("ERROR!\n"))
	ws.cxn.WriteMessage(websocket.TextMessage, errorDesc)
	ws.cxn.Close()
}

func (ws *OcenetWs) Finish(done chan int) {
	close(done)
	ws.cxn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	ws.cxn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	ws.cxn.Close()
}