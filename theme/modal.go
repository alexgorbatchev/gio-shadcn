package theme

import (
	"slices"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
)

// Modal retains the focus lifecycle of a window-level modal. Its zero value is ready to use.
type Modal struct {
	// ReturnFocus overrides the return target for a custom trigger or an
	// externally opened modal. Default component triggers configure this automatically.
	ReturnFocus func(layout.Context)
	open        *bool
	active      bool
	redraw      bool
}

// Invalidate requests another animation frame, even when the modal's input is
// disabled during dismissal. AddModal forwards it through the root input source.
func (m *Modal) Invalidate() {
	m.redraw = true
}

// LayoutRoot lays out application content and then its overlays. While a modal is
// open, the background receives a disabled native input source, so its controls
// are excluded from Gio's focus traversal as well as pointer and keyboard input.
func (t *Theme) LayoutRoot(gtx layout.Context, content layout.Widget) layout.Dimensions {
	defer func() { paint.ColorOp{Color: t.Colors.Background}.Add(gtx.Ops) }()
	background := gtx
	if t.topModal() != nil {
		background = background.Disabled()
	}
	dims := content(background)
	t.RenderOverlays(gtx)
	return dims
}

// AddModal queues a modal with native initial focus, Escape dismissal and focus
// restoration. Call LayoutRoot at the application root to contain focus, including
// focusable controls supplied by application code. Closing animations can still
// render when open is false, but receive a disabled input source.
func (t *Theme) AddModal(m *Modal, open *bool, initialFocus func(layout.Context), dismiss func(), content layout.Widget) {
	m.open = open
	if t.seenModals == nil {
		t.seenModals = make(map[*Modal]bool)
	}
	t.seenModals[m] = true
	if *open && !slices.Contains(t.modals, m) {
		t.modals = append(t.modals, m)
	}
	t.AddOverlay(func(gtx layout.Context) layout.Dimensions {
		root := gtx
		m.redraw = false
		if !*open || t.topModal() != m {
			gtx = gtx.Disabled()
		}
		if gtx.Enabled() && *open {
			if !m.active {
				m.active = true
				if initialFocus != nil {
					initialFocus(gtx)
				}
			}
			for {
				e, ok := gtx.Event(key.Filter{Name: key.NameEscape})
				if !ok {
					break
				}
				if e.(key.Event).State == key.Press {
					*open = false
					if dismiss != nil {
						dismiss()
					}
					gtx = gtx.Disabled()
					break
				}
			}
		}
		dims := content(gtx)
		if m.redraw {
			root.Execute(op.InvalidateCmd{})
		}
		return dims
	})
}

func (t *Theme) topModal() *Modal {
	for i := len(t.modals) - 1; i >= 0; i-- {
		if m := t.modals[i]; m.open != nil && *m.open {
			return m
		}
	}
	return nil
}

func (t *Theme) finishModals(gtx layout.Context) {
	active := t.modals[:0]
	for _, m := range t.modals {
		if t.seenModals[m] && *m.open {
			active = append(active, m)
			continue
		}
		if m.active {
			m.active = false
			t.restoreFocus = m.ReturnFocus
			// Background filters must be registered again before restoring focus.
			gtx.Execute(op.InvalidateCmd{})
		}
	}
	t.modals = active
	clear(t.seenModals)
}
