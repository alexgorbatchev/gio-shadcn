package hovercard_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/hovercard"
	"github.com/bnema/gio-shadcn/theme"
)

func TestHoverCardCreation(t *testing.T) {
	hc := hovercard.New(hovercard.Config{
		Title:       "Artist Profile",
		Description: "Aethelgard - Progressive House",
		Hovered:     true,
	})

	if !hc.Hovered {
		t.Errorf("expected Hovered to be true")
	}
}

func TestHoverCardLayout(t *testing.T) {
	th := theme.NewDark()
	hc := hovercard.New(hovercard.Config{
		Title:       "Artist Profile",
		Description: "Aethelgard - Progressive House",
		Hovered:     true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(250, 80)),
	}
	dims := hc.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from HoverCard.Layout")
	}
}
