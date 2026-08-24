# Select Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/select](https://ui.shadcn.com/docs/components/select)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `select` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Select Dropdown Field**
  - implementation:
    - select.go:57
  - tests:
    - select_test.go:13

### Capabilities & Features
- [x] **Option Items List**
  - implementation:
    - select.go:153
  - tests:
    - select_test.go:15
- [x] **Active Selected Value**
  - implementation:
    - select.go:183
  - tests:
    - select_test.go:21
- [x] **Open/Close Overlay Dropdown**
  - implementation:
    - select.go:79
  - tests:
    - select_test.go:34

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-autocomplete.tsx
  - implementation:
    - component: select.go:74
    - demo: demo.go:13
  - tests:
    - select_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-demo.tsx
  - implementation:
    - component: select.go:74
    - demo: demo.go:13
  - tests:
    - select_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-disabled.tsx
  - implementation:
    - component: select.go:74
    - demo: demo.go:13
  - tests:
    - select_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-groups.tsx
  - implementation:
    - component: select.go:148
    - demo: demo.go:13
  - tests:
    - select_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-invalid.tsx
  - implementation:
    - component: select.go:74
    - demo: demo.go:13
  - tests:
    - select_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/select-scrollable.tsx
  - implementation:
    - component: select.go:148
    - demo: demo.go:13
  - tests:
    - select_test.go:27

---

## Code Structure & Entry Points
- `select.go`: Primary component implementation.
- `select_test.go`: Unit test suite.
- `demo.go`: Modular component interactive demo layout (`components/select/demo.go`).
- `AGENTS.md`: Component specification.