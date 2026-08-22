# Button Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/button](https://ui.shadcn.com/docs/components/button)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `button` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default / Primary**
  - implementation:
    - button.go:132
  - tests:
    - button_test.go:11
- [x] **Secondary**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:18
- [x] **Outline**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:25
- [x] **Ghost**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:32
- [x] **Destructive**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:39
- [x] **Link**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:46

### Capabilities & Features
- [x] **Small (SM) Size**
  - implementation:
    - button.go:154
  - tests:
    - button_test.go:53
- [x] **Default Size**
  - implementation:
    - button.go:154
  - tests:
    - button_test.go:60
- [x] **Large (LG) Size**
  - implementation:
    - button.go:154
  - tests:
    - button_test.go:67
- [x] **Icon Size**
  - implementation:
    - button.go:154
  - tests:
    - button_test.go:74
- [x] **Disabled State**
  - implementation:
    - button.go:174
  - tests:
    - button_test.go:81
- [x] **Pointer Click Event**
  - implementation:
    - button.go:149
  - tests:
    - button_test.go:88

### Demos (Official shadcn Demos)
- [x] **1. Default Primary**
  - implementation:
    - button.go:132
  - tests:
    - button_test.go:11
- [x] **2. Secondary**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:18
- [x] **3. Outline**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:25
- [x] **4. Ghost**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:32
- [x] **5. Destructive**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:39
- [x] **6. Link**
  - implementation:
    - button.go:153
  - tests:
    - button_test.go:46

---

## Code Structure & Entry Points
- `button.go`: Primary component widget layout and state logic.
- `button_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
