package scrollarea_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/scrollarea"
	"github.com/bnema/gio-shadcn/theme"
)

func TestScrollAreaCreation(t *testing.T) {
	sa := scrollarea.New(scrollarea.Config{
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 200)}
		},
	})

	if sa.List == nil {
		t.Errorf("expected List to be initialized")
	}
}

func TestScrollAreaLayout(t *testing.T) {
	th := theme.NewDark()
	sa := scrollarea.New(scrollarea.Config{
		Widget: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 200)}
		},
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(100, 200)),
	}
	dims := sa.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from ScrollArea.Layout")
	}
}
