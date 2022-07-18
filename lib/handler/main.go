package handler

import (
	"net/http"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
	ServeMux = http.NewServeMux()
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
)


func init() {
	ServeMux.Handle("/", Router)
}