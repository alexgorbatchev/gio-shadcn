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
- [x] **BPM Stepper Variant (128 BPM)**
  - implementation:
    - numberinput.go:41
  - tests:
    - numberinput_test.go:13
- [x] **Gain Range Stepper Variant (-12 to +12 dB)**
  - implementation:
    - numberinput.go:46
  - tests:
    - numberinput_test.go:24
- [x] **Quantity Range Stepper Variant (1 to 10)**
  - implementation:
    - numberinput.go:48
  - tests:
    - numberinput_test.go:35
- [x] **Decimal Precision Stepper Variant (0.0 to 1.0)**
  - implementation:
    - numberinput.go:50
  - tests:
    - numberinput_test.go:46
- [x] **Min/Max Bounded Stepper Variant**
  - implementation:
    - numberinput.go:52
  - tests:
    - numberinput_test.go:57

### Capabilities & Features
- [x] **Increment & Decrement Click Buttons**
  - implementation:
    - numberinput.go:73
  - tests:
    - numberinput_test.go:68
- [x] **OnValueChange Event Callback**
  - implementation:
    - numberinput.go:78
  - tests:
    - numberinput_test.go:84
- [x] **Value Display Box Formatting**
  - implementation:
    - numberinput.go:105
  - tests:
    - numberinput_test.go:98

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/data-picker-with-dropdowns.tsx
  - implementation:
    - component: numberinput.go:41
    - demo: demo.go:18
  - tests:
    - numberinput_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/date-picker-basic.tsx
  - implementation:
    - component: numberinput.go:46
    - demo: demo.go:19
  - tests:
    - numberinput_test.go:24
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/date-picker-demo.tsx
  - implementation:
    - component: numberinput.go:48
    - demo: demo.go:20
  - tests:
    - numberinput_test.go:35
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/date-picker-input.tsx
  - implementation:
    - component: numberinput.go:50
    - demo: demo.go:21
  - tests:
    - numberinput_test.go:46

---

## Code Structure & Entry Points
- `numberinput.go`: Primary component widget layout and state logic.
- `numberinput_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Modular component interactive demo layout (`components/numberinput/demo.go`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
