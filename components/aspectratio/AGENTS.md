# Aspect Ratio Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/aspect-ratio](https://ui.shadcn.com/docs/components/aspect-ratio)  
**Official shadcn Source Spec (.mdx):** [https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/aspect-ratio.mdx](https://github.com/shadcn-ui/ui/blob/main/apps/v4/content/docs/components/aria/aspect-ratio.mdx)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `aspectratio` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **16:9 Standard Ratio**
  - implementation:
    - aspectratio.go:30
  - tests:
    - aspectratio_test.go:13
- [x] **4:3 Classic Ratio**
  - implementation:
    - aspectratio.go:30
  - tests:
    - aspectratio_test.go:27
- [x] **1:1 Square Ratio**
  - implementation:
    - aspectratio.go:30
  - tests:
    - aspectratio_test.go:41

### Capabilities & Features
- [x] **Proportional Constraint Layout**
  - implementation:
    - aspectratio.go:41
  - tests:
    - aspectratio_test.go:55
- [x] **Wrapped Child Fitting**
  - implementation:
    - aspectratio.go:48
  - tests:
    - aspectratio_test.go:70

### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/aspect-ratio-demo.tsx
  - implementation:
    - component: aspectratio.go:30
    - demo: demo.go:11
  - tests:
    - aspectratio_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/aspect-ratio-portrait.tsx
  - implementation:
    - component: aspectratio.go:30
    - demo: demo.go:20
  - tests:
    - aspectratio_test.go:27
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/aspect-ratio-rtl.tsx
  - implementation:
    - component: aspectratio.go:30
    - demo: demo.go:11
  - tests:
    - aspectratio_test.go:13
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/aspect-ratio-square.tsx
  - implementation:
    - component: aspectratio.go:30
    - demo: demo.go:29
  - tests:
    - aspectratio_test.go:41

---

## Code Structure & Entry Points
- `aspectratio.go`: Primary component widget layout and state logic.
- `demo.go`: Modular interactive demo component for gallery integration (`Demo`).
- `aspectratio_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).
