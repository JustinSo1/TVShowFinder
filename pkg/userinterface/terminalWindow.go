package userinterface

import (
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

// NewTerminalWindow creates a new terminal window
func NewTerminalWindow(file []byte) *TerminalWindow {
	window := &TerminalWindow{
		grid:  ui.NewGrid(),
		list:  NewImageLinkList(file),
		info:  NewImageInfo("Image Info", ""),
		graph: NewProgressGraph(),
	}

	for _, link := range convert.FileToList(file) {
		go func(link string) {
			if search.IsURL(link) {
				window.info.Text += search.ByLink(link)
			}
		}(link)
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
