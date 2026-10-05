package accordion_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"github.com/bnema/gio-shadcn/components/accordion"
	"github.com/bnema/gio-shadcn/theme"
)

func TestAccordionContentRunsOnce(t *testing.T) {
	th := theme.New()
	calls := 0
	item := accordion.NewItemConfig(accordion.ItemConfig{Title: "Custom", Expanded: true, ContentWidget: func(gtx layout.Context) layout.Dimensions {
		calls++
		return layout.Dimensions{Size: image.Pt(100, 100)}
	}})
	a := accordion.New(accordion.Config{Items: []*accordion.Item{item}})
	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Constraints{Max: image.Pt(400, 400)}}
	a.Layout(gtx, th)
	if calls != 1 {
		t.Fatalf("Accordion runs stateful ContentWidget %d times per layout", calls)
	}
}
