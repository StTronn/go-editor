package main

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nsf/termbox-go"
)

func initUI(conn *websocket.Conn, conf UIConfig) error {
	var wg sync.WaitGroup

	err := termbox.Init()

	if err != nil {
		return err
	}

	defer termbox.Close()
	e.SetSize(termbox.Size())
	e.SetText("DUMMY Test")
	e.SendDraw()
	e.IsConnected = true

	go drawLoop(&wg)

	err = mainLoop(conn)

	if err != nil {
		return err
	}

	return nil

}

// mainLoop is the main update loop for the UI.
func mainLoop(conn *websocket.Conn) error {
	// termboxChan is used for sending and receiving termbox events.
	termboxChan := getTermboxChan()

	// msgChan := getMsgChan(conn)

	for {
		select {
		case termboxEvent := <-termboxChan:
			err := handleTermboxEvent(termboxEvent, conn)
			if err != nil {
				return err
			}
			// case msg := <-msgChan:
			// 	handleMsg(msg, conn)
		}
	}
}

func drawLoop(wg *sync.WaitGroup) {
	for {
		<-e.DrawChan
		e.Draw()
	}
}
