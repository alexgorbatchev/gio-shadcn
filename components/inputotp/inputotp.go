/*
Package inputotp provides a PIN / verification code input component for gio-shadcn applications.

InputOTP displays single-digit input boxes for verification codes following
shadcn/ui design principles.
*/
package inputotp

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
)

// InputOTP represents a single-digit PIN box input array.
type InputOTP struct {
	Length   int
	Value    string
	Classes  string
	OnChange func(string)

	editors []*widget.Editor
}

// Config represents configuration for creating an InputOTP component.
type Config struct {
	Length   int
	Value    string
	Classes  string
	OnChange func(string)
}

// New creates a new InputOTP component.
func New(config Config) *InputOTP {
	lenVal := config.Length
	if lenVal <= 0 {
		lenVal = 6
	}
	editors := make([]*widget.Editor, lenVal)
	for i := 0; i < lenVal; i++ {
		ed := new(widget.Editor)
		ed.SingleLine = true
		ed.Submit = true
		editors[i] = ed
	}
	return &InputOTP{
		Length:   lenVal,
		Value:    config.Value,
		Classes:  config.Classes,
		OnChange: config.OnChange,
		editors:  editors,
	}
}

// Layout renders the array of PIN input boxes.
func (otp *InputOTP) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	children := make([]layout.FlexChild, 0, otp.Length*2)

	for i := 0; i < otp.Length; i++ {
		idx := i
		ed := otp.editors[idx]

		for {
			ev, ok := ed.Update(gtx)
			if !ok {
				break
			}
			if _, ok := ev.(widget.ChangeEvent); ok {
				var val string
				for _, e := range otp.editors {
					val += e.Text()
				}
				otp.Value = val
				if otp.OnChange != nil {
					otp.OnChange(otp.Value)
				}
			}
		}

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return otp.layoutBox(gtx, th, mTheme, ed)
		}))

		if i < otp.Length-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}))
		}
	}

	dims := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)

	// Reset active GPU paint color state
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (otp *InputOTP) layoutBox(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, ed *widget.Editor) layout.Dimensions {
	boxSize := gtx.Dp(th.Spacing.Space10)
	size := image.Pt(boxSize, boxSize)
	gtx.Constraints = layout.Exact(size)

	padding := layout.Inset{
		Top:    th.Spacing.Space2,
		Bottom: th.Spacing.Space2,
		Left:   th.Spacing.Space2,
		Right:  th.Spacing.Space2,
	}

	rect := image.Rectangle{Max: size}
	radius := gtx.Dp(th.Radius.RadiusMD)

	borderColor := th.Colors.Input
	if ed.Text() != "" {
		borderColor = th.Colors.Ring
	}

	theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Background)

	rr := clip.UniformRRect(rect, radius)
	theme.DrawStroke(gtx, rr.Path(gtx.Ops), 1.5, borderColor)

	_ = padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		edWidget := material.Editor(mTheme, ed, "")
		edWidget.Color = th.Colors.Foreground
		edWidget.Font.Weight = font.Bold
		edWidget.TextSize = th.Typography.FontSizeBase
		return edWidget.Layout(gtx)
	})

	return layout.Dimensions{Size: size}
}
