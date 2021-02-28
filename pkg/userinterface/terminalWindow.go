package userinterface

import (
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

// Grid updates the window
func (term *TerminalWindow) Grid() *ui.Grid {
	return term.grid
}

// NewTerminalWindow creates a new terminal window
func NewTerminalWindow(file []byte) *TerminalWindow {
	window := &TerminalWindow{
		grid:  ui.NewGrid(),
		list:  NewImageLinkList(file),
		info:  NewImageInfo("Image Info", "bLAH BLAH BLAH\nText"),
		graph: NewProgressGraph(),
	}

	// grid := ui.NewGrid()
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
