package main

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func main() {
	http.Handle("/echo", websocket.Handler(EchoHandler))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func EchoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
