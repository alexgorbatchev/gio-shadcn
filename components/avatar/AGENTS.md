# Avatar Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/avatar](https://ui.shadcn.com/docs/components/avatar)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/avatar.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/avatar.mdx)

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

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-badge-icon.tsx
  - implementation:
    - component: avatar.go:102
    - demo: demo.go:14
  - tests:
    - avatar_test.go:31
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-badge.tsx
  - implementation:
    - component: avatar.go:102
    - demo: demo.go:14
  - tests:
    - avatar_test.go:31
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-basic.tsx
  - implementation:
    - component: avatar.go:37
    - demo: demo.go:10
  - tests:
    - avatar_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-demo.tsx
  - implementation:
    - component: avatar.go:37
    - demo: demo.go:10
  - tests:
    - avatar_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-dropdown.tsx
  - implementation:
    - component: avatar.go:37
    - demo: demo.go:10
  - tests:
    - avatar_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group-count-icon.tsx
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:20
  - tests:
    - avatar_test.go:56
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group-count.tsx
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:20
  - tests:
    - avatar_test.go:56
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-group.tsx
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:20
  - tests:
    - avatar_test.go:56
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-rtl.tsx
  - implementation:
    - component: avatar.go:37
    - demo: demo.go:10
  - tests:
    - avatar_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/avatar-size.tsx
  - implementation:
    - component: avatar.go:32
    - demo: demo.go:20
  - tests:
    - avatar_test.go:56

---

## Code Structure & Entry Points
- `avatar.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `avatar_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
