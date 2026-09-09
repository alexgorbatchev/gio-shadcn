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
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := sh.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Sheet.Layout")
	}
}

func TestSheetSides(t *testing.T) {
	th := theme.NewDark()

	leftSheet := sheet.New(sheet.Config{
		Title: "Left Sheet",
		Side:  sheet.SideLeft,
		Open:  true,
	})
	rightSheet := sheet.New(sheet.Config{
		Title: "Right Sheet",
		Side:  sheet.SideRight,
		Open:  true,
	})

	gtx := layout.Context{
		Ops:         new(op.Ops),
		Constraints: layout.Exact(image.Pt(800, 600)),
	}

	dimsLeft := leftSheet.Layout(gtx, th)
	if dimsLeft.Size.X != 800 || dimsLeft.Size.Y != 600 {
		t.Errorf("expected left sheet to fill viewport 800x600, got %v", dimsLeft.Size)
	}

	dimsRight := rightSheet.Layout(gtx, th)
	if dimsRight.Size.X != 800 || dimsRight.Size.Y != 600 {
		t.Errorf("expected right sheet to fill viewport 800x600, got %v", dimsRight.Size)
	}
}
