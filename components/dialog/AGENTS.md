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
    - dialog.go:21
  - tests:
    - dialog_test.go:13

### Capabilities & Features
- [x] **Dark Backdrop Overlay**
  - implementation:
    - dialog.go:129
  - tests:
    - dialog_test.go:23
- [x] **Backdrop Click-To-Close**
  - implementation:
    - dialog.go:109
  - tests:
    - dialog_test.go:39
- [x] **Header Title & Description Body**
  - implementation:
    - dialog.go:151
  - tests:
    - dialog_test.go:53
- [x] **Confirm & Cancel Action Buttons**
  - implementation:
    - dialog.go:175
  - tests:
    - dialog_test.go:62
- [x] **Custom Content Body Widget**
  - implementation:
    - dialog.go:167
  - tests:
    - dialog_test.go:71

### Demos (Official shadcn Demos)
- [x] **1. Modal Dialog**
  - implementation:
    - dialog.go:100
  - tests:
    - dialog_test.go:13
- [x] **2. Custom Content Body**
  - implementation:
    - dialog.go:167
  - tests:
    - dialog_test.go:71
- [x] **3. Form Action Dialog**
  - implementation:
    - dialog.go:175
  - tests:
    - dialog_test.go:62
- [x] **4. Confirmation Warning**
  - implementation:
    - dialog.go:151
  - tests:
    - dialog_test.go:39

---

## Code Structure & Entry Points
- `dialog.go`: Primary component widget layout and state logic.
- `dialog_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
