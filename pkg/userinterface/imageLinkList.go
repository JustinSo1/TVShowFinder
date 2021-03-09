package userinterface

import (
	"github.com/JustinSo1/TVShowFinder/internal"
	"github.com/JustinSo1/TVShowFinder/pkg/convert"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// ImageLinkList is a list of image links
type ImageLinkList struct {
	*widgets.List
}

// SetTitle sets title of paragraph
func (list *ImageLinkList) SetTitle(title string) {
	list.Title = title
}

// SetRows sets numbered bullets of list
func (list *ImageLinkList) SetRows(rows []string) {
	list.Rows = rows
}

// SetTextStyle sets colour of text
func (list *ImageLinkList) SetTextStyle(style ui.Style) {
	list.TextStyle = style
}

// NewImageLinkList returns list of links
func NewImageLinkList(file []byte) *ImageLinkList {
	list := &ImageLinkList{widgets.NewList()}
	list.SetTitle("List")
	numberedFile, err := convert.FileToNumberedList(file)
	internal.HandleError(err)
	list.SetRows(numberedFile)
	list.SetTextStyle(ui.NewStyle(ui.ColorBlue))
	return list
}
