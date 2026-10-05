package theme

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/bnema/gio-shadcn/internal/testui"
)

func TestVariantDispatch(t *testing.T) {
	for _, th := range []*Theme{New(), NewDark()} {
		for _, v := range []Variant{VariantDefault, VariantDestructive, VariantOutline, VariantSecondary, VariantGhost, VariantLink, "unknown"} {
			button, input, bar, card := GetButtonVariant(v, &th.Colors), GetInputVariant(v, &th.Colors), GetTitleBarVariant(v, &th.Colors), GetCardVariant(v, &th.Colors)
			if button.Foreground.A == 0 || input.Foreground.A == 0 || bar.Foreground.A == 0 || card.Background != th.Colors.Card {
				t.Fatalf("variant %q lacks usable colors", v)
			}
			if v == VariantOutline && (input.BorderWidth != 1 || bar.BorderWidth != 1) {
				t.Fatal("outline lost border")
			}
			if v == VariantDestructive && input.FocusRing != th.Colors.Destructive {
				t.Fatal("destructive focus dispatch failed")
			}
			if v == VariantGhost && (input.Background.A != 0 || bar.Background.A == 0) {
				t.Fatal("ghost opacity dispatch failed")
			}
			if v == VariantSecondary && input.Foreground != th.Colors.SecondaryFg {
				t.Fatal("secondary foreground dispatch failed")
			}
		}
	}
}

func TestValidationRejectsIncompleteModes(t *testing.T) {
	if ValidateTheme(nil) == nil || validateColorScheme(nil) == nil {
		t.Fatal("nil theme/color scheme accepted")
	}
	th := New()
	if err := ValidateTheme(th); err != nil {
		t.Fatal(err)
	}
	th.Colors.Background.A = 0
	if ValidateTheme(th) == nil {
		t.Fatal("invalid light palette accepted")
	}
	th = New()
	th.DarkColors.Primary.A = 0
	if ValidateTheme(th) == nil {
		t.Fatal("invalid dark palette accepted")
	}
	th = New()
	th.MaterialTheme = nil
	th.ToggleDark()
	if !th.IsDark {
		t.Fatal("palette toggle depends on material theme")
	}
}

func TestBackdropLifecycleAndDimmerInteraction(t *testing.T) {
	th := New()
	defer th.ReleaseBackdrop()
	captures := 0
	layer := func(ops *op.Ops) { captures++; paint.Fill(ops, th.Colors.Primary) }
	th.UpdateBackdropIfNeeded(layer, image.Point{}, 4)
	if captures != 0 {
		t.Fatal("empty capture rendered")
	}
	th.UpdateBackdropIfNeeded(layer, image.Pt(160, 120), 4)
	if _, ok := th.BackdropOp(); !ok {
		t.Fatal("GPU backdrop capture unavailable")
	}
	first := captures
	th.UpdateBackdropIfNeeded(layer, image.Pt(160, 120), 4)
	if captures != first {
		t.Fatal("valid backdrop recaptured")
	}
	th.InvalidateBackdrop()
	th.UpdateBackdropIfNeeded(layer, image.Pt(160, 120), 4)
	if captures <= first {
		t.Fatal("invalidated backdrop did not recapture")
	}
	d := NewDimmer()
	defer d.Release()
	closed := 0
	h := testui.Harness{Size: image.Pt(160, 120), Widget: func(gtx layout.Context) layout.Dimensions {
		return d.LayoutWithAlpha(gtx, th, 255, func() { closed++ })
	}}
	if h.Frame().Size != h.Size {
		t.Fatal("dimmer does not cover viewport")
	}
	h.Click(10, 10)
	if closed != 1 {
		t.Fatalf("outside click callbacks=%d", closed)
	}
	if err := d.UpdateBackdrop(layer, image.Pt(160, 120), 4); err != nil {
		t.Fatal(err)
	}
	h.Widget = func(gtx layout.Context) layout.Dimensions { return d.LayoutWithAlpha(gtx, nil, 0, nil) }
	h.Frame()
	h.Click(10, 10)
	th.ReleaseBackdrop()
	if _, ok := th.BackdropOp(); ok {
		t.Fatal("released theme backdrop remained available")
	}
	var absent *Theme
	if _, ok := absent.BackdropOp(); ok {
		t.Fatal("nil theme produced backdrop")
	}
}
