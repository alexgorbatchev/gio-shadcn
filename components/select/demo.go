package selectcomp

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	FruitSel *Select
	GenreSel *Select
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	return &DemoState{
		FruitSel: New(Config{
			Options: []*Item{
				NewItem("apple", "Apple"),
				NewItem("banana", "Banana"),
				NewItem("blueberry", "Blueberry"),
				NewItem("grapes", "Grapes"),
				NewItem("pineapple", "Pineapple"),
			},
			SelectedValue: "apple",
		}),
		GenreSel: New(Config{
			Options: []*Item{
				NewItem("house", "Progressive House"),
				NewItem("techno", "Techno"),
				NewItem("trance", "Trance"),
			},
			SelectedValue: "house",
		}),
	}
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("1. Select a Fruit (Official Demo)", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.FruitSel.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("2. Music Genre Dropdown", label.H4, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.GenreSel.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
