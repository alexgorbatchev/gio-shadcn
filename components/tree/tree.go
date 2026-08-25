/*
Package tree provides a hierarchical tree view component with drag-and-drop (DnD) support
for gio-shadcn applications.

Trees display nested hierarchical items (such as file systems, categories, or navigation paths)
with expandable chevrons, optional sub-icons (folders, files), interactive node selection,
right-aligned action elements, and intuitive drag-and-drop reordering within or across tree views.
*/
package tree

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/font"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alexgorbatchev/gio-lucide"
	"github.com/bnema/gio-shadcn/theme"
)

type DropPosition int

const (
	DropNone DropPosition = iota
	DropBefore
	DropInside
	DropAfter
)

// DragSession coordinates drag-and-drop operations across single or multiple Tree instances.
type DragSession struct {
	DraggedNode    *Node
	SourceTree     *Tree
	DragStart      f32.Point
	GlobalDragX    float32
	GlobalDragY    float32
	LocalDragX     float32
	LocalDragY     float32
	DropTargetNode *Node
	DropTargetTree *Tree
	DropPosition   DropPosition
	WasDragging    bool

	registeredTrees []*Tree
}

// NewDragSession creates a shared drag-and-drop coordinator.
func NewDragSession() *DragSession {
	return &DragSession{}
}

// RegisterTree registers a tree instance for cross-tree drag-and-drop resolution.
func (s *DragSession) RegisterTree(t *Tree) {
	for _, registered := range s.registeredTrees {
		if registered == t {
			return
		}
	}
	s.registeredTrees = append(s.registeredTrees, t)
}

// ResolveDropTargetAtGlobal resolves the target tree and target node from global coordinates.
func (s *DragSession) ResolveDropTargetAtGlobal(dragged *Node, globalX, globalY float32) (*Tree, *Node, DropPosition) {
	if dragged == nil {
		return nil, nil, DropNone
	}

	// In multi-tree layouts, compute cumulative horizontal offsets
	cumX := float32(0)
	for i, tr := range s.registeredTrees {
		width := float32(tr.cachedWidth)
		if width <= 0 {
			width = 280
		}
		tr.cachedOriginX = int(cumX)

		isLast := (i == len(s.registeredTrees)-1)
		if (globalX >= cumX && globalX < cumX+width) || (isLast && globalX >= cumX) {
			localTreeY := globalY
			targetNode, pos := tr.resolveDropTargetByTreeY(localTreeY)
			if targetNode != nil {
				return tr, targetNode, pos
			}
		}
		cumX += width
	}

	// Fallback
	if len(s.registeredTrees) > 0 {
		first := s.registeredTrees[0]
		targetNode, pos := first.resolveDropTargetByTreeY(globalY)
		return first, targetNode, pos
	}

	return nil, nil, DropNone
}

// MoveNodeCrossTree moves a node from sourceTree to targetTree relative to target node.
func (s *DragSession) MoveNodeCrossTree(sourceTree, targetTree *Tree, source, target *Node, pos DropPosition) {
	if sourceTree == nil || targetTree == nil || source == nil || target == nil || source == target {
		return
	}

	// Prevent dropping a parent into its own descendant
	if sourceTree == targetTree && sourceTree.isDescendant(source, target) {
		return
	}

	// 1. Remove source from sourceTree
	if !sourceTree.removeNode(nil, sourceTree.Nodes, source) {
		return
	}

	// 2. When moving across trees, if targetTree already has an active selection,
	// maintain the targetTree's original selection by clearing source.Selected.
	if sourceTree != targetTree {
		if targetTree.hasSelected(targetTree.Nodes) {
			source.Selected = false
		}
	}

	// 3. Insert into targetTree
	if pos == DropInside {
		target.Children = append(target.Children, source)
		target.Expanded = true
		target.isFolder = true
	} else {
		parent, idx := targetTree.findParentAndIndex(nil, targetTree.Nodes, target)
		if idx < 0 {
			targetTree.Nodes = append(targetTree.Nodes, source)
		} else {
			if pos == DropAfter {
				idx++
			}
			if parent != nil {
				if idx < 0 {
					idx = 0
				}
				if idx > len(parent.Children) {
					idx = len(parent.Children)
				}
				parent.Children = insertNode(parent.Children, idx, source)
			} else {
				if idx < 0 {
					idx = 0
				}
				if idx > len(targetTree.Nodes) {
					idx = len(targetTree.Nodes)
				}
				targetTree.Nodes = insertNode(targetTree.Nodes, idx, source)
			}
		}
	}

	if sourceTree.OnMove != nil {
		sourceTree.OnMove(source, target, int(pos))
	}
	if targetTree != sourceTree && targetTree.OnMove != nil {
		targetTree.OnMove(source, target, int(pos))
	}
}

