# Sheet Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/sheet](https://ui.shadcn.com/docs/components/sheet)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `sheet` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Side Sheet Slide-Over Panel**
  - implementation:
    - sheet.go:48
  - tests:
    - sheet_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay**
  - implementation:
    - sheet.go:115
  - tests:
    - sheet_test.go:25
- [x] **Backdrop Click-To-Close**
  - implementation:
    - sheet.go:90
  - tests:
    - sheet_test.go:19
- [x] **Right Viewport Edge Alignment**
  - implementation:
    - sheet.go:124
  - tests:
    - sheet_test.go:33
- [x] **Header Title & Description**
  - implementation:
    - sheet.go:156
  - tests:
    - sheet_test.go:14
- [x] **Close Button**
  - implementation:
    - sheet.go:68
  - tests:
    - sheet_test.go:25
- [x] **Custom Content Widget**
  - implementation:
    - sheet.go:179
  - tests:
    - sheet_test.go:25

---

## Code Structure & Entry Points
- `sheet.go`: Primary component widget layout and state logic.
- `sheet_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
