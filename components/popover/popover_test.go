package popover_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/popover"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPopoverOpenCard(t *testing.T) {
	pop := popover.New(popover.Config{
		Title:       "Audio Dimensions",
		Description: "96kHz 24-bit PCM stream",
		Open:        true,
	})
	if !pop.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestPopoverClosedState(t *testing.T) {
	th := theme.NewDark()
	pop := popover.New(popover.Config{
		Title:       "Hidden Popover",
		Description: "Hidden Content",
		Open:        false,
	})
	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := pop.Layout(gtx, th)
	if dims.Size.X != 0 || dims.Size.Y != 0 {
		t.Errorf("expected 0 dimensions for closed popover")
	}
}

func TestPopoverTitleAndDescription(t *testing.T) {
	th := theme.NewDark()
	pop := popover.New(popover.Config{
		Title:       "Mixer Filter",
		Description: "Low pass cut-off at 200Hz",
		Open:        true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := pop.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