// Node represents a single item or branch in the Tree hierarchy.
type Node struct {
	ID        string
	Label     string
	Icon      *lucide.Icon
	Expanded  bool
	Selected  bool
	Disabled  bool
	Draggable bool
	Droppable bool
	Actions   []layout.Widget
	Children  []*Node

	// Internal state
	clickable  widget.Clickable
	chevronBtn widget.Clickable
	drag       gesture.Drag
	cachedY    int
	cachedH    int
	depth      int
	isFolder   bool
}

// Config represents configuration for creating a Node.
type NodeConfig struct {
	ID        string
	Label     string
	Icon      *lucide.Icon
	Expanded  bool
	Selected  bool
	Disabled  bool
	Draggable bool
	Droppable bool
	Actions   []layout.Widget
	Children  []*Node
}

// NewNode creates a new tree node.
func NewNode(config NodeConfig) *Node {
	draggable := config.Draggable
	if !draggable {
		draggable = true
	}
	droppable := config.Droppable
	if len(config.Children) > 0 {
		droppable = true
	}

	return &Node{
		ID:        config.ID,
		Label:     config.Label,
		Icon:      config.Icon,
		Expanded:  config.Expanded,
		Selected:  config.Selected,
		Disabled:  config.Disabled,
		Draggable: draggable,
		Droppable: droppable,
		Actions:   config.Actions,
		Children:  config.Children,
	}
}

// Tree represents a hierarchical tree component.
type Tree struct {
	Nodes    []*Node
	Indent   unit.Dp
	Classes  string
	Session  *DragSession
	OnSelect func(node *Node)
	OnMove   func(source *Node, targetParent *Node, targetIndex int)

	cachedOriginX    int
	cachedOriginY    int
	cachedWidth      int
	flatVisibleNodes []*Node
}

// Config represents configuration for creating a Tree.
type Config struct {
	Nodes    []*Node
	Indent   unit.Dp
	Classes  string
	Session  *DragSession
	OnSelect func(node *Node)
	OnMove   func(source *Node, targetParent *Node, targetIndex int)
}

// New creates a new Tree component.
func New(config Config) *Tree {
	indent := config.Indent
	if indent <= 0 {
		indent = unit.Dp(9)
	}

	session := config.Session
	if session == nil {
		session = NewDragSession()
	}

	t := &Tree{
		Nodes:    config.Nodes,
		Indent:   indent,
		Classes:  config.Classes,
		Session:  session,
		OnSelect: config.OnSelect,
		OnMove:   config.OnMove,
	}

	session.RegisterTree(t)
	return t
}

// SelectNode selects the specified node and unselects all other nodes in the tree.
func (t *Tree) SelectNode(target *Node) {
	if target == nil || target.Disabled {
		return
	}
	t.unselectAll(t.Nodes)
	target.Selected = true
	if t.OnSelect != nil {
		t.OnSelect(target)
	}
}

