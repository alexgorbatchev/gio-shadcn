package tree_test

import (
	"image"
	"testing"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/components/tree"
	"github.com/bnema/gio-shadcn/theme"
)

func TestTreeLayout(t *testing.T) {
	th := theme.NewDark()
	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{
			tree.NewNode(tree.NodeConfig{
				ID:       "src",
				Label:    "src",
				Expanded: true,
				Children: []*tree.Node{
					tree.NewNode(tree.NodeConfig{ID: "app", Label: "App.tsx", Icon: lucide.FileCode}),
				},
			}),
			tree.NewNode(tree.NodeConfig{ID: "readme", Label: "README.md", Icon: lucide.FileText}),
		},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	dims := tr.Layout(gtx, th)
	if dims.Size.X <= 0 || dims.Size.Y <= 0 {
		t.Errorf("invalid dimensions returned from Tree.Layout")
	}
}

func TestTreeNodeExpandCollapse(t *testing.T) {
	node := tree.NewNode(tree.NodeConfig{
		ID:       "folder",
		Label:    "components",
		Expanded: false,
		Children: []*tree.Node{
			tree.NewNode(tree.NodeConfig{ID: "btn", Label: "Button.tsx"}),
		},
	})

	if node.Expanded {
		t.Fatalf("expected node to start collapsed")
	}
	node.Expanded = true
	if !node.Expanded {
		t.Fatalf("expected node to be expanded")
	}
}

func TestTreeNodeSelection(t *testing.T) {
	th := theme.NewDark()
	selectedLabel := ""
	nodeA := tree.NewNode(tree.NodeConfig{ID: "a", Label: "FileA.ts", Selected: false})
	nodeB := tree.NewNode(tree.NodeConfig{ID: "b", Label: "FileB.ts", Selected: false})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{nodeA, nodeB},
		OnSelect: func(n *tree.Node) {
			selectedLabel = n.Label
		},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	_ = tr.Layout(gtx, th)

	// Selecting nodeB
	tr.SelectNode(nodeB)
	if !nodeB.Selected || nodeA.Selected {
		t.Errorf("expected nodeB to be selected and nodeA unselected")
	}
	if selectedLabel != "FileB.ts" {
		t.Errorf("expected OnSelect to receive FileB.ts, got %s", selectedLabel)
	}
}

func TestTreeCustomSubicons(t *testing.T) {
	node := tree.NewNode(tree.NodeConfig{
		ID:    "doc",
		Label: "Document.pdf",
		Icon:  lucide.FileText,
	})

	if node.Icon != lucide.FileText {
		t.Errorf("expected custom subicon to be lucide.FileText")
	}
}

func TestTreeMoveNodeDropBeforeAndAfter(t *testing.T) {
	nodeA := tree.NewNode(tree.NodeConfig{ID: "a", Label: "A"})
	nodeB := tree.NewNode(tree.NodeConfig{ID: "b", Label: "B"})
	nodeC := tree.NewNode(tree.NodeConfig{ID: "c", Label: "C", Selected: true})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{nodeA, nodeB, nodeC},
	})

	// Move C before A
	tr.MoveNode(nodeC, nodeA, tree.DropBefore)
	if tr.Nodes[0] != nodeC || tr.Nodes[1] != nodeA || tr.Nodes[2] != nodeB {
		t.Errorf("unexpected order after DropBefore: %v, %v, %v", tr.Nodes[0].Label, tr.Nodes[1].Label, tr.Nodes[2].Label)
	}

	// Move C after B
	tr.MoveNode(nodeC, nodeB, tree.DropAfter)
	if tr.Nodes[2] != nodeC {
		t.Errorf("expected node C to be at index 2 after DropAfter")
	}
}

func TestTreeMoveNodeInsideFolder(t *testing.T) {
	fileNode := tree.NewNode(tree.NodeConfig{ID: "f", Label: "File.go", Selected: true})
	folderNode := tree.NewNode(tree.NodeConfig{
		ID:       "dir",
		Label:    "pkg",
		Expanded: false,
	})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{fileNode, folderNode},
	})

	// Move file into folder
	tr.MoveNode(fileNode, folderNode, tree.DropInside)
	if len(folderNode.Children) != 1 || folderNode.Children[0] != fileNode {
		t.Fatalf("expected fileNode to be child of folderNode")
	}
	if !folderNode.Expanded {
		t.Errorf("expected folder to be expanded after dropping inside")
	}
}

