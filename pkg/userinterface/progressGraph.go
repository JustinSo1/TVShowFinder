package userinterface

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// ProgressGraph is loading graph of searching of image
type ProgressGraph struct {
	*widgets.Gauge
}

// SetPercent updates the percentage
func (graph *ProgressGraph) SetPercent(percent int) {
	graph.Percent = percent
}

// SetTitle updates the title
func (graph *ProgressGraph) SetTitle(title string) {
	graph.Title = title
}

// SetBarColor updates the bar color
func (graph *ProgressGraph) SetBarColor(barColor ui.Color) {
	graph.BarColor = barColor
}

// SetLabelStyle updates the label style
func (graph *ProgressGraph) SetLabelStyle(labelStyle ui.Color) {
	graph.LabelStyle = ui.NewStyle(labelStyle)
}

// SetBorderStyleFg updates the border style
func (graph *ProgressGraph) SetBorderStyleFg(borderStyle ui.Color) {
	graph.BorderStyle.Fg = borderStyle
}

// NewProgressGraph returns progress graph of searching
func NewProgressGraph() *ProgressGraph {
	graph := &ProgressGraph{widgets.NewGauge()}
	graph.SetTitle("Progress Indicator")
	graph.SetPercent(0)
	graph.SetBarColor(ui.ColorYellow)
	graph.SetLabelStyle(ui.ColorBlue)
	graph.SetBorderStyleFg(ui.ColorWhite)
	return graph
}
