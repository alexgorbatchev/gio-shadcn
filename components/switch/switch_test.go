package switchcomp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	switchcomp "github.com/bnema/gio-shadcn/components/switch"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSwitchOff(t *testing.T) {
	th := theme.NewDark()
	s := switchcomp.New(switchcomp.Config{Value: false})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 30))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestSwitchOn(t *testing.T) {
	th := theme.NewDark()
	s := switchcomp.New(switchcomp.Config{Value: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 30))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}

func TestSwitchDisabled(t *testing.T) {
	th := theme.NewDark()
	s := switchcomp.New(switchcomp.Config{Value: true, Disabled: true})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(50, 30))}
	dims := s.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions")
	}
}
