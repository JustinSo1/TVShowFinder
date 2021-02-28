package userinterface

import (
	"fmt"
	"time"

	"github.com/JustinSo1/TVShowFinder/pkg/convert"
	"github.com/JustinSo1/TVShowFinder/pkg/search"
	ui "github.com/gizak/termui/v3"
)

// TerminalWindow is the UI of the window
type TerminalWindow struct {
	grid  *ui.Grid
	list  *ImageLinkList
	info  *ImageInfo
	graph *ProgressGraph
}

// UpdatePercent updates the window
func (term *TerminalWindow) UpdatePercent(percent int) {
	term.graph.SetPercent(percent % 101)
}

// Grid returns grid
func (term *TerminalWindow) Grid() *ui.Grid {
	return term.grid
}

// Info returns search results
func (term *TerminalWindow) Info() *ImageInfo {
	return term.info
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

// Display displays JSON response
func (term *TerminalWindow) Display(file []byte) {
	fileList := convert.FileToList(file)
	delta := 100 / len(fileList)
	for _, link := range fileList {
		if search.IsURL(link) {
			term.info.SetText(term.info.Text + search.ByLink(link))
		}
		term.graph.SetPercent(term.graph.Percent + int(delta))
		ui.Render(term.grid)
	}
	term.graph.SetPercent(100)
	ui.Render(term.grid)
}

// NewTerminalWindow creates a new terminal window
func NewTerminalWindow(file []byte) *TerminalWindow {
	// defer elapsed("page")() // deferred call's arguments are evaluated immediately but execution is delayed

	window := &TerminalWindow{
		grid:  ui.NewGrid(),
		list:  NewImageLinkList(file),
		info:  NewImageInfo("Image Info", ""),
		graph: NewProgressGraph(),
	}

	termWidth, termHeight := ui.TerminalDimensions()
	window.grid.SetRect(0, 0, termWidth, termHeight)
	window.grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1, window.graph),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, window.list),
			ui.NewCol(1.0/2, window.info),
		),
	)
	return window
}
