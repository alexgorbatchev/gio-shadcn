package titlebar_test

import (
	"image"
	"image/color"
	"testing"

	"gioui.org/gpu/headless"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

func TestTitlebarClassChangesRenderedFill(t *testing.T) {
	th := theme.New()
	tb := titlebar.NewTitleBar(titlebar.WithTitle("Title"), titlebar.WithControls(false), titlebar.WithClasses("bg-red"))
	h := testui.Harness{Size: image.Pt(300, 40), Widget: func(gtx layout.Context) layout.Dimensions { return tb.Layout(gtx, th, nil) }}
	h.Frame()
	w, err := headless.NewWindow(300, 40)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Release()
	if err := w.Frame(&h.Ops); err != nil {
		t.Fatal(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 300, 40))
	if err := w.Screenshot(img); err != nil {
		t.Fatal(err)
	}
	want := color.RGBAModel.Convert(utils.ParseClasses("bg-red").Background).(color.RGBA)
	if got := img.RGBAAt(250, 10); got != want {
		t.Fatalf("class fill=%v want=%v", got, want)
	}
}
