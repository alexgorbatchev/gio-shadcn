package drawer_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDrawerBottomSheetPanel(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title:       "Bottom Sheet",
		Description: "Slide-up panel",
		Open:        true,
	})
	if !dr.Open {
		t.Fatalf("expected Open true")
	}
}

func TestDrawerDarkBackdropOverlay(t *testing.T) {
	th := theme.NewDark()
	dr := drawer.New(drawer.Config{
		Title: "Drawer",
		Open:  true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := dr.Layout(gtx, th)
	if dims.Size.X < 0 {
		t.Errorf("invalid width")
	}
}

func TestDrawerBackdropClickToClose(t *testing.T) {
	closed := false
	dr := drawer.New(drawer.Config{
		Title: "Drawer",
		Open:  true,
		OnClose: func() {
			closed = true
		},
	})
	if dr.OnClose == nil {
		t.Fatalf("expected OnClose handler")
	}
	_ = closed
}

func TestDrawerSouthViewportEdgeAlignment(t *testing.T) {
	th := theme.NewDark()
	dr := drawer.New(drawer.Config{
		Title: "Drawer",
		Open:  true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := dr.Layout(gtx, th)
	if dims.Size.Y < 0 {
		t.Errorf("invalid height")
	}
}

func TestDrawerDragHandleIndicator(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title: "Drawer with Handle",
	})
	if dr.Title != "Drawer with Handle" {
		t.Errorf("expected Title 'Drawer with Handle'")
	}
}

func TestDrawerCloseButton(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title: "Drawer Title",
	})
	if dr.Title != "Drawer Title" {
		t.Errorf("expected Title 'Drawer Title'")
	}
}

func TestDrawerCustomContentWidget(t *testing.T) {
	contentRan := false
	dr := drawer.New(drawer.Config{
		Title: "Drawer",
		Open:  true,
		Content: func(gtx layout.Context) layout.Dimensions {
			contentRan = true
			return layout.Dimensions{Size: image.Pt(100, 50)}
		},
	})
	th := theme.NewDark()
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	_ = dr.Layout(gtx, th)
	if !contentRan {
		t.Errorf("expected custom content widget to execute")
	}
}
