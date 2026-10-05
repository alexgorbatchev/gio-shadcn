/*
Package inputotp provides a PIN / verification code input component for gio-shadcn applications.

InputOTP displays single-digit input boxes for verification codes following
shadcn/ui design principles.
*/
package inputotp

import (
	"image"
	"strings"

	"gioui.org/font"
	"gioui.org/io/key"
	"gioui.org/layout"
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

	editors   []*widget.Editor
	lastValue string
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
	otp := &InputOTP{
		Length:   lenVal,
		Value:    config.Value,
		Classes:  config.Classes,
		OnChange: config.OnChange,
	}
	otp.syncValue()
	return otp
}

func (otp *InputOTP) syncValue() {
	if otp.Length <= 0 {
		otp.Length = 6
	}
	for len(otp.editors) < otp.Length {
		otp.editors = append(otp.editors, &widget.Editor{SingleLine: true, Submit: true, Filter: "0123456789"})
	}
	otp.editors = otp.editors[:otp.Length]
	value := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, otp.Value)
	value = value[:min(len(value), otp.Length)]
	for i, ed := range otp.editors {
		// Retain a whole paste until Update can distribute it across slots.
		ed.MaxLen = otp.Length + 1
		if i < len(value) {
			ed.SetText(value[i : i+1])
		} else {
			ed.SetText("")
		}
	}
	otp.Value, otp.lastValue = value, value
}

// Layout renders the array of PIN input boxes.
func (otp *InputOTP) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}
	if otp.Value != otp.lastValue || len(otp.editors) != otp.Length {
		otp.syncValue()
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
				text := ed.Text()
				if len(text) > 1 {
					for j := idx; j < min(otp.Length, idx+len(text)); j++ {
						otp.editors[j].SetText(text[j-idx : j-idx+1])
					}
					ed.SetText(text[:1])
				}
				if text != "" {
					next := min(otp.Length-1, idx+len(text))
					gtx.Execute(key.FocusCmd{Tag: otp.editors[next]})
					otp.editors[next].SetCaret(0, otp.editors[next].Len())
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
	var value strings.Builder
	for _, ed := range otp.editors {
		value.WriteString(ed.Text())
	}
	otp.Value = value.String()
	if otp.Value != otp.lastValue {
		otp.lastValue = otp.Value
		if otp.OnChange != nil {
			otp.OnChange(otp.Value)
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

	rr := theme.RRect(rect, radius)
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
