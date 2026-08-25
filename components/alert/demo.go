package alert

import (
	"gioui.org/layout"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DemoAction      *Alert
	DemoBasic       *Alert
	DemoColors      *Alert
	DemoStandard1   *Alert
	DemoStandard2   *Alert
	DemoDestructive *Alert
	DemoRTL1        *Alert
	DemoRTL2        *Alert
	BtnEnable       *button.Button
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	s.BtnEnable = button.New(button.Config{
		Text:    "Enable",
		Size:    theme.SizeSM,
		Variant: theme.VariantDefault,
	})

	// alert-action.tsx
	s.DemoAction = New(Config{
		Title:       "Dark mode is now available",
		Description: "Enable it under your profile settings to get started.",
		Action: func(gtx layout.Context) layout.Dimensions {
			return s.BtnEnable.Layout(gtx, theme.NewDark())
		},
	})

	// alert-basic.tsx
	s.DemoBasic = New(Config{
		Icon:        lucide.CircleCheck,
		Title:       "Account updated successfully",
		Description: "Your profile information has been saved. Changes will be reflected immediately.",
	})

	// alert-colors.tsx
	s.DemoColors = New(Config{
		Icon:        lucide.TriangleAlert,
		Title:       "Your subscription will expire in 3 days.",
		Description: "Renew now to avoid service interruption or upgrade to a paid plan to continue using the service.",
	})

	// alert-demo.tsx (2 items)
	s.DemoStandard1 = New(Config{
		Icon:        lucide.CircleCheck,
		Title:       "Payment successful",
		Description: "Your payment of $29.99 has been processed. A receipt has been sent to your email address.",
	})
	s.DemoStandard2 = New(Config{
		Icon:        lucide.Info,
		Title:       "New feature available",
		Description: "We've added dark mode support. You can enable it in your account settings.",
	})

	// alert-destructive.tsx
	s.DemoDestructive = New(Config{
		Variant:     theme.VariantDestructive,
		Icon:        lucide.CircleAlert,
		Title:       "Payment failed",
		Description: "Your payment could not be processed. Please check your payment method and try again.",
	})

	// alert-rtl.tsx (2 items)
	s.DemoRTL1 = New(Config{
		Icon:        lucide.CircleCheck,
		Title:       "تم الدفع بنجاح",
		Description: "تمت معالجة دفعتك البالغة 29.99 دولارًا. تم إرسال إيصال إلى عنوان بريدك الإلكتروني.",
	})
	s.DemoRTL2 = New(Config{
		Icon:        lucide.Info,
		Title:       "ميزة جديدة متاحة",
		Description: "لقد أضفنا دعم الوضع الداكن. يمكنك تفعيله في إعدادات حسابك.",
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (Action)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoAction.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (Basic)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBasic.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (Colors)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoColors.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (Demo)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoStandard1.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoStandard2.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (Destructive)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoDestructive.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Alert (RTL)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoRTL1.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoRTL2.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
