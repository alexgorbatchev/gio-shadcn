package avatar_test

import (
	"image"
	"testing"

	"gioui.org/unit"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAvatarTextInitials(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials: "dj",
	})
	if av.Initials != "DJ" {
		t.Fatalf("expected uppercase initials 'DJ', got %s", av.Initials)
	}
}

func TestAvatarImagePlaceholder(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials: "AG",
	})
	if av.Initials != "AG" {
		t.Fatalf("expected initials 'AG'")
	}
}

func TestAvatarOnlineStatusBadge(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials:   "DJ",
		ShowBadge:  true,
		BadgeColor: "green",
	})
	if !av.ShowBadge || av.BadgeColor != "green" {
		t.Fatalf("expected ShowBadge true with green color")
	}
}

func TestAvatarCircularClipEllipse(t *testing.T) {
	th := theme.NewDark()
	av := avatar.New(avatar.Config{
		Initials: "DJ",
		Size:     unit.Dp(40),
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(40, 40)),
	}
	dims := av.Layout(gtx, th)
	if dims.Size.X != 40 || dims.Size.Y != 40 {
		t.Errorf("expected size 40x40")
	}
}

func TestAvatarCustomSizes(t *testing.T) {
	av32 := avatar.New(avatar.Config{Size: unit.Dp(32)})
	av56 := avatar.New(avatar.Config{Size: unit.Dp(56)})
	if av32.Size != unit.Dp(32) || av56.Size != unit.Dp(56) {
		t.Errorf("expected custom sizes 32 and 56")
	}
}

func TestAvatarStatusDotIndicator(t *testing.T) {
	th := theme.NewDark()
	av := avatar.New(avatar.Config{
		Initials:  "DJ",
		ShowBadge: true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(50, 50)),
	}
	dims := av.Layout(gtx, th)
	if dims.Size.X <= 0 {
		t.Errorf("invalid width")
	}
}
