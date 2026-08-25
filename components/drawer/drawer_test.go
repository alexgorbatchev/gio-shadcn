package drawer_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/drawer"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDrawerCreation(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title:       "Pick a delivery time",
		Description: "Standard delivery 25-35 min",
		Open:        true,
	})
	if !dr.Open {
		t.Errorf("expected Open to be true")
	}
}

func TestDrawerProfileDialog(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title:       "Edit profile",
		Description: "Make changes to your profile",
		Open:        false,
	})
	if dr.Title != "Edit profile" {
		t.Errorf("unexpected title: %s", dr.Title)
	}
}

func TestDrawerMoveGoal(t *testing.T) {
	dr := drawer.New(drawer.Config{
		Title:       "Move Goal",
		Description: "Set daily activity goal",
		Open:        true,
	})
	if dr.Description != "Set daily activity goal" {
		t.Errorf("unexpected description: %s", dr.Description)
	}
}

func TestDrawerLayout(t *testing.T) {
	th := theme.NewDark()
	dr := drawer.New(drawer.Config{
		Title:       "Telemetry Drawer",
		Description: "CPU Usage: 2.1%",
		Open:        true,
	})
	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(600, 400)),
	}
	dims := dr.Layout(gtx, th)
	if dims.Size.X < 0 || dims.Size.Y < 0 {
		t.Errorf("invalid dimensions returned from Drawer.Layout")
	}
}
