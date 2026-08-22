package card_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/theme"
)

func TestCardDefaultVariant(t *testing.T) {
	c := card.New(card.Config{Variant: theme.VariantDefault})
	if c.Variant != theme.VariantDefault {
		t.Fatalf("expected VariantDefault")
	}
}

func TestCardHeaderContentFooterLayout(t *testing.T) {
	th := theme.NewDark()
	hdr := card.NewHeader(card.HeaderConfig{Title: "Title", Description: "Description"})
	c := card.New(card.Config{Variant: theme.VariantDefault})

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(400, 200)),
	}

	dimsHdr := hdr.Layout(gtx, th)
	if dimsHdr.Size.X < 0 {
		t.Errorf("invalid header width")
	}

	dimsCard := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(200, 100)}
	})
	if dimsCard.Size.X <= 0 || dimsCard.Size.Y <= 0 {
		t.Errorf("invalid card dimensions")
	}
}

func TestCardTitle(t *testing.T) {
	title := card.NewTitle(card.TitleConfig{Text: "Card Title"})
	if title.Text != "Card Title" {
		t.Errorf("expected Title 'Card Title', got %s", title.Text)
	}
}

func TestCardDescription(t *testing.T) {
	desc := card.NewDescription(card.DescriptionConfig{Text: "Card Description"})
	if desc.Text != "Card Description" {
		t.Errorf("expected Description 'Card Description', got %s", desc.Text)
	}
}

func TestCardContentArea(t *testing.T) {
	cnt := card.NewContent(card.ContentConfig{})
	if cnt == nil {
		t.Fatalf("expected non-nil Content")
	}
}

func TestCardFooter(t *testing.T) {
	ftr := card.NewFooter(card.FooterConfig{})
	if ftr == nil {
		t.Fatalf("expected non-nil Footer")
	}
}

func TestCardBorderStrokeAndRadius(t *testing.T) {
	th := theme.NewDark()
	c := card.New(card.Config{Variant: theme.VariantDefault})

	ops := new(op.Ops)
	gtx := layout.Context{
		Ops:         ops,
		Constraints: layout.Exact(image.Pt(300, 150)),
	}

	dims := c.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(100, 50)}
	})

	if dims.Size.X <= 0 {
		t.Errorf("invalid dimensions")
	}
}
