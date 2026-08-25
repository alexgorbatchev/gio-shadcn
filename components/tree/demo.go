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
	SharedSession  *DragSession
	WorkspaceTree  *Tree
	LibraryTree    *Tree
	SelectedPath   string
	StatusBadge    *badge.Badge

	// Action buttons
	btnNewFileSrc    widget.Clickable
	btnNewFolderSrc  widget.Clickable
	btnInfoSrc       widget.Clickable
	btnNewFileComp   widget.Clickable
	btnNewFolderComp widget.Clickable
	btnInfoTree      widget.Clickable
	btnNewAsset      widget.Clickable
	btnNewDoc        widget.Clickable
}

var defaultDemo = NewDemoState()

func NewDemoState() *DemoState {
	session := NewDragSession()

	s := &DemoState{
		SharedSession: session,
		SelectedPath:  "src/components/Tree.tsx",
		StatusBadge:   badge.New(badge.Config{Text: "Selected: src/components/Tree.tsx", Variant: theme.VariantSecondary}),
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

	// Tree A: Workspace Project Tree
	btnNode := NewNode(NodeConfig{ID: "btn", Label: "Button.tsx", Icon: lucide.FileCode})
	cardNode := NewNode(NodeConfig{ID: "card", Label: "Card.tsx", Icon: lucide.FileCode})
	dialogNode := NewNode(NodeConfig{ID: "dialog", Label: "Dialog.tsx", Icon: lucide.FileCode})
	treeNode := NewNode(NodeConfig{
		ID:       "tree",
		Label:    "Tree.tsx",
		Icon:     lucide.FileCode,
		Selected: true,
		Actions: []layout.Widget{
			makeAction(&s.btnInfoTree, lucide.Info, "Action: Info for Tree.tsx (16.8 KB)"),
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
			makeAction(&s.btnNewFileComp, lucide.FilePlus, "Action: Create new file in components/"),
			makeAction(&s.btnNewFolderComp, lucide.FolderPlus, "Action: Create subfolder in components/"),
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
			makeAction(&s.btnNewFolderSrc, lucide.FolderPlus, "Action: Create directory in src/"),
			makeAction(&s.btnInfoSrc, lucide.Info, "Action: Inspect folder src/"),
		},
	})

	pkgJson := NewNode(NodeConfig{ID: "pkg", Label: "package.json", Icon: lucide.FileCode})
	tsConfig := NewNode(NodeConfig{ID: "ts", Label: "tsconfig.json", Icon: lucide.FileCode})

	s.WorkspaceTree = New(Config{
		Nodes:   []*Node{srcFolder, pkgJson, tsConfig},
		Session: session,
		OnSelect: func(node *Node) {
			s.SelectedPath = node.Label
			s.StatusBadge.Text = fmt.Sprintf("Workspace: %s", node.Label)
		},
	})

	// Tree B: Shared Library & Assets Tree
	logoSvg := NewNode(NodeConfig{ID: "logo", Label: "logo.svg", Icon: lucide.Image})
	favicon := NewNode(NodeConfig{ID: "fav", Label: "favicon.ico", Icon: lucide.Image})
	themeJson := NewNode(NodeConfig{ID: "th_json", Label: "theme.json", Icon: lucide.FileCode})

	iconsFolder := NewNode(NodeConfig{
		ID:       "icons_folder",
		Label:    "icons",
		Expanded: true,
		Children: []*Node{logoSvg, favicon},
	})

	audioLoop := NewNode(NodeConfig{ID: "loop", Label: "deck_loop.wav", Icon: lucide.Music})
	audioFlac := NewNode(NodeConfig{ID: "flac", Label: "starlight.flac", Icon: lucide.Music})

	audioFolder := NewNode(NodeConfig{
		ID:       "audio_folder",
		Label:    "audio",
		Expanded: true,
		Children: []*Node{audioLoop, audioFlac},
	})

	assetsFolder := NewNode(NodeConfig{
		ID:       "assets_folder",
		Label:    "assets",
		Expanded: true,
		Children: []*Node{iconsFolder, audioFolder, themeJson},
		Actions: []layout.Widget{
			makeAction(&s.btnNewAsset, lucide.FolderPlus, "Action: New folder in assets/"),
		},
	})

	readmeDoc := NewNode(NodeConfig{ID: "readme_doc", Label: "README.md", Icon: lucide.FileText})
	changelogDoc := NewNode(NodeConfig{ID: "changelog", Label: "CHANGELOG.md", Icon: lucide.FileText})
	licenseDoc := NewNode(NodeConfig{ID: "license", Label: "LICENSE", Icon: lucide.FileText})

	docsFolder := NewNode(NodeConfig{
		ID:       "docs_folder",
		Label:    "docs",
		Expanded: false,
		Children: []*Node{readmeDoc, changelogDoc},
		Actions: []layout.Widget{
			makeAction(&s.btnNewDoc, lucide.FilePlus, "Action: New doc in docs/"),
		},
	})

	s.LibraryTree = New(Config{
		Nodes:   []*Node{assetsFolder, docsFolder, licenseDoc},
		Session: session,
		OnSelect: func(node *Node) {
			s.SelectedPath = node.Label
			s.StatusBadge.Text = fmt.Sprintf("Shared Library: %s", node.Label)
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
			return label.NewTypography("Dual-Tree Split View (Cross-Tree Drag & Drop)", label.H3, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return label.NewTypography("Drag files and folders across trees to move items between your Workspace and Shared Assets library.", label.Muted, "").Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space3}.Layout(gtx) }),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.StatusBadge.Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx) }),

		// 50/50 Horizontal Split Container
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Start}.Layout(gtx,
				// Left Panel (50%): Workspace Tree
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return label.NewTypography("PRIMARY WORKSPACE", label.Small, "").Layout(gtx, th)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return s.WorkspaceTree.Layout(gtx, th)
						}),
					)
				}),

				// Center Divider Spacer
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Width: th.Spacing.Space6}.Layout(gtx)
				}),

				// Right Panel (50%): Shared Library Tree
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return label.NewTypography("SHARED ASSET LIBRARY", label.Small, "").Layout(gtx, th)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx) }),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return s.LibraryTree.Layout(gtx, th)
						}),
					)
				}),
			)
		}),
	)
}

func Demo(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return defaultDemo.Layout(gtx, th)
}
