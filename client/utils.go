package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func createConn() (*websocket.Conn, *http.Response, error) {
	var u url.URL

	u = url.URL{Scheme: "ws", Host: "www.google.com", Path: "/"}

	dialer := websocket.Dialer{
		HandshakeTimeout: 2 * time.Minute,
	}
	conn, res, _ := dialer.Dial(u.String(), nil)

	return conn, res, nil

}
