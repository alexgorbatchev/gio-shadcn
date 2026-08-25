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
    - button.go:50
  - tests:
    - button_test.go:12
- [x] **Secondary**
  - implementation:
    - button.go:68
  - tests:
    - button_test.go:22
- [x] **Outline**
  - implementation:
    - button.go:68
  - tests:
    - button_test.go:32
- [x] **Ghost**
  - implementation:
    - button.go:68
  - tests:
    - button_test.go:42
- [x] **Destructive**
  - implementation:
    - button.go:68
  - tests:
    - button_test.go:52
- [x] **Link**
  - implementation:
    - button.go:68
  - tests:
    - button_test.go:62

### Capabilities & Features
- [x] **Small (SM) Size**
  - implementation:
    - button.go:210
  - tests:
    - button_test.go:92
- [x] **Default Size**
  - implementation:
    - button.go:230
  - tests:
    - button_test.go:12
- [x] **Large (LG) Size**
  - implementation:
    - button.go:217
  - tests:
    - button_test.go:92
- [x] **Icon Size**
  - implementation:
    - button.go:224
  - tests:
    - button_test.go:82
- [x] **Disabled State**
  - implementation:
    - button.go:88
  - tests:
    - button_test.go:12
- [x] **Pointer Click Event**
  - implementation:
    - button.go:107
  - tests:
    - button_test.go:12
- [x] **Icon Support**
  - implementation:
    - button.go:183
  - tests:
    - button_test.go:72

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-default.tsx
  - implementation:
    - component: button.go:50
    - demo: demo.go:32
  - tests:
    - button_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-secondary.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:33
  - tests:
    - button_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-destructive.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:34
  - tests:
    - button_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-outline.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:35
  - tests:
    - button_test.go:32
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-ghost.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:36
  - tests:
    - button_test.go:42
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-link.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:37
  - tests:
    - button_test.go:62
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-with-icon.tsx
  - implementation:
    - component: button.go:183
    - demo: demo.go:39
  - tests:
    - button_test.go:72
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-icon.tsx
  - implementation:
    - component: button.go:224
    - demo: demo.go:40
  - tests:
    - button_test.go:82
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-spinner.tsx
  - implementation:
    - component: button.go:183
    - demo: demo.go:41
  - tests:
    - button_test.go:72
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-size.tsx
  - implementation:
    - component: button.go:210
    - demo: demo.go:43
  - tests:
    - button_test.go:92
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-group-demo.tsx
  - implementation:
    - component: button.go:68
    - demo: demo.go:47
  - tests:
    - button_test.go:32
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/button-demo.tsx
  - implementation:
    - component: button.go:50
    - demo: demo.go:32
  - tests:
    - button_test.go:12

---

## Code Structure & Entry Points
- `button.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `button_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
