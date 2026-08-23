# Avatar Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/avatar](https://ui.shadcn.com/docs/components/avatar)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `avatar` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Text Initials Avatar**
  - implementation:
    - avatar.go:37
  - tests:
    - avatar_test.go:13
- [x] **Image Avatar Placeholder**
  - implementation:
    - avatar.go:37
  - tests:
    - avatar_test.go:22
- [x] **Online Status Badge Indicator**
  - implementation:
    - avatar.go:102
  - tests:
    - avatar_test.go:31

### Capabilities & Features
- [x] **Circular Clip Ellipse**
  - implementation:
    - avatar.go:66
  - tests:
    - avatar_test.go:41
- [x] **Custom Sizes (32px, 40px, 56px)**
  - implementation:
    - avatar.go:32
  - tests:
    - avatar_test.go:56
- [x] **Status Dot Indicator**
  - implementation:
    - avatar.go:102
  - tests:
    - avatar_test.go:65

### Demos (All Official shadcn Demos)
- [x] **1. Text Initials Avatar**
  - implementation:
    - component: avatar.go:37
    - demo: demo.go:10
  - tests:
    - avatar_test.go:13
- [x] **2. Online Status Badge Avatar**
  - implementation:
    - component: avatar.go:102
    - demo: demo.go:14
  - tests:
    - avatar_test.go:31
- [x] **3. Custom Small Avatar (32px)**
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:20
  - tests:
    - avatar_test.go:56
- [x] **4. Custom Large Avatar (56px)**
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:25
  - tests:
    - avatar_test.go:56

---

## Code Structure & Entry Points
- `avatar.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `avatar_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
