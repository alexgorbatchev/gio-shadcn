package drawer_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDrawerCreation(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title:       "Audio Mixer Settings",
		Description: "Configure master out gain and ASIO buffer size",
		Open:        true,
	})

	if !dr.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestDrawerLayout(t *testing.T) {
	th := theme.NewDark()
	dr := drawer.New(drawer.Config{
		Title:       "Audio Mixer Settings",
		Description: "Configure master out gain and ASIO buffer size",
		Open:        true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := dr.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Drawer.Layout")
	}
}
