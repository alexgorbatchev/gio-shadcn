# Empty Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/empty](https://ui.shadcn.com/docs/components/empty)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `empty` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Empty State Card**
  - implementation:
    - empty.go:34
  - tests:
    - empty_test.go:12

### Capabilities & Features
- [x] **Title Header & Description Body**
  - implementation:
    - empty.go:102
  - tests:
    - empty_test.go:12
- [x] **Illustrated Lucide Icon**
  - implementation:
    - empty.go:96
  - tests:
    - empty_test.go:12
- [x] **Action Widget Button**
  - implementation:
    - empty.go:119
  - tests:
    - empty_test.go:25
- [x] **Card Background & Border Stroke**
  - implementation:
    - empty.go:132
  - tests:
    - empty_test.go:12

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-demo.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-card.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-outline.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:30
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-avatar.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-avatar-group.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-background.tsx
  - implementation:
    - component: empty.go:132
    - demo: demo.go:21
  - tests:
    - empty_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-input-group.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:25
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-rtl.tsx
  - implementation:
    - component: empty.go:34
    - demo: demo.go:21
  - tests:
    - empty_test.go:12

---

## Code Structure & Entry Points
- `empty.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `empty_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
