package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/nsf/termbox-go"
	"github.com/sirupsen/logrus"
)

const (
	OperationInsert = iota
	OperationDelete
)

func handleTermboxEvent(ev termbox.Event, conn *websocket.Conn) error {
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyEsc, termbox.KeyCtrlC:
			return errors.New("pairPaid exiting")

		case termbox.KeyCtrlS:
			if fileName == "" {
				fileName = "pairpad-content.txt"
			}

			err := os.WriteFile(fileName, []byte(string(e.Text)), 0644)
			if err != nil {
				logrus.Errorf("failed to save %s", fileName)
				e.StatusChan <- fmt.Sprintf("Failed to save to %s", fileName)
				return err
			}
			e.StatusChan <- fmt.Sprintf("Saved coument to %s", fileName)

		case termbox.KeyCtrlL:
			if fileName != "" {
				logger.Log(logrus.InfoLevel, "LOADING DOCUMENT")
				content, err := os.ReadFile(fileName)
				if err != nil {
					logrus.Errorf("failed to load file %s", fileName)
					e.StatusChan <- fmt.Sprint("Failed to load %s", fileName)
				}
				e.SetText(string(content))
				e.SetX(0)
				logger.Log(logrus.InfoLevel, "SENDING DOCUMENT")
				//send to comms in future
			}

		case termbox.KeyArrowLeft, termbox.KeyCtrlB:
			e.MoveCursor(-1, 0)

		case termbox.KeyArrowRight, termbox.KeyCtrlF:
			e.MoveCursor(1, 0)

		case termbox.KeyArrowUp, termbox.KeyCtrlP:
			e.MoveCursor(0, -1)

		case termbox.KeyArrowDown, termbox.KeyCtrlN:
			e.MoveCursor(0, 1)

		case termbox.KeyBackspace, termbox.KeyBackspace2:
			e.PerformOperation(OperationDelete, ev.Ch)
		case termbox.KeyDelete:
			e.PerformOperation(OperationDelete, ev.Ch)

		case termbox.KeyTab:
			for i := 0; i < 4; i++ {
				ev.Ch = ' '
				e.PerformOperation(OperationInsert, ev.Ch)
			}

		case termbox.KeyEnter:
			ev.Ch = '\n'
			e.PerformOperation(OperationInsert, ev.Ch)

		case termbox.KeySpace:
			ev.Ch = ' '
			e.PerformOperation(OperationInsert, ev.Ch)

		default:
			if ev.Ch != 0 {
				e.PerformOperation(OperationInsert, ev.Ch)
			}
		}
	}
	e.SendDraw()
	return nil
}

func getTermboxChan() chan termbox.Event {
	termboxChan := make(chan termbox.Event)

	go func() {
		for {
			termboxChan <- termbox.PollEvent()
		}
	}()

	return termboxChan

}

func Content(doc string) string {
	return doc
}
