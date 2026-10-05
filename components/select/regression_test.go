package selectcomp_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	selectcomp "github.com/bnema/gio-shadcn/components/select"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestSelectTriggerDoesNotSelectOption(t *testing.T) {
	th := theme.New()
	calls := 0
	s := selectcomp.New(selectcomp.Config{Open: true, Options: []*selectcomp.Item{selectcomp.NewItem("apple", "Apple"), selectcomp.NewItem("banana", "Banana")}, OnChange: func(string) { calls++ }})
	h := testui.Harness{Size: image.Pt(200, 200), Widget: func(gtx layout.Context) layout.Dimensions { return s.Layout(gtx, th) }}
	h.Frame()
	h.Click(15, 15)
	if calls != 0 {
		t.Fatalf("click at trigger position selects option: OnChange calls=%d", calls)
	}
}
