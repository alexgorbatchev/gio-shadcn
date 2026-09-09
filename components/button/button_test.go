package button_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func TestButtonDefault(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Button", Variant: theme.VariantDefault})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonSecondary(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Secondary", Variant: theme.VariantSecondary})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonOutline(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Outline", Variant: theme.VariantOutline})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonGhost(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Ghost", Variant: theme.VariantGhost})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonDestructive(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Destructive", Variant: theme.VariantDestructive})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonLink(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Link", Variant: theme.VariantLink})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonWithIcon(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Text: "Login with Email", Variant: theme.VariantDefault, Icon: lucide.Mail})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(160, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonIconOnly(t *testing.T) {
	th := theme.NewDark()
	btn := button.New(button.Config{Variant: theme.VariantOutline, Size: theme.SizeIcon, Icon: lucide.ChevronRight})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(36, 36))}
	dims := btn.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestButtonClickCallback(t *testing.T) {
	th := theme.NewDark()
	clicked := false
	btn := button.New(button.Config{
		Text: "Click Me",
		OnClick: func() {
			clicked = true
		},
	})
	btn.Click()

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(100, 36))}
	btn.Layout(gtx, th)

	if !clicked {
		t.Fatalf("expected OnClick to be called when Click() event is queued on button")
	}
}

func TestButtonExactHeights(t *testing.T) {
	th := theme.NewDark()

	btnDefault := button.New(button.Config{Text: "Default"})
	btnSM := button.New(button.Config{Text: "Small", Size: theme.SizeSM})
	btnLG := button.New(button.Config{Text: "Large", Size: theme.SizeLG})
	btnIcon := button.New(button.Config{Size: theme.SizeIcon, Icon: lucide.ChevronRight})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(300, 300)}}

	dimsDef := btnDefault.Layout(gtx, th)
	if dimsDef.Size.Y != gtx.Dp(unit.Dp(36)) {
		t.Errorf("expected default button height 36dp, got %dpx", dimsDef.Size.Y)
	}

	dimsSM := btnSM.Layout(gtx, th)
	if dimsSM.Size.Y != gtx.Dp(unit.Dp(32)) {
		t.Errorf("expected small button height 32dp, got %dpx", dimsSM.Size.Y)
	}

	dimsLG := btnLG.Layout(gtx, th)
	if dimsLG.Size.Y != gtx.Dp(unit.Dp(40)) {
		t.Errorf("expected large button height 40dp, got %dpx", dimsLG.Size.Y)
	}

	dimsIcon := btnIcon.Layout(gtx, th)
	if dimsIcon.Size.Y != gtx.Dp(unit.Dp(36)) || dimsIcon.Size.X != gtx.Dp(unit.Dp(36)) {
		t.Errorf("expected icon button 36x36dp, got %v", dimsIcon.Size)
	}
}

func TestButtonGroupConnectedCorners(t *testing.T) {
	th := theme.NewDark()

	btnA := button.New(button.Config{Text: "Profile", Variant: theme.VariantOutline})
	btnB := button.New(button.Config{Text: "Settings", Variant: theme.VariantOutline})
	btnC := button.New(button.Config{Text: "Messages", Variant: theme.VariantOutline})

	group := button.NewGroup(btnA, btnB, btnC)

	if btnA.Position != button.PositionFirst {
		t.Errorf("expected btnA PositionFirst, got %v", btnA.Position)
	}
	if btnB.Position != button.PositionMiddle {
		t.Errorf("expected btnB PositionMiddle, got %v", btnB.Position)
	}
	if btnC.Position != button.PositionLast {
		t.Errorf("expected btnC PositionLast, got %v", btnC.Position)
	}

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(400, 100)}}
	dims := group.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid group layout dimensions, got %v", dims.Size)
	}
}

func TestButtonGroupSelection(t *testing.T) {
	th := theme.NewDark()
	selectedIdx := -1

	btnA := button.New(button.Config{Text: "One"})
	btnB := button.New(button.Config{Text: "Two"})
	btnC := button.New(button.Config{Text: "Three"})

	group := button.NewGroup(btnA, btnB, btnC)
	group.OnSelect = func(idx int) {
		selectedIdx = idx
	}

	// Select second button (index 1)
	group.Select(1)
	if group.SelectedIndex != 1 {
		t.Errorf("expected SelectedIndex to be 1, got %d", group.SelectedIndex)
	}
	if selectedIdx != 1 {
		t.Errorf("expected OnSelect callback to receive 1, got %d", selectedIdx)
	}

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(400, 100)}}
	dims := group.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("expected valid dimensions on group layout, got %v", dims.Size)
	}
}
