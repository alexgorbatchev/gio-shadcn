package demo

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"gioui.org/app"
	"gioui.org/gpu/headless"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func TestGalleryPagesRender(t *testing.T) {
	ids := []string{"accordion", "alert", "aspectratio", "avatar", "badge", "breadcrumb", "button", "card", "carousel", "checkbox", "collapsible", "command", "dialog", "drawer", "dropdownmenu", "empty", "hovercard", "input", "inputotp", "label", "numberinput", "pagination", "popover", "progress", "radio", "resizable", "scrollarea", "select", "separator", "sheet", "skeleton", "slider", "spinner", "switch", "table", "tabs", "textarea", "titlebar", "toast", "togglegroup", "tooltip", "tree"}
	for _, id := range ids {
		t.Run(id, func(t *testing.T) {
			th := theme.NewDark()
			h := testui.Harness{Size: image.Pt(1000, 20000), Widget: func(gtx layout.Context) layout.Dimensions {
				return renderComponentGalleryPage(gtx, th, id)
			}}
			dims := h.Frame()
			if dims.Size.X <= 0 || dims.Size.Y <= 0 {
				t.Fatalf("page has no content: %v", dims.Size)
			}
			if dims.Size.X > h.Size.X || dims.Size.Y > h.Size.Y {
				t.Fatalf("page exceeds viewport: %v", dims.Size)
			}
			// Rendering again exercises persistent component state.
			if next := h.Frame(); next.Size != dims.Size {
				t.Fatalf("idle page size changed: %v -> %v", dims.Size, next.Size)
			}
		})
	}
}

func TestGalleryFrameAndNavigation(t *testing.T) {
	w := new(app.Window)
	g := newGallery(w)
	defer g.th.ReleaseBackdrop()
	h := testui.Harness{Size: image.Pt(1200, 800), Widget: func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints = layout.Exact(image.Pt(1200, 800))
		return g.layout(gtx, w, image.Pt(1200, 800))
	}}
	if d := h.Frame(); d.Size != h.Size {
		t.Fatalf("gallery frame=%v", d.Size)
	}
	first := g.activeID
	// The third list item is the first component; the next row selects another.
	h.Click(80, 200)
	if g.activeID == first {
		t.Fatal("sidebar click did not change the active component")
	}
	g.themeToggle.Click()
	h.Frame()
	if g.th.IsDark {
		t.Fatal("theme toggle did not switch to light")
	}
	g.themeToggle.Click()
	h.Frame()
	if !g.th.IsDark {
		t.Fatal("theme toggle did not switch back to dark")
	}
	if len(g.items) != 42 {
		t.Fatalf("gallery exposes %d components", len(g.items))
	}
}

func TestUnknownGalleryPageIsEmpty(t *testing.T) {
	h := testui.Harness{Size: image.Pt(800, 600), Widget: func(gtx layout.Context) layout.Dimensions {
		return renderComponentGalleryPage(gtx, theme.New(), "unknown")
	}}
	if dims := h.Frame(); dims.Size != (image.Point{}) {
		t.Fatalf("unknown page has content: %v", dims)
	}
}

func TestGalleryGPURendersBothThemes(t *testing.T) {
	window, err := headless.NewWindow(1200, 800)
	if err != nil {
		t.Fatal(err)
	}
	defer window.Release()
	g := newGallery(new(app.Window))
	defer g.th.ReleaseBackdrop()
	h := testui.Harness{Size: image.Pt(1200, 800), Widget: func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints = layout.Exact(image.Pt(1200, 800))
		return g.layout(gtx, nil, image.Pt(1200, 800))
	}}
	for _, name := range []string{"dark", "light"} {
		h.Frame()
		if err := window.Frame(&h.Ops); err != nil {
			t.Fatal(err)
		}
		img := image.NewRGBA(image.Rect(0, 0, 1200, 800))
		if err := window.Screenshot(img); err != nil {
			t.Fatal(err)
		}
		if got, want := img.RGBAAt(5, 100), g.th.Colors.Background; got.R != want.R || got.G != want.G || got.B != want.B || got.A != want.A {
			t.Fatalf("%s frame background=%v, active theme=%v", name, got, want)
		}
		colors := make(map[uint32]bool)
		for y := 0; y < 800; y += 5 {
			for x := 0; x < 1200; x += 5 {
				c := img.RGBAAt(x, y)
				colors[uint32(c.R)<<24|uint32(c.G)<<16|uint32(c.B)<<8|uint32(c.A)] = true
			}
		}
		if len(colors) < 20 {
			t.Fatalf("%s gallery has only %d colors; content missing", name, len(colors))
		}
		if dir := os.Getenv("GIO_TEST_ARTIFACTS"); dir != "" {
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatal(err)
			}
			file, err := os.Create(filepath.Join(dir, "gallery-"+name+".png"))
			if err != nil {
				t.Fatal(err)
			}
			encodeErr := png.Encode(file, img)
			closeErr := file.Close()
			if encodeErr != nil {
				t.Fatal(encodeErr)
			}
			if closeErr != nil {
				t.Fatal(closeErr)
			}
		}
		g.themeToggle.Click()
	}
}
