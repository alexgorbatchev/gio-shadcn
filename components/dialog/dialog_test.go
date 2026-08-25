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
	d := dialog.New(dialog.Config{
		Title:       "Edit profile",
		Description: "Make changes to your profile here.",
		Open:        true,
	})
	if !d.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestAlertDialogConfirmation(t *testing.T) {
	d := dialog.New(dialog.Config{
		Title:       "Are you absolutely sure?",
		Description: "This action cannot be undone.",
		ConfirmText: "Continue",
		CancelText:  "Cancel",
		Open:        true,
	})
	if d.ConfirmText != "Continue" {
		t.Errorf("expected ConfirmText 'Continue', got %s", d.ConfirmText)
	}
}

func TestAlertDialogDestructive(t *testing.T) {
	d := dialog.New(dialog.Config{
		Title:       "Delete chat?",
		Description: "Permanently delete this chat conversation.",
		ConfirmText: "Delete",
		Open:        true,
	})
	if d.Title != "Delete chat?" {
		t.Errorf("unexpected title: %s", d.Title)
	}
}

func TestDialogLayout(t *testing.T) {
	th := theme.NewDark()
	d := dialog.New(dialog.Config{
		Title:       "Test Dialog",
		Description: "Dialog body description",
		Open:        true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := d.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Dialog.Layout")
	}
}
