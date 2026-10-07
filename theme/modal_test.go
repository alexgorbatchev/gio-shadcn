package theme_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/components/command"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/components/sheet"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSheetKeyboardLifecycle(t *testing.T) {
	testModalKeyboardLifecycle(t, "sheet")
}

func TestDrawerKeyboardLifecycle(t *testing.T) {
	testModalKeyboardLifecycle(t, "drawer")
}

func TestCommandKeyboardLifecycle(t *testing.T) {
	testModalKeyboardLifecycle(t, "command")
}

func testModalKeyboardLifecycle(t *testing.T, name string) {
	t.Helper()
	th := theme.New()
	var background widget.Editor
	var render layout.Widget
	var activate func()
	var open *bool
	closes := 0
	switch name {
	case "sheet":
		s := sheet.New(sheet.Config{TriggerText: "Open", OnClose: func() { closes++ }})
		open, activate = &s.Open, s.TriggerButton.Click
		render = func(gtx layout.Context) layout.Dimensions { return s.Layout(gtx, th) }
	case "drawer":
		d := drawer.New(drawer.Config{TriggerText: "Open", OnClose: func() { closes++ }})
		open, activate = &d.Open, d.TriggerButton.Click
		render = func(gtx layout.Context) layout.Dimensions { return d.Layout(gtx, th) }
	case "command":
		c := command.New(command.Config{TriggerText: "Open", Items: []*command.Item{command.NewItem("Action", "")}})
		open, activate = &c.Open, c.TriggerButton.Click
		render = func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, th) }
	}
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		return th.LayoutRoot(gtx, func(gtx layout.Context) layout.Dimensions {
			material.Editor(th.MaterialTheme, &background, "Background").Layout(gtx)
			return render(gtx)
		})
	}}
	h.Frame()
	activate()
	for range 80 {
		h.Frame()
	}
	for _, mods := range []key.Modifiers{0, key.ModShift} {
		for range 6 {
			h.Key(key.NameTab, mods)
			if h.Router.Source().Focused(&background) {
				t.Fatal("native focus traversal escaped the modal")
			}
		}
	}
	h.Key(key.NameEscape, 0)
	if *open {
		t.Fatal("Escape did not dismiss the modal")
	}
	if name != "command" && closes != 1 {
		t.Fatalf("close callback count=%d", closes)
	}
	h.Key(key.NameReturn, 0)
	if !*open {
		t.Fatal("Enter did not reopen the modal from its restored trigger")
	}
}
