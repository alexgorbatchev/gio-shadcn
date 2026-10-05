package card_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCardTitleRendersText(t *testing.T) {
	c := card.NewTitle(card.TitleConfig{Text: "Account"})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, theme.New()) }}
	if d := h.Frame(); d.Size.X <= 0 || d.Size.Y <= 0 {
		t.Fatal("title rendered no text")
	}
}

func TestCardDescriptionRendersText(t *testing.T) {
	c := card.NewDescription(card.DescriptionConfig{Text: "Account settings"})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, theme.New()) }}
	if d := h.Frame(); d.Size.X <= 0 || d.Size.Y <= 0 {
		t.Fatal("description rendered no text")
	}
}

func TestCardHeaderIncludesTitleAndDescription(t *testing.T) {
	c := card.NewHeader(card.HeaderConfig{Title: "Account", Description: "Account settings"})
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions { return c.Layout(gtx, theme.New()) }}
	if d := h.Frame(); d.Size.Y < 30 || d.Size.X <= 0 {
		t.Fatal("header omitted title or description")
	}
}

func TestCardContentRunsOnceWithInset(t *testing.T) {
	c := card.NewContent(card.ContentConfig{Classes: "p-2"})
	calls := 0
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions {
		return c.Layout(gtx, theme.New(), func(gtx layout.Context) layout.Dimensions { calls++; return layout.Dimensions{Size: image.Pt(100, 20)} })
	}}
	if d := h.Frame(); d.Size != image.Pt(116, 36) || calls != 1 {
		t.Fatalf("content=%v, calls=%d", d.Size, calls)
	}
}

func TestCardFooterRunsOnceWithInset(t *testing.T) {
	c := card.NewFooter(card.FooterConfig{Classes: "px-3 py-2"})
	calls := 0
	h := testui.Harness{Size: image.Pt(300, 100), Widget: func(gtx layout.Context) layout.Dimensions {
		return c.Layout(gtx, theme.New(), func(gtx layout.Context) layout.Dimensions { calls++; return layout.Dimensions{Size: image.Pt(100, 20)} })
	}}
	if d := h.Frame(); d.Size != image.Pt(124, 36) || calls != 1 {
		t.Fatalf("footer=%v, calls=%d", d.Size, calls)
	}
}