func TestTreeMoveNodeCircularDropPrevention(t *testing.T) {
	child := tree.NewNode(tree.NodeConfig{ID: "child", Label: "Child"})
	parent := tree.NewNode(tree.NodeConfig{
		ID:       "parent",
		Label:    "Parent",
		Children: []*tree.Node{child},
	})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{parent},
	})

	// Dragging a parent into its own child MUST be safely rejected
	tr.MoveNode(parent, child, tree.DropInside)
	if len(tr.Nodes) != 1 || tr.Nodes[0] != parent {
		t.Errorf("expected parent to remain at root")
	}
}

func TestTreeHitTestingWithLocalOffset(t *testing.T) {
	th := theme.NewDark()
	node1 := tree.NewNode(tree.NodeConfig{ID: "1", Label: "Node 1", Selected: true})
	node2 := tree.NewNode(tree.NodeConfig{ID: "2", Label: "Node 2"})
	node3 := tree.NewNode(tree.NodeConfig{ID: "3", Label: "Node 3"})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{node1, node2, node3},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	_ = tr.Layout(gtx, th)

	target, pos := tr.ResolveDropTarget(node1, 16.0)
	if target != node1 {
		t.Errorf("expected target node1, got %v", target)
	}
	if pos != tree.DropInside && pos != tree.DropAfter && pos != tree.DropBefore {
		t.Errorf("expected valid drop position, got %v", pos)
	}

	target3, _ := tr.ResolveDropTarget(node1, 70.0)
	if target3 != node3 {
		t.Errorf("expected target node3 at Y=70, got %v", target3)
	}
}

func TestTreeDropInsideGenerousHitZone(t *testing.T) {
	th := theme.NewDark()
	fileNode := tree.NewNode(tree.NodeConfig{ID: "f", Label: "File.ts", Selected: true})
	folderNode := tree.NewNode(tree.NodeConfig{
		ID:       "folder",
		Label:    "components",
		Expanded: false,
		Droppable: true,
	})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{fileNode, folderNode},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 200))}
	_ = tr.Layout(gtx, th)

	h := float32(32.0)
	target, pos := tr.ResolveDropTarget(fileNode, h+h*0.30)
	if target != folderNode || pos != tree.DropInside {
		t.Errorf("expected DropInside on folder at 30%%, got target=%v, pos=%v", target, pos)
	}

	targetMid, posMid := tr.ResolveDropTarget(fileNode, h+h*0.50)
	if targetMid != folderNode || posMid != tree.DropInside {
		t.Errorf("expected DropInside on folder at 50%%, got target=%v, pos=%v", targetMid, posMid)
	}

	target70, pos70 := tr.ResolveDropTarget(fileNode, h+h*0.70)
	if target70 != folderNode || pos70 != tree.DropInside {
		t.Errorf("expected DropInside on folder at 70%%, got target=%v, pos=%v", target70, pos70)
	}
}

func TestTreeConfigurableIndent(t *testing.T) {
	// Default indent must be 9dp (50% reduction from old 18dp)
	trDefault := tree.New(tree.Config{})
	if trDefault.Indent != unit.Dp(9) {
		t.Errorf("expected default indent 9dp, got %v", trDefault.Indent)
	}

	// Custom configured indent
	trCustom := tree.New(tree.Config{
		Indent: unit.Dp(14),
	})
	if trCustom.Indent != unit.Dp(14) {
		t.Errorf("expected custom indent 14dp, got %v", trCustom.Indent)
	}
}

func TestTreeNodeTrailingActions(t *testing.T) {
	th := theme.NewDark()
	actionRendered := false

	node := tree.NewNode(tree.NodeConfig{
		ID:    "src",
		Label: "src",
		Actions: []layout.Widget{
			func(gtx layout.Context) layout.Dimensions {
				actionRendered = true
				return layout.Dimensions{Size: image.Pt(16, 16)}
			},
		},
	})

	tr := tree.New(tree.Config{
		Nodes: []*tree.Node{node},
	})

	gtx := layout.Context{Ops: new(op.Ops), Constraints: layout.Exact(image.Pt(300, 100))}
	_ = tr.Layout(gtx, th)

	if !actionRendered {
		t.Errorf("expected trailing action to be rendered")
	}
}
