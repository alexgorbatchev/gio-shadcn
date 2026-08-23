# Accordion Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/accordion](https://ui.shadcn.com/docs/components/accordion)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `accordion` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Single Open Mode (`TypeSingle`)**
  - implementation:
    - accordion.go:107
  - tests:
    - accordion_test.go:13
- [x] **Multiple Open Mode (`TypeMultiple`)**
  - implementation:
    - accordion.go:114
  - tests:
    - accordion_test.go:25
- [x] **Disabled Item State (`Disabled`)**
  - implementation:
    - accordion.go:106
  - tests:
    - accordion_test.go:37
- [x] **Borderless Variant (`Borderless`)**
  - implementation:
    - accordion.go:220
  - tests:
    - accordion_test.go:61

### Capabilities & Features
- [x] **Expand / Collapse Toggle**
  - implementation:
    - accordion.go:107
  - tests:
    - accordion_test.go:13
- [x] **Custom Icon Indicator**
  - implementation:
    - accordion.go:180
  - tests:
    - accordion_test.go:49
- [x] **Custom Header Widget**
  - implementation:
    - accordion.go:167
  - tests:
    - accordion_test.go:49
- [x] **Nested Content Widget**
  - implementation:
    - accordion.go:204
  - tests:
    - accordion_test.go:72
- [x] **Border & Background Drawing**
  - implementation:
    - accordion.go:221
  - tests:
    - accordion_test.go:87

### Demos (All 8 Official shadcn Demos)
- [x] **1. Single Open Accordion (Default)**
  - implementation:
    - component: accordion.go:107
    - demo: demo.go:9
  - tests:
    - accordion_test.go:13
- [x] **2. Multiple Open Accordion**
  - implementation:
    - component: accordion.go:114
    - demo: demo.go:17
  - tests:
    - accordion_test.go:25
- [x] **3. Disabled Item Accordion**
  - implementation:
    - component: accordion.go:106
    - demo: demo.go:25
  - tests:
    - accordion_test.go:37
- [x] **4. Chevron Icon Accordion**
  - implementation:
    - component: accordion.go:180
    - demo: demo.go:31
  - tests:
    - accordion_test.go:49
- [x] **5. Custom Header Badge Section**
  - implementation:
    - component: accordion.go:167
    - demo: demo.go:37
  - tests:
    - accordion_test.go:49
- [x] **6. Borderless Variant Accordion**
  - implementation:
    - component: accordion.go:220
    - demo: demo.go:43
  - tests:
    - accordion_test.go:61
- [x] **7. Nested Accordion**
  - implementation:
    - component: accordion.go:204
    - demo: demo.go:57
  - tests:
    - accordion_test.go:72
- [x] **8. Controlled Accordion State**
  - implementation:
    - component: accordion.go:107
    - demo: demo.go:68
  - tests:
    - accordion_test.go:25

---

## Code Structure & Entry Points
- `accordion.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `accordion_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
