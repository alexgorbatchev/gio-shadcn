package hovercard

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	HoverCardItem *HoverCard
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		HoverCardItem: New(Config{Title: "Artist Profile", Description: "Aethelgard - Progressive House", Hovered: false}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return s.HoverCardItem.Layout(gtx, th)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
