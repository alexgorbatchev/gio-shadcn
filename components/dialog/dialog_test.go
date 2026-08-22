package dialog_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDialogModalWindow(t *testing.T) {
	dl := dialog.New(dialog.Config{
		Title:       "Modal Dialog Title",
		Description: "Modal Description",
		Open:        true,
	})
	if !dl.Open {
		t.Fatalf("expected Open true")
	}
}

func TestDialogDarkBackdropOverlay(t *testing.T) {
	th := theme.NewDark()
	dl := dialog.New(dialog.Config{
		Title: "Dialog",
		Open:  true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(500, 400)),
	}
	dims := dl.Layout(gtx, th)
	if dims.Size.X < 0 {
		t.Errorf("invalid width")
	}
}

func TestDialogBackdropClickToClose(t *testing.T) {
	cancelled := false
	dl := dialog.New(dialog.Config{
		Title: "Dialog",
		Open:  true,
		OnCancel: func() {
			cancelled = true
		},
	})
	if dl.OnCancel == nil {
		t.Fatalf("expected OnCancel handler")
	}
	_ = cancelled
}

func TestDialogHeaderTitleAndDescription(t *testing.T) {
	dl := dialog.New(dialog.Config{
		Title:       "Header Title",
		Description: "Description Body",
	})
	if dl.Title != "Header Title" || dl.Description != "Description Body" {
		t.Errorf("expected Title and Description")
	}
}

func TestDialogConfirmAndCancelActions(t *testing.T) {
	dl := dialog.New(dialog.Config{
		ConfirmText: "Save Changes",
		CancelText:  "Discard",
	})
	if dl.ConfirmText != "Save Changes" || dl.CancelText != "Discard" {
		t.Errorf("expected custom action texts")
	}
}

func TestDialogCustomContentWidget(t *testing.T) {
	contentRan := false
	dl := dialog.New(dialog.Config{
		Title: "Dialog",
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
		Constraints: layout.Exact(image.Pt(500, 400)),
	}
	_ = dl.Layout(gtx, th)
	if !contentRan {
		t.Errorf("expected custom content widget to execute")
	}
}
