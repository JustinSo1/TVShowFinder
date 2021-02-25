package userinterface

import (
	ui "github.com/gizak/termui/v3"
)

// NewTerminalWindow creates a new terminal window
func NewTerminalWindow(file []byte) *ui.Grid {
	// Graph, Link List, Info

	l := NewImageLinkList(file)

	g2 := NewProgressGraph()

	p2 := NewImageInfo("Image Info", "Blah blah blah\nText")

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1, g2),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, l),
			ui.NewCol(1.0/2, p2),
		),
	)
	return grid
}
