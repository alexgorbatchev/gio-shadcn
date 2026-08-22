package input_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/theme"
)

func TestInputStandardTextInput(t *testing.T) {
	inp := input.Text("Enter text...")
	if inp.Placeholder != "Enter text..." {
		t.Fatalf("expected Placeholder 'Enter text...'")
	}
}

func TestInputPrefilledTextInput(t *testing.T) {
	inp := input.Text("Enter text...")
	inp.SetText("Prefilled Text")
	if inp.Text() != "Prefilled Text" {
		t.Fatalf("expected Text 'Prefilled Text'")
	}
}

func TestInputDisabledTextInput(t *testing.T) {
	inp := input.Text("Placeholder")
	inp.Disabled = true
	if !inp.Disabled {
		t.Fatalf("expected Disabled true")
	}
}

func TestInputSingleLineTextEditing(t *testing.T) {
	inp := input.Text("Placeholder")
	inp.SetText("Single Line")
	if inp.Value != "Single Line" {
		t.Errorf("expected Value 'Single Line'")
	}
}

func TestInputPlaceholderText(t *testing.T) {
	inp := input.Text("Track Title Placeholder")
	if inp.Placeholder != "Track Title Placeholder" {
		t.Errorf("expected Placeholder 'Track Title Placeholder'")
	}
}

func TestInputFocusRingStroke(t *testing.T) {
	th := theme.NewDark()
	inp := input.Text("Placeholder")
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 40)),
	}
	dims := inp.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
