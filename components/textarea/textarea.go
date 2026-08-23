/*
Package textarea provides a multiline text area input component for gio-shadcn applications.

TextAreas allow users to enter longer text descriptions following
shadcn/ui design principles.
*/
package textarea

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// TextArea represents a multiline text field input component.
type TextArea struct {
	editor *widget.Editor

	Text        string
	Placeholder string
	Height      unit.Dp
	Classes     string
	OnChange    func(string)
}

// Config represents configuration for creating a TextArea.
type Config struct {
	Text        string
	Placeholder string
	Height      unit.Dp
	Classes     string
	OnChange    func(string)
}

// New creates a new TextArea input component.
func New(config Config) *TextArea {
	ed := new(widget.Editor)
	ed.SingleLine = false
	if config.Text != "" {
		ed.SetText(config.Text)
	}

	h := config.Height
	if h <= 0 {
		h = unit.Dp(80)
	}

	return &TextArea{
		editor:      ed,
		Text:        config.Text,
		Placeholder: config.Placeholder,
		Height:      h,
		Classes:     config.Classes,
		OnChange:    config.OnChange,
	}
}

// Layout renders the multiline text editor.
func (ta *TextArea) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	// Process editor text changes
	for {
		ev, ok := ta.editor.Update(gtx)
		if !ok {
			break
		}
		if _, ok := ev.(widget.ChangeEvent); ok {
			ta.Text = ta.editor.Text()
			if ta.OnChange != nil {
				ta.OnChange(ta.Text)
			}
		}
	}

	bgColor := th.Colors.Background
	borderColor := th.Colors.Input

	styles := utils.ParseClasses(ta.Classes)
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	hPx := gtx.Dp(ta.Height)
	gtx.Constraints.Min.Y = hPx

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space3,
		Right:  th.Spacing.Space3,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderEditor := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.Editor(mTheme, ta.editor, ta.Placeholder).Layout(gtx)
		})
	}

	contentDims := renderEditor(gtxContent)
	finalHeight := contentDims.Size.Y
	if finalHeight < hPx {
		finalHeight = hPx
	}

	areaSize := image.Pt(contentDims.Size.X, finalHeight)

	dims := layout.Stack{}.Layout(gtx,
		// Background & Border drawn FIRST
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := image.Rectangle{Max: areaSize}
			radius := gtx.Dp(th.Radius.RadiusMD)

			theme.DrawRRectBackground(gtx, rect, radius, bgColor)

			rr := clip.UniformRRect(rect, radius)
			theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.0, borderColor)

			return layout.Dimensions{Size: areaSize}
		}),

		// Text Editor drawn ON TOP of background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return renderEditor(gtx)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}
