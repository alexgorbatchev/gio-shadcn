package tree

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/badge"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

type DemoState struct {
	FileSystemTree *Tree
	SelectedPath   string
	StatusBadge    *badge.Badge

	// Action buttons
	btnNewFileSrc    widget.Clickable
	btnNewFolderSrc  widget.Clickable
	btnInfoSrc       widget.Clickable
	btnNewFileComp   widget.Clickable
	btnNewFolderComp widget.Clickable
	btnInfoTree      widget.Clickable
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	s := &DemoState{
		SelectedPath: "src/components/Tree.tsx",
		StatusBadge:  badge.New(badge.Config{Text: "Selected: src/components/Tree.tsx", Variant: theme.VariantSecondary}),
	}

	makeAction := func(btn *widget.Clickable, icon *lucide.Icon, actionDesc string) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			for btn.Clicked(gtx) {
				s.StatusBadge.Text = actionDesc
			}
			return btn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				fg := theme.DarkColorScheme().MutedFg
				if btn.Hovered() {
					fg = theme.DarkColorScheme().Foreground
				}
				return icon.LayoutSize(gtx, unit.Dp(14), fg)
			})
		}
	}

	btnNode := NewNode(NodeConfig{ID: "btn", Label: "Button.tsx", Icon: lucide.FileCode})
	cardNode := NewNode(NodeConfig{ID: "card", Label: "Card.tsx", Icon: lucide.FileCode})
	dialogNode := NewNode(NodeConfig{ID: "dialog", Label: "Dialog.tsx", Icon: lucide.FileCode})
	treeNode := NewNode(NodeConfig{
		ID:       "tree",
		Label:    "Tree.tsx",
		Icon:     lucide.FileCode,
		Selected: true,
		Actions: []layout.Widget{
			makeAction(&s.btnInfoTree, lucide.Info, "Action: Info for Tree.tsx (16.8 KB, 420 lines)"),
		},
	})

	useAuthNode := NewNode(NodeConfig{ID: "auth", Label: "useAuth.ts", Icon: lucide.FileCode})
	useThemeNode := NewNode(NodeConfig{ID: "theme", Label: "useTheme.ts", Icon: lucide.FileCode})

	appNode := NewNode(NodeConfig{ID: "app", Label: "App.tsx", Icon: lucide.FileCode})
	mainNode := NewNode(NodeConfig{ID: "main", Label: "main.tsx", Icon: lucide.FileCode})

	compFolder := NewNode(NodeConfig{
		ID:       "comp_folder",
		Label:    "components",
		Expanded: true,
		Children: []*Node{btnNode, cardNode, dialogNode, treeNode},
		Actions: []layout.Widget{
			makeAction(&s.btnNewFileComp, lucide.FilePlus, "Action: Create new component file in src/components/"),
			makeAction(&s.btnNewFolderComp, lucide.FolderPlus, "Action: Create new subfolder in src/components/"),
		},
	})

	hooksFolder := NewNode(NodeConfig{
		ID:       "hooks_folder",
		Label:    "hooks",
		Expanded: false,
		Children: []*Node{useAuthNode, useThemeNode},
	})

	srcFolder := NewNode(NodeConfig{
		ID:       "src_folder",
		Label:    "src",
		Expanded: true,
		Children: []*Node{compFolder, hooksFolder, appNode, mainNode},
		Actions: []layout.Widget{
			makeAction(&s.btnNewFileSrc, lucide.FilePlus, "Action: Create new file in src/"),
			makeAction(&s.btnNewFolderSrc, lucide.FolderPlus, "Action: Create new directory in src/"),
			makeAction(&s.btnInfoSrc, lucide.Info, "Action: Inspect folder src/ (8 items)"),
		},
	})

	publicFolder := NewNode(NodeConfig{
		ID:       "public_folder",
		Label:    "public",
		Expanded: false,
		Children: []*Node{
			NewNode(NodeConfig{ID: "fav", Label: "favicon.ico", Icon: lucide.Image}),
			NewNode(NodeConfig{ID: "robots", Label: "robots.txt", Icon: lucide.FileText}),
		},
	})

	pkgJson := NewNode(NodeConfig{ID: "pkg", Label: "package.json", Icon: lucide.FileCode})
	tsConfig := NewNode(NodeConfig{ID: "ts", Label: "tsconfig.json", Icon: lucide.FileCode})
	readme := NewNode(NodeConfig{ID: "readme", Label: "README.md", Icon: lucide.FileText})

	s.FileSystemTree = New(Config{
		Nodes: []*Node{srcFolder, publicFolder, pkgJson, tsConfig, readme},
		OnSelect: func(node *Node) {
			s.SelectedPath = node.Label
			s.StatusBadge.Text = fmt.Sprintf("Selected: %s", node.Label)
		},
	})

	return s
}

func (s *DemoState) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("File System Tree View (Drag & Drop Reordering)", label.H3, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Click items to select, drag to reorder/nest, or use right-aligned action buttons (New File, New Folder, Info).", label.Muted, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.StatusBadge.Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.FileSystemTree.Layout(gtx, th)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
