package toast_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/toast"
	"github.com/bnema/gio-shadcn/theme"
)

func TestToastSuccess(t *testing.T) {
	th := theme.NewDark()
	tst := toast.New(toast.Config{
		Title:       "Event created",
		Description: "Sunday, December 03",
		Visible:     true,
		Icon:        lucide.Check,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 60))}
	dims := tst.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestToastDestructive(t *testing.T) {
	th := theme.NewDark()
	tst := toast.New(toast.Config{
		Title:       "Error",
		Description: "Request failed",
		Variant:     theme.VariantDestructive,
		Visible:     true,
		Icon:        lucide.CircleAlert,
	})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 60))}
	dims := tst.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
