# Label & Typography Component Specification (`gio-shadcn`)

**Official shadcn Reference:** [https://ui.shadcn.com/docs/components/label](https://ui.shadcn.com/docs/components/label)

---

## Mandatory Developer & AI Agent Instructions
Whenever adding, modifying, or refactoring the `label` component implementation or unit tests in this directory, developers and AI agents MUST:
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
- [x] **Heading 1 (H1)**
  - implementation:
    - label.go:167
  - tests:
    - label_test.go:13
- [x] **Heading 2 (H2)**
  - implementation:
    - label.go:168
  - tests:
    - label_test.go:20
- [x] **Heading 3 (H3)**
  - implementation:
    - label.go:169
  - tests:
    - label_test.go:27
- [x] **Heading 4 (H4)**
  - implementation:
    - label.go:170
  - tests:
    - label_test.go:34
- [x] **Body Paragraph (P)**
  - implementation:
    - label.go:171
  - tests:
    - label_test.go:41
- [x] **Muted & Small Text**
  - implementation:
    - label.go:172
  - tests:
    - label_test.go:48

### Capabilities & Features
- [x] **Typography Font Scale**
  - implementation:
    - label.go:197
  - tests:
    - label_test.go:56
- [x] **Theme Color Integration**
  - implementation:
    - label.go:214
  - tests:
    - label_test.go:69

---



### Demos
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/label-demo.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-h1.tsx
  - implementation:
    - component: label.go:250
    - demo: demo.go:18
  - tests:
    - label_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-h2.tsx
  - implementation:
    - component: label.go:252
    - demo: demo.go:19
  - tests:
    - label_test.go:18
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-h3.tsx
  - implementation:
    - component: label.go:254
    - demo: demo.go:20
  - tests:
    - label_test.go:25
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-h4.tsx
  - implementation:
    - component: label.go:256
    - demo: demo.go:21
  - tests:
    - label_test.go:32
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-p.tsx
  - implementation:
    - component: label.go:258
    - demo: demo.go:22
  - tests:
    - label_test.go:39
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-muted.tsx
  - implementation:
    - component: label.go:264
    - demo: demo.go:23
  - tests:
    - label_test.go:46
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-small.tsx
  - implementation:
    - component: label.go:260
    - demo: demo.go:24
  - tests:
    - label_test.go:46
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-demo.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-lead.tsx
  - implementation:
    - component: label.go:262
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-large.tsx
  - implementation:
    - component: label.go:266
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-blockquote.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-inline-code.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-list.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-table.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:53
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/label-rtl.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:11
* [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/typography-rtl.tsx
  - implementation:
    - component: label.go:197
    - demo: demo.go:17
  - tests:
    - label_test.go:11

---

## Code Structure & Entry Points
- `label.go`: Primary component widget layout and state logic.
- `label_test.go`: Automated unit test suite verifying layout dimensions and state updates.
- `demo.go`: Exported interactive demo widget (`Demo`).
- `AGENTS.md`: Component specification, shadcn reference URL, and maintainer guidelines (this file).