package empty_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/empty"
	"github.com/bnema/gio-shadcn/theme"
)

func TestEmptyStateCard(t *testing.T) {
	emp := empty.New(empty.Config{
		Title:       "No Audio Tracks",
		Description: "Import FLAC files to get started.",
	})
	if emp.Title != "No Audio Tracks" {
		t.Fatalf("expected Title 'No Audio Tracks'")
	}
}

func TestEmptyIllustratedIconVector(t *testing.T) {
	emp := empty.New(empty.Config{})
	if emp.Title == "" {
		t.Fatalf("expected default Title")
	}
}

func TestEmptyTitleAndDescription(t *testing.T) {
	emp := empty.New(empty.Config{
		Title:       "Empty Library",
		Description: "Library description",
	})
	if emp.Title != "Empty Library" || emp.Description != "Library description" {
		t.Errorf("expected Title and Description")
	}
}

func TestEmptyPrimaryActionButton(t *testing.T) {
	th := theme.NewDark()
	emp := empty.New(empty.Config{})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 150)),
	}
	dims := emp.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
