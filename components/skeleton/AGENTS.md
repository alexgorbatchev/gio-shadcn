# Skeleton Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/skeleton](https://ui.shadcn.com/docs/components/skeleton)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `skeleton` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Text Line Skeleton**
  - implementation:
    - skeleton.go:31
  - tests:
    - skeleton_test.go:13
- [x] **Circular Avatar Skeleton**
  - implementation:
    - skeleton.go:31
  - tests:
    - skeleton_test.go:24
- [x] **Card Skeleton Container**
  - implementation:
    - skeleton.go:31
  - tests:
    - skeleton_test.go:35

### Capabilities & Features
- [x] **Muted Shimmer Background**
  - implementation:
    - skeleton.go:56
  - tests:
    - skeleton_test.go:46
- [x] **Custom Width & Height Dimensions**
  - implementation:
    - skeleton.go:51
  - tests:
    - skeleton_test.go:13

### Demos (All Official shadcn Demos)
- [x] **1. Text Line Skeleton**
  - implementation:
    - component: skeleton.go:47
    - demo: demo.go:29
  - tests:
    - skeleton_test.go:13
- [x] **2. Circular Avatar Skeleton**
  - implementation:
    - component: skeleton.go:47
    - demo: demo.go:41
  - tests:
    - skeleton_test.go:24
- [x] **3. Card Container Skeleton**
  - implementation:
    - component: skeleton.go:47
    - demo: demo.go:53
  - tests:
    - skeleton_test.go:35

---

## Code Structure & Entry Points
- `skeleton.go`: Primary component widget layout and state logic.
- `skeleton_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Interactive component gallery demo.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