func (t *Tree) unselectAll(list []*Node) {
	for _, n := range list {
		n.Selected = false
		if len(n.Children) > 0 {
			t.unselectAll(n.Children)
		}
	}
}

func (t *Tree) hasSelected(list []*Node) bool {
	for _, n := range list {
		if n.Selected {
			return true
		}
		if len(n.Children) > 0 && t.hasSelected(n.Children) {
			return true
		}
	}
	return false
}

// Layout renders the tree hierarchy, chevrons, sub-icons, and interactive DnD indicators.
func (t *Tree) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if th == nil {
		th = theme.New()
	}

	mTheme := th.MaterialTheme
	if mTheme == nil {
		mTheme = material.NewTheme()
	}

	// Register with session
	t.Session.RegisterTree(t)
	t.cachedWidth = gtx.Constraints.Max.X

	// 1. Flatten visible nodes for hit testing and linear rendering
	t.flatVisibleNodes = t.flatVisibleNodes[:0]
	t.flattenNodes(t.Nodes, 0)

	// 2. Process DnD gestures only on selected draggable nodes
	t.processDnD(gtx)

	// 3. Render flattened visible nodes with exact cumulative Y coordinate tracking
	children := make([]layout.FlexChild, 0, len(t.flatVisibleNodes))
	curY := 0

	for _, n := range t.flatVisibleNodes {
		node := n
		node.cachedY = curY

		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			dims := t.layoutNode(gtx, th, mTheme, node)
			node.cachedH = dims.Size.Y
			return dims
		}))

		h := node.cachedH
		if h <= 0 {
			h = gtx.Dp(unit.Dp(32))
		}
		curY += h
	}

	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)

	// 4. If dragging, render floating drag preview badge at cursor location (anchored to source tree)
	if t.Session.DraggedNode != nil && t.Session.WasDragging && t.Session.SourceTree == t {
		t.layoutDragGhost(gtx, th, mTheme)
	}

	// Reset active GPU paint color state back to background
	paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)

	return dims
}

func (t *Tree) flattenNodes(nodes []*Node, depth int) {
	for _, node := range nodes {
		node.depth = depth
		node.isFolder = len(node.Children) > 0 || node.Droppable
		t.flatVisibleNodes = append(t.flatVisibleNodes, node)

		if node.Expanded && len(node.Children) > 0 {
			t.flattenNodes(node.Children, depth+1)
		}
	}
}

func (t *Tree) processDnD(gtx layout.Context) {
	session := t.Session

	for _, node := range t.flatVisibleNodes {
		if !node.Selected || node.Disabled || !node.Draggable {
			continue
		}

		for {
			ev, ok := node.drag.Update(gtx.Metric, gtx.Source, gesture.Both)
			if !ok {
				break
			}
			switch ev.Kind {
			case pointer.Press:
				session.LocalDragX = ev.Position.X
				session.LocalDragY = ev.Position.Y
				session.GlobalDragX = ev.Position.X + float32(t.cachedOriginX)
				session.GlobalDragY = float32(node.cachedY+t.cachedOriginY) + ev.Position.Y
				session.DragStart = f32.Pt(session.GlobalDragX, session.GlobalDragY)
				session.DraggedNode = node
				session.SourceTree = t
				session.WasDragging = false
			case pointer.Drag:
				session.LocalDragX = ev.Position.X
				session.LocalDragY = ev.Position.Y
				session.GlobalDragX = ev.Position.X + float32(t.cachedOriginX)
				session.GlobalDragY = float32(node.cachedY+t.cachedOriginY) + ev.Position.Y

				dx := session.GlobalDragX - session.DragStart.X
				dy := session.GlobalDragY - session.DragStart.Y
				if dx*dx+dy*dy >= 9 {
					session.DraggedNode = node
					session.SourceTree = t
					session.WasDragging = true
					targetTree, targetNode, pos := session.ResolveDropTargetAtGlobal(node, session.GlobalDragX, session.GlobalDragY)
					session.DropTargetTree = targetTree
					session.DropTargetNode = targetNode
					session.DropPosition = pos
				}
			case pointer.Release, pointer.Cancel:
				if session.WasDragging && session.DraggedNode != nil && session.DropTargetTree != nil && session.DropTargetNode != nil && session.DraggedNode != session.DropTargetNode {
					session.MoveNodeCrossTree(session.SourceTree, session.DropTargetTree, session.DraggedNode, session.DropTargetNode, session.DropPosition)
				}
				session.DraggedNode = nil
				session.DropTargetTree = nil
				session.DropTargetNode = nil
				session.DropPosition = DropNone
				session.WasDragging = false
			}
		}
	}
}

