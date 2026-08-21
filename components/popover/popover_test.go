package popover_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/popover"
	"github.com/bnema/gio-shadcn/theme"
)

func TestPopoverCreation(t *testing.T) {
	pop := popover.New(popover.Config{
		Title:       "Audio Dimensions",
		Description: "96kHz 24-bit PCM stream",
		Open:        true,
	})

	if !pop.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestPopoverLayout(t *testing.T) {
	th := theme.NewDark()
	pop := popover.New(popover.Config{
		Title:       "Audio Dimensions",
		Description: "96kHz 24-bit PCM stream",
		Open:        true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 100)),
	}
	dims := pop.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Popover.Layout")
	}
}
