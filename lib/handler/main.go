package handler

import (
	"fmt"
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
	fmt.Println("starting server...")
	ServeMux.Handle("/", Router)
}
