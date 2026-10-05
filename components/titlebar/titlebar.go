/*
Package titlebar provides a custom window title bar component for gio-shadcn applications.
*/
package titlebar

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

type TitleBar struct {
	window *app.Window

	title        string
	variant      theme.Variant
	classes      string
	showControls bool

	decorations widget.Decorations
}

type Option func(*TitleBar)

func WithTitle(title string) Option {
	return func(tb *TitleBar) {
		tb.title = title
	}
}

func WithWindow(w *app.Window) Option {
	return func(tb *TitleBar) {
		tb.window = w
	}
}

func WithVariant(variant theme.Variant) Option {
	return func(tb *TitleBar) {
		tb.variant = variant
	}
}

func WithClasses(classes string) Option {
	return func(tb *TitleBar) {
		tb.classes = classes
	}
}

func WithControls(show bool) Option {
	return func(tb *TitleBar) {
		tb.showControls = show
	}
}

func NewTitleBar(options ...Option) *TitleBar {
	tb := &TitleBar{
		variant:      theme.VariantDefault,
		showControls: true,
	}

	for _, option := range options {
		option(tb)
	}

	return tb
}

func (tb *TitleBar) Layout(gtx layout.Context, th *theme.Theme, window *app.Window) layout.Dimensions {
	if tb.window == nil {
		tb.window = window
	}

	if th == nil {
		th = theme.New()
	}
	if actions := tb.decorations.Update(gtx); actions != 0 && tb.window != nil {
		tb.window.Perform(actions)
	}

	height := gtx.Dp(40)
	gtx.Constraints.Min.Y = height
	gtx.Constraints.Max.Y = height

	variantConfig := theme.GetTitleBarVariant(tb.variant, &th.Colors)
	if background := utils.ParseClasses(tb.classes).Background; background.A > 0 {
		variantConfig.Background = background
	}

	// Safely draw background and border with push/pop clips FIRST
	bgClip := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	paint.ColorOp{Color: variantConfig.Background}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	bgClip.Pop()

	if variantConfig.BorderWidth > 0 {
		borderHeight := int(variantConfig.BorderWidth)
		bRect := clip.Rect{
			Min: image.Pt(0, gtx.Constraints.Max.Y-borderHeight),
			Max: gtx.Constraints.Max,
		}
		bClip := bRect.Push(gtx.Ops)
		paint.ColorOp{Color: variantConfig.Border}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		bClip.Pop()
	}

	dims := layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min = gtx.Constraints.Max
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return tb.decorations.LayoutMove(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min = gtx.Constraints.Max
						return layout.Inset{Left: unit.Dp(16)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							mTheme := th.MaterialTheme
							if mTheme == nil {
								mTheme = material.NewTheme()
							}
							label := material.Label(mTheme, unit.Sp(14), tb.title)
							label.Color = variantConfig.Foreground
							label.Font.Weight = font.Medium
							return label.Layout(gtx)
						})
					})
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !tb.showControls {
						return layout.Dimensions{}
					}
					return tb.layoutWindowControls(gtx, th, variantConfig)
				}),
			)
		}),
	)

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

// SetMode synchronizes decoration state with app.ConfigEvent.Config.Mode.
func (tb *TitleBar) SetMode(mode app.WindowMode) {
	tb.decorations.Maximized = mode == app.Maximized || mode == app.Fullscreen
}

func (tb *TitleBar) layoutWindowControls(gtx layout.Context, th *theme.Theme, variantConfig theme.VariantConfig) layout.Dimensions {
	buttonWidth := gtx.Dp(46)
	buttonHeight := gtx.Dp(40)

	btnConstraints := layout.Exact(image.Point{X: buttonWidth, Y: buttonHeight})

	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = btnConstraints
			return tb.layoutControlItem(gtx, tb.decorations.Clickable(system.ActionMinimize), "−", variantConfig.Foreground)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = btnConstraints
			symbol := "□"
			if tb.decorations.Maximized {
				symbol = "❐"
			}
			return tb.layoutControlItem(gtx, tb.decorations.Clickable(system.ActionMaximize), symbol, variantConfig.Foreground)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints = btnConstraints
			return tb.layoutControlItem(gtx, tb.decorations.Clickable(system.ActionClose), "✕", variantConfig.Foreground)
		}),
	)
}

func (tb *TitleBar) layoutControlItem(gtx layout.Context, clickable *widget.Clickable, symbol string, fgColor color.NRGBA) layout.Dimensions {
	mTheme := material.NewTheme()
	label := material.Label(mTheme, unit.Sp(12), symbol)
	label.Color = fgColor

	return clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return label.Layout(gtx)
		})
	})
}
