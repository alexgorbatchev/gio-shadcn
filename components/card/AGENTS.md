# Card Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/card](https://ui.shadcn.com/docs/components/card)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `card` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default Card**
  - implementation:
    - card.go:42
  - tests:
    - card_test.go:13
- [x] **Header + Content + Footer Layout**
  - implementation:
    - card.go:133
  - tests:
    - card_test.go:20

### Capabilities & Features
- [x] **Card Title**
  - implementation:
    - card.go:215
  - tests:
    - card_test.go:43
- [x] **Card Description**
  - implementation:
    - card.go:240
  - tests:
    - card_test.go:50
- [x] **Card Content Area**
  - implementation:
    - card.go:258
  - tests:
    - card_test.go:57
- [x] **Card Footer**
  - implementation:
    - card.go:276
  - tests:
    - card_test.go:64
- [x] **Border Stroke & Rounded Radius**
  - implementation:
    - card.go:142
  - tests:
    - card_test.go:71

---

## Code Structure & Entry Points
- `card.go`: Primary component widget layout and state logic.
- `card_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
