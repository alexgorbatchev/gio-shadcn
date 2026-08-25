package popover_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/popover"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPopoverStandard(t *testing.T) {
	th := theme.NewDark()
	p := popover.New(popover.Config{
		Title:       "Dimensions",
		Description: "Width: 100%",
		Open:        true,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 150))}
	dims := p.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestPopoverClosed(t *testing.T) {
	th := theme.NewDark()
	p := popover.New(popover.Config{
		Title:       "Closed",
		Description: "Should not render",
		Open:        false,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 150))}
	dims := p.Layout(gtx, th)
	if dims.Size.X != 0 || dims.Size.Y != 0 {
		t.Errorf("expected 0 dimensions for closed popover")
	}
}
