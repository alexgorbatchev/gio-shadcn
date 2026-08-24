# Badge Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/badge](https://ui.shadcn.com/docs/components/badge)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/badge.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/badge.mdx)

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
    - badge.go:133
  - tests:
    - badge_test.go:13
- [x] **Secondary Badge**
  - implementation:
    - badge.go:127
  - tests:
    - badge_test.go:23
- [x] **Outline Badge**
  - implementation:
    - badge.go:129
  - tests:
    - badge_test.go:33
- [x] **Destructive Badge**
  - implementation:
    - badge.go:131
  - tests:
    - badge_test.go:43

### Capabilities & Features
- [x] **Full Rounded Radius**
  - implementation:
    - badge.go:96
  - tests:
    - badge_test.go:53
- [x] **XS Typography Label**
  - implementation:
    - badge.go:80
  - tests:
    - badge_test.go:67
- [x] **Compact Padding**
  - implementation:
    - badge.go:58
  - tests:
    - badge_test.go:80

---


### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-colors.tsx
  - implementation:
    - component: badge.go:127
    - demo: demo.go:14
  - tests:
    - badge_test.go:23
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-demo.tsx
  - implementation:
    - component: badge.go:133
    - demo: demo.go:13
  - tests:
    - badge_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-icon.tsx
  - implementation:
    - component: badge.go:133
    - demo: demo.go:13
  - tests:
    - badge_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-link.tsx
  - implementation:
    - component: badge.go:129
    - demo: demo.go:15
  - tests:
    - badge_test.go:33
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-rtl.tsx
  - implementation:
    - component: badge.go:133
    - demo: demo.go:13
  - tests:
    - badge_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-spinner.tsx
  - implementation:
    - component: badge.go:133
    - demo: demo.go:13
  - tests:
    - badge_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/badge-variants.tsx
  - implementation:
    - component: badge.go:127
    - demo: demo.go:14
  - tests:
    - badge_test.go:23

## Code Structure & Entry Points
- `badge.go`: Primary component widget layout and state logic.
- `badge_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.md`: Component interactive demo snippets and layout specs (`components/badge/demo.md`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
