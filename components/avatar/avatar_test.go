package avatar_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAvatarBasicDemo(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials: "CN",
	})
	if av.Initials != "CN" {
		t.Errorf("expected Initials CN, got %s", av.Initials)
	}
}

func TestAvatarBadgeDemo(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials:   "CN",
		ShowBadge:  true,
		BadgeColor: "green",
	})
	if !av.ShowBadge || av.BadgeColor != "green" {
		t.Errorf("expected green badge enabled")
	}
}

func TestAvatarBadgeIconDemo(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials:  "PP",
		ShowBadge: true,
	})
	if !av.ShowBadge {
		t.Errorf("expected badge to be enabled")
	}
}

func TestAvatarSizeDemo(t *testing.T) {
	sm := avatar.New(avatar.Config{Initials: "SM", Size: unit.Dp(32)})
	md := avatar.New(avatar.Config{Initials: "MD", Size: unit.Dp(40)})
	lg := avatar.New(avatar.Config{Initials: "LG", Size: unit.Dp(56)})

	if sm.Size != unit.Dp(32) || md.Size != unit.Dp(40) || lg.Size != unit.Dp(56) {
		t.Errorf("size mismatch in avatar configs")
	}
}

func TestAvatarGroupDemo(t *testing.T) {
	g1 := avatar.New(avatar.Config{Initials: "CN"})
	g2 := avatar.New(avatar.Config{Initials: "LR"})
	g3 := avatar.New(avatar.Config{Initials: "+3"})

	if g1.Initials != "CN" || g2.Initials != "LR" || g3.Initials != "+3" {
		t.Errorf("avatar group initials mismatch")
	}
}

func TestAvatarLayout(t *testing.T) {
	th := theme.NewDark()
	av := avatar.New(avatar.Config{
		Initials:  "AG",
		ShowBadge: true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(100, 100)),
	}
	dims := av.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions from Avatar.Layout: %v", dims.Size)
	}
}
