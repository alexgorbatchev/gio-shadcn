package theme_test

import (
	"image"
	"image/color"
	"testing"
	"time"

	"gioui.org/gpu/headless"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func settleScheduledFrames(t *testing.T, h *testui.Harness) {
	t.Helper()
	for range 240 {
		at, requested := h.Router.WakeupTime()
		if !requested {
			return
		}
		h.Now = h.Now.Add(time.Second / 60)
		if at.After(h.Now) {
			h.Now = at
		}
		h.Frame()
	}
	t.Fatal("animation never stopped requesting frames")
}

func TestSheetCloseReleasesDimmer(t *testing.T) {
	testAnimatedModalClose(t, []string{"right", "left", "top", "bottom"})
}

func TestDrawerCloseReleasesDimmer(t *testing.T) {
	testAnimatedModalClose(t, []string{"drawer"})
}

func testAnimatedModalClose(t *testing.T, names []string) {
	t.Helper()
	for _, name := range names {
		for _, closeWith := range []string{"button", "backdrop", "escape"} {
			t.Run(name+"/"+closeWith, func(t *testing.T) {
				th := theme.NewDark()
				var render layout.Widget
				var open *bool
				closes, clicks := 0, 0
				closeX, closeY := float32(360), float32(44)
				outsideX, outsideY := float32(10), float32(10)
				if name == "drawer" {
					d := drawer.New(drawer.Config{Open: true, Height: 120, Title: "Drawer", OnClose: func() { closes++ }})
					open = &d.Open
					render = func(gtx layout.Context) layout.Dimensions { return d.Layout(gtx, th) }
					closeY = 230
				} else {
					side := sheet.SideRight
					switch name {
					case "left":
						side, closeX, outsideX = sheet.SideLeft, 80, 390
					case "top":
						side, outsideY = sheet.SideTop, 290
					case "bottom":
						side, closeY = sheet.SideBottom, 224
					}
					s := sheet.New(sheet.Config{Open: true, Width: 120, Height: 120, Side: side, Title: "Sheet", OnClose: func() { closes++ }})
					open = &s.Open
					render = func(gtx layout.Context) layout.Dimensions { return s.Layout(gtx, th) }
				}
				var background widget.Clickable
				h := testui.Harness{Size: image.Pt(400, 300), Widget: func(gtx layout.Context) layout.Dimensions {
					return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions {
						if background.Clicked(gtx) {
							clicks++
						}
						background.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							theme.DrawRRectBackground(gtx, image.Rectangle{Max: gtx.Constraints.Max}, 0, th.Colors.Background)
							return layout.Dimensions{Size: gtx.Constraints.Max}
						})
						return render(gtx)
					})
				}}
				h.Frame()
				settleScheduledFrames(t, &h)
				switch closeWith {
				case "button":
					h.Click(closeX, closeY)
				case "backdrop":
					h.Click(outsideX, outsideY)
				case "escape":
					h.Key(key.NameEscape, 0)
				}
				if *open || closes != 1 {
					t.Fatalf("close did not fire once: open=%v, callbacks=%d", *open, closes)
				}
				settleScheduledFrames(t, &h)
				w, err := headless.NewWindow(h.Size.X, h.Size.Y)
				if err != nil {
					t.Fatal(err)
				}
				defer w.Release()
				if err := w.Frame(&h.Ops); err != nil {
					t.Fatal(err)
				}
				img := image.NewRGBA(image.Rectangle{Max: h.Size})
				if err := w.Screenshot(img); err != nil {
					t.Fatal(err)
				}
				want := color.RGBAModel.Convert(th.Colors.Background).(color.RGBA)
				if got := img.RGBAAt(int(outsideX), int(outsideY)); got != want {
					t.Errorf("idle frame still contains dimmer: pixel=%v, want %v", got, want)
				}
				h.Click(outsideX, outsideY)
				if clicks != 1 || closes != 1 {
					t.Fatalf("background clicks=%d, close callbacks=%d after dismissal", clicks, closes)
				}
			})
		}
	}
}
