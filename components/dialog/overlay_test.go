package dialog_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/dialog"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestDialogQueuesAtWindowRootOnce(t *testing.T) {
	th := theme.New()
	calls := 0
	cancelled := 0
	d := dialog.New(dialog.Config{Open: true, TriggerText: "Open", Title: "Dialog", OnCancel: func() { cancelled++ }, Content: func(gtx layout.Context) layout.Dimensions {
		calls++
		return layout.Dimensions{Size: image.Pt(200, 100)}
	}})
	h := testui.Harness{Size: image.Pt(600, 400), Widget: func(gtx layout.Context) layout.Dimensions {
		child := gtx
		child.Constraints.Max = image.Pt(100, 40)
		d.Layout(child, th)
		if d.Open && !th.HasOverlays() {
			t.Error("dialog not queued at root")
		}
		return th.RenderOverlays(gtx)
	}}
	if h.Frame().Size != h.Size || calls != 1 {
		t.Fatalf("root bounds/content calls=%d", calls)
	}
	h.Click(300, 200)
	if !d.Open {
		t.Fatal("interior dismissed")
	}
	h.Click(10, 390)
	if d.Open || cancelled != 1 {
		t.Fatal("root outside click failed")
	}
}
