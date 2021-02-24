package userinterface

import (
	"github.com/JustinSo1/TVShowFinder/pkg/search"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// NewImageInfo returns info
func NewImageInfo() *widgets.Paragraph {
	p2 := widgets.NewParagraph()
	p2.Title = "Multiline"
	p2.Text = ""
	p2.BorderStyle.Fg = ui.ColorYellow
	return p2
}

// UpdateText updates text of paragraph
func UpdateText(p2 *widgets.Paragraph, lines []string) {
	for _, link := range lines {
		if search.IsURL(link) {
			p2.Text += search.ByLink(link)
		}
	}
}
