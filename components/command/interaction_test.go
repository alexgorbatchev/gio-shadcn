package command

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSearchFiltersAndDisabledSelection(t *testing.T) {
	th := theme.New()
	selected := -1
	a, b := NewItemFull("Alpha", "A", "Files", lucide.File, false), NewItemFull("Beta", "B", "Files", lucide.File, true)
	c := New(Config{Items: []*Item{a, b}, Classes: "bg-blue-500", OnSelectItem: func(i int) { selected = i }})
	h := testui.Harness{Size: image.Pt(400, 300), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) }}
	full := h.Frame()
	b.clickable.Click()
	h.Frame()
	if selected != -1 {
		t.Fatal("disabled selected")
	}
	c.searchEditor.SetText("AlPhA")
	filtered := h.Frame()
	if filtered.Size.Y >= full.Size.Y {
		t.Fatal("search did not filter case-insensitively")
	}
	a.clickable.Click()
	h.Frame()
	if selected != 0 {
		t.Fatal("filtered selection lost original index")
	}
}

func TestModalQueuesAndSelectCloses(t *testing.T) {
	th := theme.New()
	selected := -1
	a := NewItem("Alpha", "")
	c := New(Config{Open: true, TriggerText: "Open", Items: []*Item{a}, OnSelectItem: func(i int) { selected = i }})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		child := gtx
		child.Constraints.Max = image.Pt(100, 40)
		c.Layout(child, th)
		if c.Open && !th.HasOverlays() {
			t.Error("modal not queued")
		}
		return th.RenderOverlays(gtx)
	}}
	if h.Frame().Size != h.Size {
		t.Fatal("modal failed viewport bounds")
	}
	a.clickable.Click()
	h.Frame()
	if c.Open || selected != 0 {
		t.Fatal("modal selection did not close")
	}
}
