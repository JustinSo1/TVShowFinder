package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/JustinSo1/TVShowFinder/internal"
	"github.com/JustinSo1/TVShowFinder/pkg"
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

	l := pkg.NewImageLinkList(data)

	g2 := pkg.NewProgressGraph()

	p2 := pkg.NewImageInfo()

	draw := func(count int) {
		g2.Percent = count % 101
	}

	grid := pkg.NewTerminalWindow(l, p2, g2)
	ui.Render(grid)

	lines := pkg.ReformatFile(data)
	pkg.UpdateText(p2, lines)

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
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if tickerCount == 100 {
				return
			}
			draw(tickerCount)
			ui.Render(grid)
			tickerCount++
		}
	}
}