// ResolveDropTarget resolves the drop target node and drop position for a dragged node at a local Y offset.
func (t *Tree) ResolveDropTarget(dragged *Node, localY float32) (*Node, DropPosition) {
	if dragged == nil {
		return nil, DropNone
	}
	absoluteY := float32(dragged.cachedY) + localY
	return t.resolveDropTargetByTreeY(absoluteY)
}

func (t *Tree) resolveDropTargetByTreeY(treeY float32) (*Node, DropPosition) {
	for _, node := range t.flatVisibleNodes {
		top := float32(node.cachedY)
		h := float32(node.cachedH)
		if h <= 0 {
			h = 32
		}
		bottom := top + h

		if treeY >= top && treeY < bottom {
			relY := treeY - top

			if node.isFolder {
				if relY >= h*0.15 && relY <= h*0.85 {
					return node, DropInside
				}
				if relY < h*0.15 {
					return node, DropBefore
				}
				return node, DropAfter
			}

			if relY < h*0.5 {
				return node, DropBefore
			}
			return node, DropAfter
		}
	}

	// Clamped boundary checks
	if len(t.flatVisibleNodes) > 0 {
		last := t.flatVisibleNodes[len(t.flatVisibleNodes)-1]
		if treeY >= float32(last.cachedY+last.cachedH) {
			return last, DropAfter
		}
		first := t.flatVisibleNodes[0]
		if treeY < float32(first.cachedY) {
			return first, DropBefore
		}
	}

	return nil, DropNone
}

// MoveNode moves source relative to target within this Tree.
func (t *Tree) MoveNode(source, target *Node, pos DropPosition) {
	t.Session.MoveNodeCrossTree(t, t, source, target, pos)
}

func insertNode(slice []*Node, index int, item *Node) []*Node {
	if index < 0 {
		index = 0
	}
	if index > len(slice) {
		index = len(slice)
	}
	result := make([]*Node, len(slice)+1)
	copy(result, slice[:index])
	result[index] = item
	copy(result[index+1:], slice[index:])
	return result
}

func (t *Tree) isDescendant(parent, target *Node) bool {
	if parent == nil || target == nil {
		return false
	}
	for _, child := range parent.Children {
		if child == target || t.isDescendant(child, target) {
			return true
		}
	}
	return false
}

func (t *Tree) removeNode(parent *Node, list []*Node, target *Node) bool {
	for i, n := range list {
		if n == target {
			if parent != nil {
				parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			} else {
				t.Nodes = append(t.Nodes[:i], t.Nodes[i+1:]...)
			}
			return true
		}
		if len(n.Children) > 0 {
			if t.removeNode(n, n.Children, target) {
				return true
			}
		}
	}
	return false
}

func (t *Tree) findParentAndIndex(parent *Node, list []*Node, target *Node) (*Node, int) {
	for i, n := range list {
		if n == target {
			return parent, i
		}
		if len(n.Children) > 0 {
			if p, idx := t.findParentAndIndex(n, n.Children, target); p != nil || idx != -1 {
				return p, idx
			}
		}
	}
	return nil, -1
}

