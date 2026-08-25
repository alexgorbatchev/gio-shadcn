# Tree View Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/tree-view](https://ui.shadcn.com/docs/components/tree-view)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `tree` component implementation or unit tests in this directory, developers and AI agents MUST:
1. **Maintain Parity:** Preserve strict 100% visual token and feature parity with `shadcn/ui` and the source-of-truth Flutter POC implementation.
2. **Keep Checklist Updated:** Update the feature checklist below if new variants, subcomponents, configuration options, or interactive states are added.
3. **GPU State Safety:** Ensure all rounded rectangle drawings pass through `theme.DrawRRectBackground` / `theme.DrawStroke` with radius clamped to half the component height (`heightPx / 2`).
4. **Prevent Color State Leaks:** End all `Layout` functions with an explicit GPU color state reset:
   ```go
   paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
   ```
5. **Unit Test Coverage:** Verify that `go test ./...` passes cleanly with zero compilation or `go vet` errors.

---

## Component Variants & Features Checklist

### Variants
- [x] **File System Hierarchy Tree**
  - implementation:
    - tree.go:41
  - tests:
    - tree_test.go:13
- [x] **Folder & File Nodes**
  - implementation:
    - tree.go:73
  - tests:
    - tree_test.go:34

### Capabilities & Features
- [x] **Configurable Indentation Level (`Indent: unit.Dp(9)` default)**
  - implementation:
    - tree.go:129
  - tests:
    - tree_test.go:217
- [x] **Right-Aligned Trailing Action Elements (Spaced & Vertically Centered)**
  - implementation:
    - tree.go:550
  - tests:
    - tree_test.go:235
- [x] **Chevron Expand / Collapse Toggle (Default Icon)**
  - implementation:
    - tree.go:345
  - tests:
    - tree_test.go:34
- [x] **Custom Subicons (Folders, Files, Code, Media)**
  - implementation:
    - tree.go:318
  - tests:
    - tree_test.go:68
- [x] **Cross-Tree Drag & Drop (`DragSession`)**
  - implementation:
    - tree.go:107
  - tests:
    - tree_test.go:253
- [x] **Target Tree Selection Preservation on Drop**
  - implementation:
    - tree.go:114
  - tests:
    - tree_test.go:338
- [x] **Visual Drop Insertion Line Indicators & Drop Inside Highlight**
  - implementation:
    - tree.go:405
  - tests:
    - tree_test.go:148
- [x] **Floating Drag Preview Ghost Badge**
  - implementation:
    - tree.go:450
  - tests:
    - tree_test.go:148
- [x] **Node Selection & Highlighting (Select on Click)**
  - implementation:
    - tree.go:135
  - tests:
    - tree_test.go:49

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/collapsible-file-tree.tsx
  - implementation:
    - component: tree.go:125
    - demo: demo.go:28
  - tests:
    - tree_test.go:13

---

## Code Structure & Entry Points
- `tree.go`: Primary tree view widget implementation with right-aligned action elements, configurable indentation, target selection preservation, and cross-tree DnD gesture support.
- `demo.go`: Interactive Dual-Tree File System Split View demo widget (`Demo`).
- `tree_test.go`: Automated unit test suite verifying layout dimensions, expand/collapse, selection, configurable indent, trailing actions, cross-tree DnD, and selection preservation.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
