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
    - avatar.go:35
  - tests:
    - avatar_test.go:12
- [x] **Image Avatar Placeholder**
  - implementation:
    - avatar.go:65
  - tests:
    - avatar_test.go:12
- [x] **Online Status Badge Indicator**
  - implementation:
    - avatar.go:102
  - tests:
    - avatar_test.go:21

### Capabilities & Features
- [x] **Circular Clip Ellipse**
  - implementation:
    - avatar.go:76
  - tests:
    - avatar_test.go:62
- [x] **Custom Sizes (32px, 40px, 56px)**
  - implementation:
    - avatar.go:54
  - tests:
    - avatar_test.go:42
- [x] **Status Dot Indicator**
  - implementation:
    - avatar.go:107
  - tests:
    - avatar_test.go:21

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-badge-icon.tsx
  - implementation:
    - component: avatar.go:102
    - demo: demo.go:40
  - tests:
    - avatar_test.go:32
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-badge.tsx
  - implementation:
    - component: avatar.go:102
    - demo: demo.go:33
  - tests:
    - avatar_test.go:21
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-basic.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:27
  - tests:
    - avatar_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-demo.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:27
  - tests:
    - avatar_test.go:62
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-dropdown.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:33
  - tests:
    - avatar_test.go:21
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group-count-icon.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:53
  - tests:
    - avatar_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group-count.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:53
  - tests:
    - avatar_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:50
  - tests:
    - avatar_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-rtl.tsx
  - implementation:
    - component: avatar.go:35
    - demo: demo.go:50
  - tests:
    - avatar_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-size.tsx
  - implementation:
    - component: avatar.go:54
    - demo: demo.go:46
  - tests:
    - avatar_test.go:42

---

## Code Structure & Entry Points
- `avatar.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `avatar_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
