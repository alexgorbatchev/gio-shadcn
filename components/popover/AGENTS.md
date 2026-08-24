# Popover Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/popover](https://ui.shadcn.com/docs/components/popover)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `popover` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Anchored Card Popover Variant**
  - implementation:
    - popover.go:37
  - tests:
    - popover_test.go:13

### Capabilities & Features
- [x] **Open / Closed Visibility State**
  - implementation:
    - popover.go:50
  - tests:
    - popover_test.go:24
- [x] **Title & Description Layout**
  - implementation:
    - popover.go:81
  - tests:
    - popover_test.go:39
- [x] **Border & Background Card Drawing**
  - implementation:
    - popover.go:111
  - tests:
    - popover_test.go:39

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/popover-alignments.tsx
  - implementation:
    - component: popover.go:49
    - demo: demo.go:13
  - tests:
    - popover_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/popover-basic.tsx
  - implementation:
    - component: popover.go:49
    - demo: demo.go:13
  - tests:
    - popover_test.go:24
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/popover-demo.tsx
  - implementation:
    - component: popover.go:49
    - demo: demo.go:13
  - tests:
    - popover_test.go:39
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/popover-form.tsx
  - implementation:
    - component: popover.go:49
    - demo: demo.go:13
  - tests:
    - popover_test.go:39

---

## Code Structure & Entry Points
- `popover.go`: Primary component widget layout and state logic.
- `popover_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Modular component interactive demo layout (`components/popover/demo.go`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
