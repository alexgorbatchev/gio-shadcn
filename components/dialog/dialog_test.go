package dialog_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDialogCreation(t *testing.T) {
	dl := dialog.New(dialog.Config{
		Title:       "Confirm Audio Engine Reset",
		Description: "Are you sure you want to reset all mixer channels?",
		Open:        true,
	})

	if !dl.Open {
		t.Errorf("expected Open to be true")
	}

	if dl.Title != "Confirm Audio Engine Reset" {
		t.Errorf("expected Title to be 'Confirm Audio Engine Reset', got %s", dl.Title)
	}
}

func TestDialogLayout(t *testing.T) {
	th := theme.NewDark()
	dl := dialog.New(dialog.Config{
		Title:       "Confirm Audio Engine Reset",
		Description: "Are you sure you want to reset all mixer channels?",
		Open:        true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(500, 400)),
	}
	dims := dl.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Dialog.Layout")
	}
}
