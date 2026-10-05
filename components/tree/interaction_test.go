package tree

import (
	"image"
	"testing"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"github.com/bnema/gio-shadcn/internal/testui"
	"github.com/bnema/gio-shadcn/theme"
)

func treeHarness(tr *Tree) *testui.Harness {
	th := theme.New()
	return &testui.Harness{Size: image.Pt(300, 300), Widget: func(gtx layout.Context) layout.Dimensions { return tr.Layout(gtx, th) }}
}

func TestPointerSelectionAndExpansion(t *testing.T) {
	child := NewNode(NodeConfig{Label: "child"})
	folder := NewNode(NodeConfig{Label: "folder", Children: []*Node{child}})
	selected := 0
	tr := New(Config{Nodes: []*Node{folder}, OnSelect: func(n *Node) {
		selected++
		if n != child && n != folder {
			t.Errorf("selected %q", n.Label)
		}
	}})
	h := treeHarness(tr)
	h.Frame()
	h.Click(12, 12)
	if !folder.Expanded || len(tr.flatVisibleNodes) != 2 {
		t.Fatal("chevron did not expose child")
	}
	h.Click(100, float32(child.cachedY+10))
	if !child.Selected || folder.Selected || selected != 2 {
		t.Fatalf("selection child=%v folder=%v callbacks=%d", child.Selected, folder.Selected, selected)
	}
	child.Disabled = true
	tr.SelectNode(child)
	tr.SelectNode(nil)
	if selected != 2 {
		t.Fatal("disabled/nil selection invoked callback")
	}
	h.Frame()
}

func TestPointerDragReordersAndCancelPreservesOrder(t *testing.T) {
	for _, cancel := range []bool{false, true} {
		t.Run(map[bool]string{false: "release", true: "cancel"}[cancel], func(t *testing.T) {
			a, b := NewNode(NodeConfig{Label: "A", Selected: true}), NewNode(NodeConfig{Label: "B"})
			moves := 0
			tr := New(Config{Nodes: []*Node{a, b}, OnMove: func(source, parent *Node, index int) {
				moves++
				if source != a || parent != nil || index != 1 {
					t.Errorf("wrong move destination: %p %p %d", source, parent, index)
				}
			}})
			h := treeHarness(tr)
			h.Frame()
			h.Pointer(pointer.Press, 100, 10, pointer.ButtonPrimary)
			h.Pointer(pointer.Move, 100, float32(b.cachedY+b.cachedH-2), pointer.ButtonPrimary)
			if !tr.Session.WasDragging || tr.Session.DropTargetNode != b {
				t.Fatal("native drag did not find target")
			}
			kind := pointer.Release
			if cancel {
				kind = pointer.Cancel
			}
			h.Pointer(kind, 100, 62, 0)
			h.Frame()
			if cancel {
				if tr.Nodes[0] != a || moves != 0 {
					t.Fatal("cancel committed a move")
				}
			} else if tr.Nodes[1] != a || moves != 1 {
				t.Fatal("release did not move node")
			}
			if tr.Session.DraggedNode != nil || tr.Session.WasDragging {
				t.Fatal("drag session did not clear")
			}
		})
	}
}

func TestInvalidDropDoesNotDetachSource(t *testing.T) {
	for _, pos := range []DropPosition{DropNone, DropPosition(99), DropInside} {
		a, b := NewNode(NodeConfig{Label: "A"}), NewNode(NodeConfig{Label: "B"})
		tr := New(Config{Nodes: []*Node{a, b}})
		target := b
		if pos == DropInside {
			target = NewNode(NodeConfig{Label: "absent"})
		}
		tr.MoveNode(a, target, pos)
		if len(tr.Nodes) != 2 || tr.Nodes[0] != a || len(target.Children) != 0 {
			t.Fatalf("invalid drop %d changed tree", pos)
		}
	}
}

func TestNestedMovesAndDestinationCallbacks(t *testing.T) {
	a, b, c := NewNode(NodeConfig{Label: "A"}), NewNode(NodeConfig{Label: "B"}), NewNode(NodeConfig{Label: "C"})
	folder := NewNode(NodeConfig{Children: []*Node{a, b}})
	session := NewDragSession()
	source := New(Config{Nodes: []*Node{folder, c}, Session: session})
	target := New(Config{Nodes: []*Node{NewNode(NodeConfig{Selected: true})}, Session: session})
	source.MoveNode(c, b, DropBefore)
	if len(folder.Children) != 3 || folder.Children[1] != c {
		t.Fatal("nested insertion failed")
	}
	source.MoveNode(a, b, DropAfter)
	if folder.Children[2] != a {
		t.Fatal("nested removal/reordering failed")
	}
	moves := 0
	target.OnMove = func(n, p *Node, i int) {
		moves++
		if n != a || p != nil || i != 1 {
			t.Errorf("destination %p %p %d", n, p, i)
		}
	}
	session.MoveNodeCrossTree(source, target, a, target.Nodes[0], DropAfter)
	if len(folder.Children) != 2 || target.Nodes[1] != a || moves != 1 {
		t.Fatal("cross-tree nested move failed")
	}
	session.MoveNodeCrossTree(source, target, NewNode(NodeConfig{}), target.Nodes[0], DropBefore)
	if len(target.Nodes) != 2 {
		t.Fatal("absent source inserted")
	}
}

func TestDropHitZonesAndIndicators(t *testing.T) {
	a := NewNode(NodeConfig{Label: "A", Selected: true})
	folder := NewNode(NodeConfig{Label: "folder", Droppable: true})
	tr := New(Config{Nodes: []*Node{a, folder}})
	h := treeHarness(tr)
	h.Frame()
	h.Frame()
	for _, tc := range []struct {
		y    float32
		node *Node
		pos  DropPosition
	}{{-1, a, DropBefore}, {float32(a.cachedH - 1), a, DropAfter}, {float32(folder.cachedY + 1), folder, DropBefore}, {float32(folder.cachedY + folder.cachedH/2), folder, DropInside}, {float32(folder.cachedY + folder.cachedH - 1), folder, DropAfter}, {400, folder, DropAfter}} {
		n, p := tr.ResolveDropTarget(a, tc.y)
		if n != tc.node || p != tc.pos {
			t.Fatalf("y=%v target=%p pos=%d", tc.y, n, p)
		}
		tr.Session.DraggedNode = a
		tr.Session.SourceTree = tr
		tr.Session.DropTargetTree = tr
		tr.Session.DropTargetNode = tc.node
		tr.Session.DropPosition = tc.pos
		tr.Session.WasDragging = true
		if h.Frame().Size.Y <= 0 {
			t.Fatal("drag feedback disappeared")
		}
	}
	if n, p := tr.ResolveDropTarget(nil, 0); n != nil || p != DropNone {
		t.Fatal("nil drag resolved")
	}
}
