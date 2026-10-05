package utils_test

import (
	"image/color"
	"testing"

	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/utils"
)

func TestUtilityClassesApplyInOrder(t *testing.T) {
	style := utils.ParseClasses("p-4 px-2 py-3 pt-1 pr-5 pb-6 pl-8", "m-2 mx-3 my-4 bg-blue border border-red rounded-lg opacity-75")
	if style.Padding != (layout.Inset{Top: 4, Right: 20, Bottom: 24, Left: 32}) {
		t.Fatalf("padding=%v", style.Padding)
	}
	if style.Margin != (layout.Inset{Top: 16, Right: 12, Bottom: 16, Left: 12}) {
		t.Fatalf("margin=%v", style.Margin)
	}
	if style.Background != (color.NRGBA{R: 59, G: 130, B: 246, A: 255}) {
		t.Fatalf("background=%v", style.Background)
	}
	if style.Border.Width != 1 || style.Border.Color.R != 239 || style.Radius != 8 || style.Opacity != .75 {
		t.Fatalf("decoration=%+v", style)
	}
}

func TestUnknownClassesPreserveEarlierStyles(t *testing.T) {
	style := utils.ParseClasses("p-2 rounded bg-white opacity-50", "p-bad px-bad py-bad pt-bad pr-bad pb-bad pl-bad m-bad mx-bad my-bad rounded-bad bg-bad border-bad opacity-bad unknown")
	if style.Padding != layout.UniformInset(8) || style.Radius != 4 || style.Background.R != 255 || style.Opacity != .5 {
		t.Fatalf("invalid classes changed styles: %+v", style)
	}
	if got := utils.ClassNames("a", "", "b"); got != "a b" {
		t.Fatalf("class names=%q", got)
	}
}

func TestUtilityScaleValues(t *testing.T) {
	for _, tt := range []struct {
		class  string
		radius unit.Dp
	}{{"none", 0}, {"sm", 2}, {"md", 6}, {"lg", 8}, {"xl", 12}, {"2xl", 16}, {"3xl", 24}, {"full", 9999}} {
		t.Run(tt.class, func(t *testing.T) {
			if got := utils.ParseClasses("rounded-" + tt.class).Radius; got != tt.radius {
				t.Fatalf("radius=%v", got)
			}
		})
	}
	for _, name := range []string{"transparent", "black", "white", "red", "green", "blue", "yellow", "purple", "pink", "indigo", "gray"} {
		t.Run(name, func(t *testing.T) {
			style := utils.ParseClasses("bg-" + name)
			if name == "transparent" {
				if style.Background.A != 0 {
					t.Fatal(style.Background)
				}
			} else if style.Background.A != 255 {
				t.Fatal(style.Background)
			}
		})
	}
}