func (t *Tree) layoutNode(gtx layout.Context, th *theme.Theme, mTheme *material.Theme, node *Node) layout.Dimensions {
	indentPx := gtx.Dp(t.Indent) * node.depth

	// Toggle expand on chevron click
	for node.chevronBtn.Clicked(gtx) {
		if !node.Disabled {
			node.Expanded = !node.Expanded
		}
	}

	// Select on node click (evaluated before layout)
	for node.clickable.Clicked(gtx) {
		if !node.Disabled && !t.Session.WasDragging {
			t.SelectNode(node)
		}
	}

	rowBg := color.NRGBA{}
	fgColor := th.Colors.Foreground
	mutedColor := th.Colors.MutedFg

	if node.Selected {
		rowBg = th.Colors.Secondary
	} else if node.clickable.Hovered() {
		rowBg = th.Colors.Muted
	}

	if node.Disabled {
		fgColor.A = 100
		mutedColor.A = 100
	}

	// Determine icon
	nodeIcon := node.Icon
	if nodeIcon == nil {
		if node.isFolder {
			if node.Expanded {
				nodeIcon = lucide.FolderOpen
			} else {
				nodeIcon = lucide.Folder
			}
		} else {
			nodeIcon = lucide.File
		}
	}

	padding := layout.Inset{
		Top:    th.Spacing.Space1,
		Bottom: th.Spacing.Space1,
		Left:   unit.Dp(float32(indentPx) + 6),
		Right:  th.Spacing.Space2,
	}

	gtxContent := gtx
	gtxContent.Constraints.Min = image.Pt(0, 0)

	renderContent := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				// 1. Expand/Collapse Chevron (Default Chevron icon)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if node.isFolder {
						return node.chevronBtn.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Right: th.Spacing.Space1}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								if node.Expanded {
									return lucide.ChevronDown.LayoutSize(gtx, unit.Dp(14), mutedColor)
								}
								return lucide.ChevronRight.LayoutSize(gtx, unit.Dp(14), mutedColor)
							})
						})
					}
					// Leaf spacer matching chevron width
					return layout.Spacer{Width: unit.Dp(18)}.Layout(gtx)
				}),

				// 2. Subicon (Folder, File, or custom icon)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if nodeIcon != nil {
						return layout.Inset{Right: th.Spacing.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return nodeIcon.LayoutSize(gtx, unit.Dp(16), mutedColor)
						})
					}
					return layout.Dimensions{}
				}),

				// 3. Label Text
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeSM, node.Label)
					lbl.Color = fgColor
					lbl.Font.Weight = font.Medium
					lbl.Alignment = text.Start
					return lbl.Layout(gtx)
				}),

				// 4. Right-Aligned Trailing Actions (Spaced by default, Vert Center-aligned, NOT stretched vertically)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if len(node.Actions) == 0 {
						return layout.Dimensions{}
					}
					actionChildren := make([]layout.FlexChild, 0, len(node.Actions)*2)
					for idx, act := range node.Actions {
						actionWidget := act
						actionChildren = append(actionChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return actionWidget(gtx)
						}))
						if idx < len(node.Actions)-1 {
							actionChildren = append(actionChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Width: th.Spacing.Space1}.Layout(gtx)
							}))
						}
					}
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, actionChildren...)
				}),
			)
		})
	}

	contentDims := renderContent(gtxContent)
	rowSize := image.Pt(gtx.Constraints.Max.X, contentDims.Size.Y)

	isDropTarget := t.Session.DropTargetTree == t && t.Session.DropTargetNode == node
	isDragged := t.Session.DraggedNode == node && t.Session.WasDragging

	if isDragged {
		rowBg.A = 80
	}
	if isDropTarget && t.Session.DropPosition == DropInside {
		rowBg = th.Colors.Secondary
		fgColor = th.Colors.Foreground
	}

	return node.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Only register drag gesture when node is selected
		if node.Selected && node.Draggable && !node.Disabled {
			node.drag.Add(gtx.Ops)
		}

		return layout.Stack{}.Layout(gtx,
			// Row Background drawn FIRST
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{Max: rowSize}
				radius := gtx.Dp(th.Radius.RadiusSM)
				theme.DrawRRectBackground(gtx, rect, radius, rowBg)

				// Draw Drop Indicators (Before / After insertion lines)
				if isDropTarget && t.Session.WasDragging {
					lineColor := th.Colors.Primary
					indicatorX := indentPx + gtx.Dp(unit.Dp(6))

					if t.Session.DropPosition == DropBefore {
						lineRect := image.Rectangle{
							Min: image.Pt(indicatorX, 0),
							Max: image.Pt(rowSize.X-gtx.Dp(unit.Dp(8)), gtx.Dp(unit.Dp(2))),
						}
						theme.DrawRRectBackground(gtx, lineRect, 0, lineColor)

						// Left indicator dot
						dotRect := image.Rectangle{
							Min: image.Pt(indicatorX-4, -2),
							Max: image.Pt(indicatorX+2, 4),
						}
						theme.DrawRRectBackground(gtx, dotRect, 3, lineColor)
					} else if t.Session.DropPosition == DropAfter {
						lineRect := image.Rectangle{
							Min: image.Pt(indicatorX, rowSize.Y-gtx.Dp(unit.Dp(2))),
							Max: image.Pt(rowSize.X-gtx.Dp(unit.Dp(8)), rowSize.Y),
						}
						theme.DrawRRectBackground(gtx, lineRect, 0, lineColor)

						// Left indicator dot
						dotRect := image.Rectangle{
							Min: image.Pt(indicatorX-4, rowSize.Y-4),
							Max: image.Pt(indicatorX+2, rowSize.Y+2),
						}
						theme.DrawRRectBackground(gtx, dotRect, 3, lineColor)
					} else if t.Session.DropPosition == DropInside {
						theme.DrawRRectBackground(gtx, rect, radius, th.Colors.Muted)
					}
				}

				return layout.Dimensions{Size: rowSize}
			}),

			// Row content drawn ON TOP of background
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return renderContent(gtx)
			}),
		)
	})
}

