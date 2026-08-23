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
    - input.go:250
  - tests:
    - input_test.go:13
- [x] **Prefilled Text Input**
  - implementation:
    - input.go:251
  - tests:
    - input_test.go:21
- [x] **Disabled Text Input**
  - implementation:
    - input.go:275
  - tests:
    - input_test.go:30

### Capabilities & Features
- [x] **Single-Line Text Editing**
  - implementation:
    - input.go:270
  - tests:
    - input_test.go:40
- [x] **Placeholder Text**
  - implementation:
    - input.go:323
  - tests:
    - input_test.go:49
- [x] **Focus Ring Stroke**
  - implementation:
    - input.go:338
  - tests:
    - input_test.go:57

---



### Demos
- [x] **1. Text Input Field**
  - implementation:
    - component: input.go:50
    - demo: demo.md:13
  - tests:
    - input_test.go:15

## Code Structure & Entry Points
- `input.go`: Primary component implementation.
- `input_test.go`: Unit test suite.
- `demo.md`: Component interactive demo snippets.
- `AGENTS.md`: Component specification.