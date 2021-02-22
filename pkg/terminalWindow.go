package pkg

import (
	ui "github.com/gizak/termui/v3"
)

// NewTerminalWindow creates a new terminal window
func NewTerminalWindow(items ...ui.Drawable) *ui.Grid {
	// Graph, Link List, Info
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1, items[2]),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, items[0]),
			ui.NewCol(1.0/2, items[1]),
		),
	)
	return grid
}
