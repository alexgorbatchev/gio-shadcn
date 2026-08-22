# Tabs Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/tabs](https://ui.shadcn.com/docs/components/tabs)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `tabs` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Tabs Navigation Header**
  - implementation:
    - tabs.go:2
  - tests:
    - tabs_test.go:1

### Capabilities & Features
- [x] **Tab Selection Buttons**
  - implementation:
    - tabs.go:2
  - tests:
    - tabs_test.go:1
- [x] **Active Tab Highlight**
  - implementation:
    - tabs.go:4
  - tests:
    - tabs_test.go:19
- [x] **Selected Key State**
  - implementation:
    - tabs.go:69
  - tests:
    - tabs_test.go:13

---

## Code Structure & Entry Points
- `tabs.go`: Primary component widget layout and state logic.
- `tabs_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
