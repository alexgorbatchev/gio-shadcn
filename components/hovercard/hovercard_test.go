package hovercard_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/hovercard"
	"github.com/bnema/gio-shadcn/theme"
)

func TestHoverCardHoverPreviewCard(t *testing.T) {
	hc := hovercard.New(hovercard.Config{
		Title:       "Artist Profile",
		Description: "Aethelgard - Progressive House",
		Hovered:     true,
	})
	if !hc.Hovered {
		t.Fatalf("expected Hovered true")
	}
}

func TestHoverCardHoveredTriggerDetector(t *testing.T) {
	hc := hovercard.New(hovercard.Config{
		Hovered: false,
	})
	if hc.Hovered {
		t.Fatalf("expected Hovered false")
	}
}

func TestHoverCardPopoverProfileDetails(t *testing.T) {
	th := theme.NewDark()
	hc := hovercard.New(hovercard.Config{
		Title:       "Artist Profile",
		Description: "Aethelgard - Progressive House",
		Hovered:     true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(250, 80)),
	}
	dims := hc.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
