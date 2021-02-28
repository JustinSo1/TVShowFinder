package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JustinSo1/TVShowFinder/internal"
	"github.com/JustinSo1/TVShowFinder/pkg/userinterface"
	ui "github.com/gizak/termui/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	internal.HandleError(err)

	err = ui.Init()
	internal.HandleError(err)
	defer ui.Close()

	window := userinterface.NewTerminalWindow(data)

	ui.Render(window.Grid())

	window.Display(data)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				window.Grid().SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(window.Grid())
			}
		}
	}
}
