package theme_test

import (
	"image"
	"image/color"
	"testing"

	"gioui.org/gpu/headless"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/bnema/gio-shadcn/components/scrollarea"
	"github.com/bnema/gio-shadcn/theme"
)

func TestRoundedFillClampsRadiusAndIsolatesClip(t *testing.T) {
	w, err := headless.NewWindow(100, 40)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Release()
	var ops op.Ops
	gtx := layout.Context{Ops: &ops}
	red := color.NRGBA{R: 255, A: 255}
	blue := color.NRGBA{B: 255, A: 255}
	theme.DrawRRectBackground(gtx, image.Rect(0, 0, 40, 20), 9999, red)
	theme.DrawRRectBackground(gtx, image.Rect(60, 0, 100, 20), 0, blue)
	if err := w.Frame(&ops); err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 100, 40))
	if err := w.Screenshot(img); err != nil {
		t.Fatal(err)
	}
	if got := img.RGBAAt(20, 10); got.R < 250 || got.A < 250 {
		t.Fatalf("capsule center = %v", got)
	}
	if got := img.RGBAAt(0, 0); got.A != 0 {
		t.Fatalf("rounded corner = %v", got)
	}
	if got := img.RGBAAt(80, 10); got.B < 250 || got.R != 0 {
		t.Fatalf("subsequent fill = %v", got)
	}
	if got := img.RGBAAt(50, 10); got.A != 0 {
		t.Fatalf("clip leaked into gap: %v", got)
	}
}

func TestEmptyLayoutsRestorePaintColor(t *testing.T) {
	th := theme.New()
	for name, widget := range map[string]layout.Widget{
		"dimmer": func(gtx layout.Context) layout.Dimensions {
			d := theme.NewDimmer()
			defer d.Release()
			return d.LayoutWithAlpha(gtx, th, 160, nil)
		},
		"overlays":   th.RenderOverlays,
		"scrollarea": func(gtx layout.Context) layout.Dimensions { return scrollarea.New(scrollarea.Config{}).Layout(gtx, th) },
	} {
		t.Run(name, func(t *testing.T) {
			w, err := headless.NewWindow(100, 80)
			if err != nil {
				t.Fatal(err)
			}
			defer w.Release()
			var ops op.Ops
			gtx := layout.Context{Ops: &ops, Constraints: layout.Constraints{Max: image.Pt(100, 80)}}
			paint.ColorOp{Color: th.Colors.Primary}.Add(&ops)
			widget(gtx)
			area := clip.Rect(image.Rect(60, 50, 80, 70)).Push(&ops)
			paint.PaintOp{}.Add(&ops)
			area.Pop()
			if err := w.Frame(&ops); err != nil {
				t.Fatal(err)
			}
			img := image.NewRGBA(image.Rect(0, 0, 100, 80))
			if err := w.Screenshot(img); err != nil {
				t.Fatal(err)
			}
			want := color.RGBAModel.Convert(th.Colors.Background).(color.RGBA)
			if got := img.RGBAAt(70, 60); got != want {
				t.Fatalf("paint state=%v, want %v", got, want)
			}
		})
	}
}
