package userinterface

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// ImageInfo is a paragraph
type ImageInfo struct {
	*widgets.Paragraph
}

// SetTitle sets title of paragraph
func (paragraph *ImageInfo) SetTitle(title string) {
	paragraph.Title = title
}

// SetText sets text
func (paragraph *ImageInfo) SetText(text string) {
	paragraph.Text = text
}

// SetBorderStyleFg sets colour of border
func (paragraph *ImageInfo) SetBorderStyleFg(color ui.Color) {
	paragraph.BorderStyle.Fg = color
}

// NewImageInfo returns info
func NewImageInfo(title, text string) *ImageInfo {
	info := &ImageInfo{widgets.NewParagraph()}
	info.SetTitle(title)
	info.SetText(text)
	return info
}

// // UpdateText updates text of paragraph
// func UpdateText(p2 *widgets.Paragraph, lines []string) {
// 	for _, link := range lines {
// 		if search.IsURL(link) {
// 			p2.Text += search.ByLink(link)
// 		}
// 	}
// }
