package toast_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/toast"
	"github.com/bnema/gio-shadcn/theme"
)

func TestToastCreation(t *testing.T) {
	tst := toast.New(toast.Config{
		Title:       "Track Loaded",
		Description: "Starlight Symphony loaded to Deck A.",
		Visible:     true,
	})

	if !tst.Visible {
		t.Errorf("expected Visible to be true")
	}
}

func TestToastLayout(t *testing.T) {
	th := theme.NewDark()
	tst := toast.New(toast.Config{
		Title:       "Track Loaded",
		Description: "Starlight Symphony loaded to Deck A.",
		Visible:     true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(300, 80)),
	}
	dims := tst.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Toast.Layout")
	}
}
