package pkg

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// NewImageLinkList returns list of links
func NewImageLinkList(data []byte) *widgets.List {
	l := widgets.NewList()
	l.Title = "List"
	l.Rows = ReformatFile(data)
	l.TextStyle = ui.NewStyle(ui.ColorBlue)
	l.WrapText = true
	return l
}
