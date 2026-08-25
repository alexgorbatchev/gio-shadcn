package accordion

import (
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	DemoBasic     *Accordion
	DemoBorders   *Accordion
	DemoCard      *Accordion
	DemoStandard  *Accordion
	DemoDisabled  *Accordion
	DemoMultiple  *Accordion
	DemoRTL       *Accordion
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{}

	// accordion-basic.tsx
	s.DemoBasic = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("How do I reset my password?", "Click on 'Forgot Password' on the login page, enter your email address, and we'll send you a link to reset your password. The link will expire in 24 hours.", true),
			NewItem("Can I change my subscription plan?", "Yes, you can upgrade or downgrade your plan at any time from your account settings. Changes will be reflected in your next billing cycle.", false),
			NewItem("What payment methods do you accept?", "We accept all major credit cards, PayPal, and bank transfers. All payments are processed securely through our payment partners.", false),
		},
	})

	// accordion-borders.tsx
	s.DemoBorders = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("How does billing work?", "We offer monthly and annual subscription plans. Billing is charged at the beginning of each cycle, and you can cancel anytime. All plans include automatic backups, 24/7 support, and unlimited team members.", true),
			NewItem("Is my data secure?", "Yes. We use end-to-end encryption, SOC 2 Type II compliance, and regular third-party security audits. All data is encrypted at rest and in transit using industry-standard protocols.", false),
			NewItem("What integrations do you support?", "We integrate with 500+ popular tools including Slack, Zapier, Salesforce, HubSpot, and more. You can also build custom integrations using our REST API and webhooks.", false),
		},
	})

	// accordion-card.tsx
	s.DemoCard = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("What subscription plans do you offer?", "We offer three subscription tiers: Starter ($9/month), Professional ($29/month), and Enterprise ($99/month). Each plan includes increasing storage limits, API access, priority support, and team collaboration features.", true),
			NewItem("How does billing work?", "Billing occurs automatically at the start of each billing cycle. We accept all major credit cards, PayPal, and ACH transfers for enterprise customers. You'll receive an invoice via email after each payment.", false),
			NewItem("How do I cancel my subscription?", "You can cancel your subscription anytime from your account settings. There are no cancellation fees or penalties. Your access will continue until the end of your current billing period.", false),
		},
	})

	// accordion-demo.tsx
	s.DemoStandard = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("What are your shipping options?", "We offer standard (5-7 days), express (2-3 days), and overnight shipping. Free shipping on international orders.", true),
			NewItem("What is your return policy?", "Returns accepted within 30 days. Items must be unused and in original packaging. Refunds processed within 5-7 business days.", false),
			NewItem("How can I contact customer support?", "Reach us via email, live chat, or phone. We respond within 24 hours during business days.", false),
		},
	})

	// accordion-disabled.tsx
	s.DemoDisabled = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("Can I access my account history?", "Yes, you can view your complete account history including all transactions, plan changes, and support tickets in the Account History section of your dashboard.", false),
			NewItemConfig(ItemConfig{Title: "Premium feature information", Content: "This section contains information about premium features. Upgrade your plan to access this content.", Disabled: true}),
			NewItem("How do I update my email address?", "You can update your email address in your account settings. You'll receive a verification email at your new address to confirm the change.", false),
		},
	})

	// accordion-multiple.tsx
	s.DemoMultiple = New(Config{
		Type: TypeMultiple,
		Items: []*Item{
			NewItem("Notification Settings", "Manage how you receive notifications. You can enable email alerts for updates or push notifications for mobile devices.", true),
			NewItem("Privacy & Security", "Control your privacy settings and security preferences. Enable two-factor authentication, manage connected devices, review active sessions, and configure data sharing preferences. You can also download your data or delete your account.", false),
			NewItem("Billing & Subscription", "View your current plan, payment history, and upcoming invoices. Update your payment method, change your subscription tier, or cancel your subscription.", false),
		},
	})

	// accordion-rtl.tsx
	s.DemoRTL = New(Config{
		Type: TypeSingle,
		Items: []*Item{
			NewItem("كيف يمكنني إعادة تعيين كلمة المرور؟", "انقر على 'نسيت كلمة المرور' في صفحة تسجيل الدخول، أدخل عنوان بريدك الإلكتروني، وسنرسل لك رابطًا لإعادة تعيين كلمة المرور.", true),
			NewItem("هل يمكنني تغيير خطة الاشتراك الخاصة بي؟", "نعم، يمكنك ترقية أو تخفيض خطتك في أي وقت من إعدادات حسابك. ستظهر التغييرات في دورة الفوترة التالية.", false),
			NewItem("ما هي طرق الدفع التي تقبلونها؟", "نقبل جميع بطاقات الائتمان الرئيسية و PayPal والتحويلات المصرفية.", false),
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Basic)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBasic.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Borders)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoBorders.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Card)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoCard.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Demo)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoStandard.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Disabled)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoDisabled.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (Multiple)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoMultiple.Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return label.NewTypography("Accordion (RTL)", label.H4, "").Layout(gtx, th) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return s.DemoRTL.Layout(gtx, th) }),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
