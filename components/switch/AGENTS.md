# Switch Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/switch](https://ui.shadcn.com/docs/components/switch)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `switch` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Off State Switch**
  - implementation:
    - switch.go:36
  - tests:
    - switch_test.go:13
- [x] **On State Switch**
  - implementation:
    - switch.go:36
  - tests:
    - switch_test.go:23
- [x] **Disabled Switch State**
  - implementation:
    - switch.go:75
  - tests:
    - switch_test.go:33

### Capabilities & Features
- [x] **Track Fill Transition**
  - implementation:
    - switch.go:68
  - tests:
    - switch_test.go:23
- [x] **Sliding Thumb Knob**
  - implementation:
    - switch.go:106
  - tests:
    - switch_test.go:56
- [x] **Interactive Click Toggle**
  - implementation:
    - switch.go:52
  - tests:
    - switch_test.go:43

---


### Demos
- [x] **1. Switch Toggle Button**
  - implementation:
    - component: switch.go:50
    - demo: demo.md:13
  - tests:
    - switch_test.go:15

## Code Structure & Entry Points
- `switch.go`: Primary component widget layout and state logic.
- `switch_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.md`: Component interactive demo snippets and layout specs (`components/switch/demo.md`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
