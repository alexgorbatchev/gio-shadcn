# Number Input Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/number-input](https://ui.shadcn.com/docs/components/number-input)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `numberinput` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Number Stepper**
  - implementation:
    - numberinput.go:2
  - tests:
    - numberinput_test.go:1

### Capabilities & Features
- [x] **Increment / Decrement Buttons**
  - implementation:
    - numberinput.go:4
  - tests:
    - numberinput_test.go:13
- [x] **Min / Max Range Bounds**
  - implementation:
    - numberinput.go:27
  - tests:
    - numberinput_test.go:17
- [x] **Step Size Configuration**
  - implementation:
    - numberinput.go:2
  - tests:
    - numberinput_test.go:16
- [x] **Direct Number Formatting**
  - implementation:
    - numberinput.go:75
  - tests:
    - numberinput_test.go:13

---

## Code Structure & Entry Points
- `numberinput.go`: Primary component widget layout and state logic.
- `numberinput_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
