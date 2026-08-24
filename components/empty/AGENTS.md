# Empty Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/empty](https://ui.shadcn.com/docs/components/empty)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `empty` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Empty State Card**
  - implementation:
    - empty.go:21
  - tests:
    - empty_test.go:13

### Capabilities & Features
- [x] **Title Header & Description Body**
  - implementation:
    - empty.go:75
  - tests:
    - empty_test.go:30
- [x] **Default Fallback Text**
  - implementation:
    - empty.go:34
  - tests:
    - empty_test.go:23
- [x] **Card Background & Border Stroke**
  - implementation:
    - empty.go:92
  - tests:
    - empty_test.go:40

---


### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-avatar-group.tsx
  - implementation:
    - component: empty.go:50
    - demo: demo.go:14
  - tests:
    - empty_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-avatar.tsx
  - implementation:
    - component: empty.go:50
    - demo: demo.go:14
  - tests:
    - empty_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-background.tsx
  - implementation:
    - component: empty.go:94
    - demo: demo.go:14
  - tests:
    - empty_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-card.tsx
  - implementation:
    - component: empty.go:50
    - demo: demo.go:14
  - tests:
    - empty_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-demo.tsx
  - implementation:
    - component: empty.go:50
    - demo: demo.go:14
  - tests:
    - empty_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-input-group.tsx
  - implementation:
    - component: empty.go:68
    - demo: demo.go:14
  - tests:
    - empty_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-outline.tsx
  - implementation:
    - component: empty.go:97
    - demo: demo.go:14
  - tests:
    - empty_test.go:37
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/empty-rtl.tsx
  - implementation:
    - component: empty.go:50
    - demo: demo.go:14
  - tests:
    - empty_test.go:11

## Code Structure & Entry Points
- `empty.go`: Primary component widget layout and state logic.
- `empty_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.md`: Component interactive demo snippets and layout specs (`components/empty/demo.md`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
