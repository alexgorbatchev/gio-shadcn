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
    - collapsible.go:58
  - tests:
    - collapsible_test.go:13
- [x] **Collapsed State**
  - implementation:
    - collapsible.go:103
  - tests:
    - collapsible_test.go:24

### Capabilities & Features
- [x] **Trigger Button Header**
  - implementation:
    - collapsible.go:88
  - tests:
    - collapsible_test.go:35
- [x] **Content Body Visibility Toggle**
  - implementation:
    - collapsible.go:111
  - tests:
    - collapsible_test.go:44

---

### Demos
- [x] **Simple Collapsible**
  - implementation:
    - collapsible.go:3
  - tests:
    - collapsible_test.go:25
- [x] **With Trigger Button**
  - implementation:
    - collapsible.go:55
  - tests:
    - collapsible_test.go:37
- [x] **Nested Content Panel**
  - implementation:
    - collapsible.go:56
  - tests:
    - collapsible_test.go:47

## Code Structure & Entry Points
- `collapsible.go`: Primary component widget layout and state logic.
- `collapsible_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
