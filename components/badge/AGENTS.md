# Badge Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/badge](https://ui.shadcn.com/docs/components/badge)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `badge` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default / Primary Badge**
  - implementation:
    - badge.go:135
  - tests:
    - badge_test.go:12
- [x] **Secondary Badge**
  - implementation:
    - badge.go:129
  - tests:
    - badge_test.go:22
- [x] **Outline Badge**
  - implementation:
    - badge.go:131
  - tests:
    - badge_test.go:32
- [x] **Destructive Badge**
  - implementation:
    - badge.go:133
  - tests:
    - badge_test.go:42

### Capabilities & Features
- [x] **Full Rounded Radius**
  - implementation:
    - badge.go:94
  - tests:
    - badge_test.go:52
- [x] **XS Typography Label**
  - implementation:
    - badge.go:78
  - tests:
    - badge_test.go:66
- [x] **Compact Padding**
  - implementation:
    - badge.go:58
  - tests:
    - badge_test.go:80

---

## Code Structure & Entry Points
- `badge.go`: Primary component widget layout and state logic.
- `badge_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
