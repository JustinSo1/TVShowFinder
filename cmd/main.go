package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	window := userinterface.NewTerminalWindow(data)

	ui.Render(window.Grid())

	tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
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
		case <-ticker:
			if tickerCount == 100 {
				return
			}
			window.UpdatePercent(tickerCount)
			ui.Render(window.Grid())
			tickerCount++
		}
	}
}
