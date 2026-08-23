package accordion

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/theme"
)

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	singleAcc := New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("Section 1 (Single Open)", "Content panel for section 1.", true),
			NewItem("Section 2", "Content panel for section 2.", false),
		},
	})

	multiAcc := New(Config{
		Type: TypeMultiple,
		Items: []*Item{
			NewItem("Multi Section 1", "Content 1.", true),
			NewItem("Multi Section 2", "Content 2.", true),
		},
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return singleAcc.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return multiAcc.Layout(gtx, th) }),
	)
}
