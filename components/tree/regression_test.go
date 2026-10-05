package tree_test

import (
	"testing"

	"github.com/bnema/gio-shadcn/components/tree"
)

func TestTreeMoveCallbackReportsDestination(t *testing.T) {
	a := tree.NewNode(tree.NodeConfig{ID: "a", Label: "A"})
	b := tree.NewNode(tree.NodeConfig{ID: "b", Label: "B"})
	c := tree.NewNode(tree.NodeConfig{ID: "c", Label: "C"})
	var gotParent *tree.Node
	gotIndex := -1
	tr := tree.New(tree.Config{Nodes: []*tree.Node{a, b, c}, OnMove: func(source, parent *tree.Node, index int) { gotParent, gotIndex = parent, index }})
	tr.MoveNode(c, a, tree.DropBefore)
	if gotParent != nil || gotIndex != 0 {
		t.Fatalf("root move before A reports parent=%v, index=%d; want nil, 0", gotParent, gotIndex)
	}
}
