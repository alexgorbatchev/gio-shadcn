package dropdownmenu_test

import (
	"image"
	"image/color"
	"testing"

	"gioui.org/gpu/headless"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/widget"
	"github.com/bnema/gio-shadcn/components/dropdownmenu"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDropdownOverlayEscapesClipAndCoversLaterSibling(t *testing.T) {
	th := theme.NewDark()
	selected, behind := 0, 0
	var sibling widget.Clickable
	item := dropdownmenu.NewItem("Profile", "")
	item.OnSelect = func() { selected++ }
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: true, Items: []*dropdownmenu.Item{item},
		Trigger: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 40)}
		},
	})
	h := testui.Harness{Size: image.Pt(300, 200), Widget: func(gtx layout.Context) layout.Dimensions {
		origin := op.Offset(image.Pt(40, 20)).Push(gtx.Ops)
		parentClip := clip.Rect(image.Rect(0, 0, 100, 40)).Push(gtx.Ops)
		dims := dm.Layout(gtx, th)
		parentClip.Pop()
		origin.Pop()
		if sibling.Clicked(gtx) {
			behind++
		}
		later := op.Offset(image.Pt(0, 60)).Push(gtx.Ops)
		sibling.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			theme.DrawRRectBackground(gtx, image.Rect(0, 0, 300, 140), 0, th.Colors.Destructive)
			return layout.Dimensions{Size: image.Pt(300, 140)}
		})
		later.Pop()
		return dims
	}}
	h.Frame()
	img := dropdownScreenshot(t, &h)
	want := color.RGBAModel.Convert(th.Colors.Popover).(color.RGBA)
	if got := img.RGBAAt(45, 72); got != want {
		t.Errorf("floating menu pixel=%v, want popover %v", got, want)
	}
	h.Click(45, 80)
	if selected != 1 || behind != 0 || dm.Open {
		t.Fatalf("menu selection=%d, covered sibling clicks=%d, open=%v", selected, behind, dm.Open)
	}
}

func TestDisabledDropdownDoesNotEscapeAboveModal(t *testing.T) {
	th := theme.NewDark()
	dm := dropdownmenu.New(dropdownmenu.Config{
		Open: true, Items: []*dropdownmenu.Item{dropdownmenu.NewItem("Profile", "")},
		Trigger: func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: image.Pt(100, 40)}
		},
	})
	h := testui.Harness{Size: image.Pt(300, 200), Widget: func(gtx layout.Context) layout.Dimensions {
		dims := dm.Layout(gtx.Disabled(), th)
		theme.DrawRRectBackground(gtx, image.Rectangle{Max: gtx.Constraints.Max}, 0, th.Colors.Destructive)
		return dims
	}}
	h.Frame()
	img := dropdownScreenshot(t, &h)
	want := color.RGBAModel.Convert(th.Colors.Destructive).(color.RGBA)
	if got := img.RGBAAt(5, 52); got != want {
		t.Fatalf("disabled popup painted above modal: pixel=%v, want %v", got, want)
	}
}

func dropdownScreenshot(t *testing.T, h *testui.Harness) *image.RGBA {
	t.Helper()
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
	return img
}
