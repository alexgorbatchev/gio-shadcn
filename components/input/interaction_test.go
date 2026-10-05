package input_test

import (
	"image"
	"testing"

	"gioui.org/io/key"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestFocusChangeAndSubmitCallbacks(t *testing.T) {
	focus, blur, changes, submits := 0, 0, 0, 0
	i := input.NewInput(input.WithInputVariant(input.InputFilled), input.WithInputSize(input.InputSizeLarge), input.WithLabel("Name"), input.WithHelper("Display name"), input.WithRequired(true), input.WithOnChange(func(string) { changes++ }), input.WithOnSubmit(func() { submits++ }))
	i.WithOnFocus(func() { focus++ }).WithOnBlur(func() { blur++ })
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions {
		return i.Layout(gtx, theme.New())
	}}
	if d := h.Frame(); d.Size.Y < 52 {
		t.Fatalf("large input height=%d", d.Size.Y)
	}
	h.Click(20, 20)
	h.Edit("Alex", 0, 0)
	if i.Text() != "Alex" || i.Value != "Alex" || focus != 1 || changes != 1 {
		t.Fatalf("text=%q, focus=%d, change=%d", i.Text(), focus, changes)
	}
	h.Router.Queue(key.Event{Name: key.NameReturn, State: key.Press})
	h.Frame()
	if submits != 1 {
		t.Fatalf("submit callbacks=%d", submits)
	}
	h.Router.Source().Execute(key.FocusCmd{})
	h.Frame()
	if blur != 1 {
		t.Fatalf("blur callbacks=%d", blur)
	}
	i.SetText("Programmatic")
	h.Frame()
	if i.Text() != "Programmatic" || changes != 1 {
		t.Fatal("SetText lost value or fired user change callback")
	}
}

func TestInputTypeSwitchClearsFilterAndMask(t *testing.T) {
	i := input.Number("Number")
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return i.Layout(gtx, theme.New()) }}
	h.Frame()
	h.Click(20, 20)
	h.Edit("12abc", 0, 0)
	if i.Text() != "12" {
		t.Fatalf("number filter text=%q", i.Text())
	}
	i.Type = input.InputPassword
	i.SetText("")
	h.Frame()
	h.Edit("secret!", 0, 0)
	if i.Text() != "secret!" {
		t.Fatalf("password retained numeric filter: %q", i.Text())
	}
}

func TestDisabledInputCannotEdit(t *testing.T) {
	i := input.Email("Email").WithVariant(input.InputGhost).WithSize(input.InputSizeSmall).WithLabel("Email").WithHelper("Contact").WithRequired(true).WithError("Invalid").WithDisabled(true).WithOnChange(func(string) { t.Fatal("disabled OnChange") }).WithOnSubmit(func() { t.Fatal("disabled OnSubmit") })
	i.SetText("before")
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return i.Layout(gtx, theme.New()) }}
	h.Frame()
	h.Click(20, 20)
	h.Edit("after", 0, 0)
	if i.Text() != "before" {
		t.Fatal("disabled input changed")
	}
	state := i.Update(layout.Context{})
	if !state.IsDisabled() || state.IsPressed() || state.IsHovered() {
		t.Fatal("disabled input state incorrect")
	}
	_ = state.IsActive()
}
