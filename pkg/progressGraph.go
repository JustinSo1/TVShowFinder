package pkg

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// NewProgressGraph returns progress graph of searching
func NewProgressGraph() *widgets.Gauge {
	g2 := widgets.NewGauge()
	g2.Title = "Slim Gauge"
	g2.Percent = 0
	g2.BarColor = ui.ColorYellow
	g2.LabelStyle = ui.NewStyle(ui.ColorBlue)
	g2.BorderStyle.Fg = ui.ColorWhite
	return g2
}
