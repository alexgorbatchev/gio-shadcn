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
    - empty.go:3
  - tests:
    - empty_test.go:2

### Capabilities & Features
- [x] **Illustrated Icon / Vector**
  - implementation:
    - empty.go:81
  - tests:
    - empty_test.go:15
- [x] **Title & Description**
  - implementation:
    - empty.go:24
  - tests:
    - empty_test.go:19
- [x] **Primary Action Button**
  - implementation:
    - empty.go:55
  - tests:
    - empty_test.go:17

---

## Code Structure & Entry Points
- `empty.go`: Primary component widget layout and state logic.
- `empty_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
