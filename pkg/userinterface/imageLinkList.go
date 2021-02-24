package userinterface

import (
	"github.com/JustinSo1/TVShowFinder/pkg/convert"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// NewImageLinkList returns list of links
func NewImageLinkList(file []byte) *widgets.List {
	l := widgets.NewList()
	l.Title = "List"
	l.Rows = convert.FileToNumberedList(file)
	l.TextStyle = ui.NewStyle(ui.ColorBlue)
	l.WrapText = true
	return l
}
