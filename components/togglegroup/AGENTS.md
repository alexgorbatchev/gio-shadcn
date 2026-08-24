# Toggle Group Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/toggle-group](https://ui.shadcn.com/docs/components/toggle-group)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `togglegroup` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Toggle Group Segmented Buttons**
  - implementation:
    - togglegroup.go:69
  - tests:
    - togglegroup_test.go:13

### Capabilities & Features
- [x] **Single Key Selection**
  - implementation:
    - togglegroup.go:78
  - tests:
    - togglegroup_test.go:21
- [x] **Segmented Button Styling**
  - implementation:
    - togglegroup.go:121
  - tests:
    - togglegroup_test.go:30
- [x] **Active Highlight**
  - implementation:
    - togglegroup.go:126
  - tests:
    - togglegroup_test.go:21

---



### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-demo.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-disabled.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-demo.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:12
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-disabled.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-font-weight-selector.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-outline.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-sizes.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-spacing.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-group-vertical.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-outline.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-sizes.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/toggle-text.tsx
  - implementation:
    - component: togglegroup.go:69
    - demo: demo.go:16
  - tests:
    - togglegroup_test.go:28

## Code Structure & Entry Points
- `togglegroup.go`: Primary component widget layout and state logic.
- `togglegroup_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Interactive component gallery demo.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).