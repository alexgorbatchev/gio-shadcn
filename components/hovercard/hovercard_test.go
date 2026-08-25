package hovercard_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/hovercard"
	"github.com/bnema/gio-shadcn/theme"
)

func TestHoverCardBasic(t *testing.T) {
	th := theme.NewDark()
	hc := hovercard.New(hovercard.Config{
		Title:       "@nextjs",
		Description: "The React Framework",
		Hovered:     true,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	dims := hc.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestHoverCardHovered(t *testing.T) {
	th := theme.NewDark()
	hc := hovercard.New(hovercard.Config{
		Title:       "@nextjs",
		Description: "The React Framework",
		Hovered:     true,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 150))}
	dims := hc.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
