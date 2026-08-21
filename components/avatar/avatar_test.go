package avatar_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/avatar"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAvatarCreation(t *testing.T) {
	av := avatar.New(avatar.Config{
		Initials:  "dj",
		ShowBadge: true,
	})

	if av.Initials != "DJ" {
		t.Errorf("expected Initials to be 'DJ', got %s", av.Initials)
	}
}

func TestAvatarLayout(t *testing.T) {
	th := theme.NewDark()
	av := avatar.New(avatar.Config{
		Initials:  "AG",
		ShowBadge: true,
	})

	gtx := layout.Context{
		Ops: new(op.Ops),
		Constraints: layout.Exact(image.Pt(50, 50)),
	}
	dims := av.Layout(gtx, th)

	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Avatar.Layout")
	}
}
