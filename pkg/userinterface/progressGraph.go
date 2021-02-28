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

// NewProgressGraph returns progress graph of searching
func NewProgressGraph() *ProgressGraph {
	graph := &ProgressGraph{widgets.NewGauge()}
	graph.Title = "Progress Indicator"
	graph.Percent = 0
	graph.BarColor = ui.ColorYellow
	graph.LabelStyle = ui.NewStyle(ui.ColorBlue)
	graph.BorderStyle.Fg = ui.ColorWhite
	return graph
}
