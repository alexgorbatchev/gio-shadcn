# Dialog Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/dialog](https://ui.shadcn.com/docs/components/dialog)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `dialog` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Modal Dialog Window**
  - implementation:
    - dialog.go:61
  - tests:
    - dialog_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay**
  - implementation:
    - dialog.go:138
  - tests:
    - dialog_test.go:24
- [x] **Backdrop Click-To-Close**
  - implementation:
    - dialog.go:117
  - tests:
    - dialog_test.go:40
- [x] **Header Title & Description**
  - implementation:
    - dialog.go:158
  - tests:
    - dialog_test.go:54
- [x] **Confirm & Cancel Actions**
  - implementation:
    - dialog.go:183
  - tests:
    - dialog_test.go:64
- [x] **Custom Content Widget**
  - implementation:
    - dialog.go:173
  - tests:
    - dialog_test.go:74

---

## Code Structure & Entry Points
- `dialog.go`: Primary component widget layout and state logic.
- `dialog_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
