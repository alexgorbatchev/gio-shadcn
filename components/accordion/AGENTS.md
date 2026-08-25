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
    - accordion_test.go:76
- [x] **Disabled Item State (`Disabled`)**
  - implementation:
    - accordion.go:106
  - tests:
    - accordion_test.go:64
- [x] **Borderless Variant (`Borderless`)**
  - implementation:
    - accordion.go:220
  - tests:
    - accordion_test.go:27

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
    - accordion_test.go:52
- [x] **Custom Header Widget**
  - implementation:
    - accordion.go:167
  - tests:
    - accordion_test.go:64
- [x] **Nested Content Widget**
  - implementation:
    - accordion.go:204
  - tests:
    - accordion_test.go:40
- [x] **Border & Background Drawing**
  - implementation:
    - accordion.go:221
  - tests:
    - accordion_test.go:99

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-basic.tsx
  - implementation:
    - component: accordion.go:107
    - demo: demo.go:26
  - tests:
    - accordion_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-borders.tsx
  - implementation:
    - component: accordion.go:220
    - demo: demo.go:37
  - tests:
    - accordion_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-card.tsx
  - implementation:
    - component: accordion.go:221
    - demo: demo.go:48
  - tests:
    - accordion_test.go:40
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-demo.tsx
  - implementation:
    - component: accordion.go:107
    - demo: demo.go:59
  - tests:
    - accordion_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-disabled.tsx
  - implementation:
    - component: accordion.go:106
    - demo: demo.go:70
  - tests:
    - accordion_test.go:64
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-multiple.tsx
  - implementation:
    - component: accordion.go:114
    - demo: demo.go:81
  - tests:
    - accordion_test.go:76
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/accordion-rtl.tsx
  - implementation:
    - component: accordion.go:107
    - demo: demo.go:92
  - tests:
    - accordion_test.go:88

---

## Code Structure & Entry Points
- `accordion.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `accordion_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
