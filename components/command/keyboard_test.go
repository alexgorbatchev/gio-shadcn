package command

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCommandArrowKeysSkipDisabledAndEnterActivates(t *testing.T) {
	th := theme.New()
	var selected []int
	c := New(Config{Open: true, TriggerText: "Commands", Items: []*Item{
		NewItem("Alpha", ""), NewItemFull("Disabled", "", "", nil, true), NewItem("Gamma", ""),
	}, OnSelectItem: func(i int) { selected = append(selected, i) }})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) })
	}}
	h.Frame()
	h.Frame()
	h.Key(key.NameDownArrow, 0)
	h.Key(key.NameReturn, 0)
	if len(selected) != 1 || selected[0] != 2 || c.Open {
		t.Fatalf("keyboard activation: selections=%v, modal open=%v", selected, c.Open)
	}
}

func TestCommandKeyboardUsesFilteredResults(t *testing.T) {
	for _, query := range []string{"gAm", "missing"} {
		t.Run(query, func(t *testing.T) {
			th := theme.New()
			selected := -1
			c := New(Config{Open: true, TriggerText: "Commands", Items: []*Item{NewItem("Alpha", ""), NewItem("Gamma", "")}, OnSelectItem: func(i int) { selected = i }})
			h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
				return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) })
			}}
			h.Frame()
			h.Frame()
			h.Edit(query, 0, 0)
			h.Key(key.NameDownArrow, 0)
			h.Key(key.NameEnter, 0)
			if query == "gAm" && (selected != 1 || c.Open) {
				t.Fatalf("filtered command selected=%d, open=%v", selected, c.Open)
			}
			if query == "missing" && (selected != -1 || !c.Open) {
				t.Fatalf("empty result activated: selected=%d, open=%v", selected, c.Open)
			}
		})
	}
}

func TestCommandHomeEndAndUpNavigation(t *testing.T) {
	for _, test := range []struct {
		name string
		keys []key.Name
		want int
	}{
		{"end", []key.Name{key.NameEnd}, 2},
		{"home", []key.Name{key.NameEnd, key.NameHome}, 0},
		{"up", []key.Name{key.NameEnd, key.NameUpArrow}, 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			th := theme.New()
			selected := -1
			c := New(Config{Open: true, TriggerText: "Commands", Items: []*Item{NewItem("First", ""), NewItemFull("Disabled", "", "", nil, true), NewItem("Last", "")}, OnSelectItem: func(i int) { selected = i }})
			h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
				return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) })
			}}
			h.Frame()
			h.Frame()
			for _, name := range test.keys {
				h.Key(name, 0)
			}
			h.Key(key.NameReturn, 0)
			if selected != test.want {
				t.Fatalf("selected=%d, want %d", selected, test.want)
			}
		})
	}
}

func TestCommandArrowNavigationFromFocusedItem(t *testing.T) {
	th := theme.New()
	selected := -1
	c := New(Config{Open: true, TriggerText: "Commands", Items: []*Item{NewItem("First", ""), NewItem("Last", "")}, OnSelectItem: func(i int) { selected = i }})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) })
	}}
	h.Frame()
	h.Frame()
	h.Key(key.NameTab, 0)
	if !h.Router.Source().Focused(c.Items[0].clickable) {
		t.Fatal("Tab did not focus the first command")
	}
	h.Key(key.NameDownArrow, 0)
	if !h.Router.Source().Focused(c.Items[1].clickable) {
		t.Error("Down did not transfer native focus to the next command")
	}
	h.Key(key.NameReturn, 0)
	if selected != 1 || c.Open {
		t.Fatalf("focused-row activation selected=%d, open=%v", selected, c.Open)
	}
}

func TestCommandNavigationBoundaryAndLoop(t *testing.T) {
	for _, test := range []struct {
		name string
		loop bool
		keys []key.Name
		want int
	}{
		{"clamp-first", false, []key.Name{key.NameUpArrow}, 0},
		{"clamp-last", false, []key.Name{key.NameEnd, key.NameDownArrow}, 2},
		{"wrap-first", true, []key.Name{key.NameUpArrow}, 2},
		{"wrap-last", true, []key.Name{key.NameEnd, key.NameDownArrow}, 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			th := theme.New()
			selected := -1
			c := New(Config{Open: true, TriggerText: "Commands", Items: []*Item{NewItem("First", ""), NewItemFull("Disabled", "", "", nil, true), NewItem("Last", "")}, OnSelectItem: func(i int) { selected = i }})
			c.Loop = test.loop
			h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
				return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) })
			}}
			h.Frame()
			h.Frame()
			for _, name := range test.keys {
				h.Key(name, 0)
			}
			h.Key(key.NameReturn, 0)
			if selected != test.want {
				t.Fatalf("selected=%d, want %d", selected, test.want)
			}
		})
	}
}