func (t *Tree) layoutDragGhost(gtx layout.Context, th *theme.Theme, mTheme *material.Theme) {
	if t.Session.DraggedNode == nil {
		return
	}

	ghostOffset := image.Pt(int(t.Session.LocalDragX)+12, int(t.Session.DraggedNode.cachedY)+int(t.Session.LocalDragY)+8)
	trans := op.Offset(ghostOffset).Push(gtx.Ops)
	defer trans.Pop()

	padding := layout.Inset{
		Top:    unit.Dp(4),
		Bottom: unit.Dp(4),
		Left:   unit.Dp(8),
		Right:  unit.Dp(8),
	}

	renderGhost := func(gtx layout.Context) layout.Dimensions {
		return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					ic := t.Session.DraggedNode.Icon
					if ic == nil {
						if t.Session.DraggedNode.isFolder {
							ic = lucide.Folder
						} else {
							ic = lucide.File
						}
					}
					return layout.Inset{Right: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return ic.LayoutSize(gtx, unit.Dp(14), th.Colors.PrimaryFg)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Label(mTheme, th.Typography.FontSizeXS, t.Session.DraggedNode.Label)
					lbl.Color = th.Colors.PrimaryFg
					lbl.Font.Weight = font.SemiBold
					return lbl.Layout(gtx)
				}),
			)
		})
	}

	gtxGhost := gtx
	gtxGhost.Constraints.Min = image.Pt(0, 0)
	ghostDims := renderGhost(gtxGhost)

	ghostRect := image.Rectangle{Max: ghostDims.Size}
	theme.DrawRRectBackground(gtx, ghostRect, gtx.Dp(th.Radius.RadiusSM), th.Colors.Primary)
	renderGhost(gtx)
}
