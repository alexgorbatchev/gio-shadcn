package sheet_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSheetCreation(t *testing.T) {
	sh := sheet.New(sheet.Config{
		Title:       "Track Inspector",
		Description: "Detailed audio parameters and FLAC metadata",
		Open:        true,
	})

	if !sh.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestSheetLayout(t *testing.T) {
	th := theme.NewDark()
	sh := sheet.New(sheet.Config{
		Title:       "Track Inspector",
		Description: "Detailed audio parameters and FLAC metadata",
		Open:        true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := sh.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Sheet.Layout")
	}
}
