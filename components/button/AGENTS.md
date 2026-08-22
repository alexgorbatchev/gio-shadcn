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
    - button.go:107
  - tests:
    - button_test.go:13
- [x] **Secondary**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:21
- [x] **Outline**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:29
- [x] **Ghost**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:37
- [x] **Destructive**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:45
- [x] **Link**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:53

### Capabilities & Features
- [x] **Small (SM) Size**
  - implementation:
    - button.go:296
  - tests:
    - button_test.go:61
- [x] **Default Size**
  - implementation:
    - button.go:318
  - tests:
    - button_test.go:68
- [x] **Large (LG) Size**
  - implementation:
    - button.go:304
  - tests:
    - button_test.go:75
- [x] **Icon Size**
  - implementation:
    - button.go:311
  - tests:
    - button_test.go:82
- [x] **Disabled State**
  - implementation:
    - button.go:170
  - tests:
    - button_test.go:89
- [x] **Pointer Click Event**
  - implementation:
    - button.go:151
  - tests:
    - button_test.go:96

---

### Demos
- [x] **Default Primary**
  - implementation:
    - button.go:85
  - tests:
    - button_test.go:11
- [x] **Secondary**
  - implementation:
    - button.go:113
  - tests:
    - button_test.go:17
- [x] **Outline**
  - implementation:
    - button.go:114
  - tests:
    - button_test.go:25
- [x] **Ghost**
  - implementation:
    - button.go:115
  - tests:
    - button_test.go:31
- [x] **Destructive**
  - implementation:
    - button.go:116
  - tests:
    - button_test.go:39
- [x] **Link**
  - implementation:
    - button.go:117
  - tests:
    - button_test.go:45
- [x] **Icon Only**
  - implementation:
    - button.go:28
  - tests:
    - button_test.go:74
- [x] **With Icon**
  - implementation:
    - button.go:41
  - tests:
    - button_test.go:66
- [x] **Loading State**
  - implementation:
    - button.go:120
  - tests:
    - button_test.go:74

## Code Structure & Entry Points
- `button.go`: Primary component widget layout and state logic.
- `button_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
