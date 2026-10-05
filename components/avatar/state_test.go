package avatar_test

import (
	"image"
	"image/color"
	"testing"

	"gioui.org/gpu/headless"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAvatarRestoresBackgroundPaint(t *testing.T) {
	th := theme.New()
	a := avatar.New(avatar.Config{Initials: "AB"})
	var ops op.Ops
	gtx := layout.Context{Ops: &ops, Metric: unit.Metric{PxPerDp: 1, PxPerSp: 1}, Constraints: layout.Constraints{Max: image.Pt(100, 80)}}
	a.Layout(gtx, th)
	area := clip.Rect(image.Rect(60, 50, 80, 70)).Push(&ops)
	paint.PaintOp{}.Add(&ops)
	area.Pop()
	w, err := headless.NewWindow(100, 80)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Release()
	if err := w.Frame(&ops); err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 100, 80))
	if err := w.Screenshot(img); err != nil {
		t.Fatal(err)
	}
	want := color.RGBAModel.Convert(th.Colors.Background).(color.RGBA)
	if got := img.RGBAAt(70, 60); got != want {
		t.Fatalf("paint state=%v, want background %v", got, want)
	}
}
