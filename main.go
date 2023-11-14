package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/sttronn/go-editor/editor"
)

type UIConfig struct {
	EditorConfig editor.EditorConfig
}

func initUI(conf UIConfig) error {

	err := termbox.Init()

	if err != nil {
		return err
	}

	e := editor.NewEditor(conf.EditorConfig)

	e.SetSize(termbox.Size())
	e.SetText("hello I am your editor")
	e.SendDraw()
	e.IsConnected = true

	e.Draw()

	return nil

}

func main() {
	conf := UIConfig{
		EditorConfig: editor.EditorConfig{
			ScrollEnabled: true,
		},
	}

	if err := initUI(conf); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize UI: %v\n", err)
		os.Exit(1)
	}

	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					return // Exit on ESC
				}
				// Handle other keys or custom functionality here
			}

			// Redraw the editor
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			// drawEditor() // Implement this function based on your editor's drawing logic
			termbox.Flush()
		}
	}
}
