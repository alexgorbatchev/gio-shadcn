# Input Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/input](https://ui.shadcn.com/docs/components/input)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `input` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Standard Text Input**
  - implementation:
    - input.go:26
  - tests:
    - input_test.go:2
- [x] **Prefilled Text Input**
  - implementation:
    - input.go:4
  - tests:
    - input_test.go:3
- [x] **Disabled Text Input**
  - implementation:
    - input.go:147
  - tests:
    - input_test.go:4

### Capabilities & Features
- [x] **Single-Line Text Editing**
  - implementation:
    - input.go:3
  - tests:
    - input_test.go:20
- [x] **Placeholder Text**
  - implementation:
    - input.go:7
  - tests:
    - input_test.go:16
- [x] **Focus Ring Stroke**
  - implementation:
    - input.go:33
  - tests:
    - input_test.go:14

---

## Code Structure & Entry Points
- `input.go`: Primary component widget layout and state logic.
- `input_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
