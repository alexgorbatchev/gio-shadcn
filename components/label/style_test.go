package label_test

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
	"image"
	"testing"
)

func TestExplicitLabelSizeOverridesPreset(t *testing.T) {
	l := label.NewLabel(label.WithLabelText("Sample"), label.WithTextStyle(theme.TextStyle{Size: 32}))
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return l.Layout(gtx, theme.New()) }}
	if d := h.Frame(); d.Size.Y < 32 {
		t.Fatalf("explicit 32sp size ignored: %v", d.Size)
	}
}
