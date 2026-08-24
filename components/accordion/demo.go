package accordion

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	SingleAcc       *Accordion
	MultiAcc        *Accordion
	DisabledAcc     *Accordion
	ChevronAcc      *Accordion
	CustomHeaderAcc *Accordion
	BorderlessAcc   *Accordion
	NestedAcc       *Accordion
	ControlledAcc   *Accordion
	BtnExpandAll    *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.SingleAcc = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("Is it accessible?", "Yes. It adheres to the WAI-ARIA design pattern.", true),
			NewItem("Is it styled?", "Yes. It comes with default styles that match the other components' aesthetic.", false),
			NewItem("Is it animated?", "Yes. It's animated by default, but you can disable it if you prefer.", false),
		},
	})

	s.MultiAcc = New(Config{
		Type: TypeMultiple,
		Items: []*Item{
			NewItem("Audio Engine Specs", "Runs at 96kHz 24-bit floating point precision.", true),
			NewItem("GPU Vector Pipeline", "Gio immediate-mode engine renders vector paths directly on the GPU.", true),
			NewItem("Low Latency ASIO Buffer", "Hardware buffer configured for ultra-low latency playback.", false),
		},
	})

	s.DisabledAcc = New(Config{
		Items: []*Item{
			NewItemConfig(ItemConfig{Title: "Disabled Section (Locked)", Content: "Locked content panel.", Disabled: true}),
		},
	})

	s.ChevronAcc = New(Config{
		Items: []*Item{
			NewItemConfig(ItemConfig{Title: "Chevron Indicator Item", Content: "Uses custom chevron symbol.", Icon: "v", Expanded: true}),
		},
	})

	s.CustomHeaderAcc = New(Config{
		Items: []*Item{
			NewItemConfig(ItemConfig{Title: "Custom Header Badge Section", Content: "Custom header badge.", Expanded: true}),
		},
	})

	s.BorderlessAcc = New(Config{
		Borderless: true,
		Items: []*Item{
			NewItem("Borderless Item 1", "Flush borderless container style.", true),
		},
	})

	s.NestedAcc = New(Config{
		Items: []*Item{
			NewItemConfig(ItemConfig{
				Title:    "Parent Section (Contains Inner Accordion)",
				Expanded: true,
				ContentWidget: func(gtx layout.Context) layout.Dimensions {
					return s.SingleAcc.Layout(gtx, theme.NewDark())
				},
			}),
		},
	})

	s.ControlledAcc = New(Config{
		Type: TypeMultiple,
		Items: []*Item{
			NewItem("Controlled Panel 1", "State toggled externally by buttons.", true),
		},
	})

	s.BtnExpandAll = button.New(button.Config{
		Text:    "Expand All",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			for _, item := range s.ControlledAcc.Items {
				item.Expanded = true
			}
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.SingleAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.MultiAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DisabledAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.ChevronAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.CustomHeaderAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BorderlessAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.NestedAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.BtnExpandAll.Layout(gtx, th) }),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx) }),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions { return s.ControlledAcc.Layout(gtx, th) }),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
