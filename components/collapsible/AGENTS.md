# Collapsible Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/collapsible](https://ui.shadcn.com/docs/components/collapsible)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `collapsible` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Expanded State**
  - implementation:
    - collapsible.go:53
  - tests:
    - collapsible_test.go:13
- [x] **Collapsed State**
  - implementation:
    - collapsible.go:108
  - tests:
    - collapsible_test.go:23

### Capabilities & Features
- [x] **Trigger Button Header**
  - implementation:
    - collapsible.go:90
  - tests:
    - collapsible_test.go:33
- [x] **Content Body Visibility Toggle**
  - implementation:
    - collapsible.go:58
  - tests:
    - collapsible_test.go:42

### Demos
- [x] **1. Simple Collapsible**
  - implementation:
    - component: collapsible.go:53
    - demo: demo.go:10
  - tests:
    - collapsible_test.go:13
- [x] **2. Collapsed State Container**
  - implementation:
    - component: collapsible.go:108
    - demo: demo.go:15
  - tests:
    - collapsible_test.go:23

---

## Code Structure & Entry Points
- `collapsible.go`: Primary component widget layout and state logic.
- `collapsible_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
