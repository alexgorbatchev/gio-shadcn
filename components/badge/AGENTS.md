# Badge Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/badge](https://ui.shadcn.com/docs/components/badge)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `badge` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Default / Primary Badge**
  - implementation:
    - badge.go:142
  - tests:
    - badge_test.go:12
- [x] **Secondary Badge**
  - implementation:
    - badge.go:136
  - tests:
    - badge_test.go:22
- [x] **Outline Badge**
  - implementation:
    - badge.go:138
  - tests:
    - badge_test.go:32
- [x] **Destructive Badge**
  - implementation:
    - badge.go:140
  - tests:
    - badge_test.go:42

### Capabilities & Features
- [x] **Full Rounded Radius**
  - implementation:
    - badge.go:103
  - tests:
    - badge_test.go:12
- [x] **XS Typography Label**
  - implementation:
    - badge.go:88
  - tests:
    - badge_test.go:12
- [x] **Compact Padding**
  - implementation:
    - badge.go:61
  - tests:
    - badge_test.go:12
- [x] **Icon Support**
  - implementation:
    - badge.go:73
  - tests:
    - badge_test.go:52

---

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-demo.tsx
  - implementation:
    - component: badge.go:142
    - demo: demo.go:22
  - tests:
    - badge_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-variants.tsx
  - implementation:
    - component: badge.go:136
    - demo: demo.go:23
  - tests:
    - badge_test.go:22
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-icon.tsx
  - implementation:
    - component: badge.go:73
    - demo: demo.go:26
  - tests:
    - badge_test.go:52
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-spinner.tsx
  - implementation:
    - component: badge.go:73
    - demo: demo.go:27
  - tests:
    - badge_test.go:62
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-link.tsx
  - implementation:
    - component: badge.go:53
    - demo: demo.go:28
  - tests:
    - badge_test.go:72
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-colors.tsx
  - implementation:
    - component: badge.go:68
    - demo: demo.go:29
  - tests:
    - badge_test.go:82
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-rtl.tsx
  - implementation:
    - component: badge.go:53
    - demo: demo.go:22
  - tests:
    - badge_test.go:12

---

## Code Structure & Entry Points
- `badge.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `badge_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
