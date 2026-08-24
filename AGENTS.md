---
created_on: 2026-08-22 14:45
last_modified: 2026-08-22 15:45
status: current
---

# `gio-shadcn` Package Specification & Component Governance Policy

Native Go UI component library bringing `shadcn/ui` design tokens, Flex/Grid layout semantics, and GPU-accelerated immediate-mode vector graphics to `gioui.org` (Gio).

---

## Commands
- **Run Interactive Demo:** `just demo` (or `go run ./demo/cmd`)
- **Build Demo Binary:** `just build-demo` (creates `bin/demo-app`)
- **Run Unit Tests:** `just test` (or `go test ./...`)
- **Run Static Analysis:** `just vet` (or `go vet ./...`)
- **Format Code:** `just fmt-all` (or `go fmt ./...`)

---

## 1. Component Specification & Line-Numbered Mapping Policy

Every UI component in `gio-shadcn` MUST maintain its own dedicated specification file located at:
`components/<component>/AGENTS.md`

### Mandatory Requirements for `components/<component>/AGENTS.md`:

1. **Official Reference Link:**
   Must include a direct link to the corresponding `shadcn/ui` specification (`https://ui.shadcn.com/docs/components/<slug>`).

2. **Line-Numbered Checklist Mapping:**
   Every variant, capability/feature, and official `shadcn/ui` demo example supported by the component MUST be listed in the checklist with nested sub-bullets:
   ```markdown
   ### Variants
   - [x] **Variant Name**
     - implementation:
       - <component>.go:<line_number>
     - tests:
       - <component>_test.go:<line_number>

   ### Capabilities & Features
   - [x] **Feature Name**
     - implementation:
       - <component>.go:<line_number>
     - tests:
       - <component>_test.go:<line_number>

   ### Demos
   * [ ] https://github.com/shadcn-ui/ui/blob/main/apps/v4/examples/aria/<component>-<variant>.tsx
     - implementation:
       - component: <component>.go:<line_number>
       - demo: demo.go:<line_number>
     - tests:
       - <component>_test.go:<line_number>
   ```

3. **Distinct Test Function Requirement:**
   - Every variant, feature capability, and official demo MUST have dedicated, distinct unit test assertions in `<component>_test.go`.
   - Shared or generic line numbers across different checklist items are strictly prohibited. Each checklist item must point to its own unique, manually verified implementation and test code location.

4. **Manual Documentation Policy (NO SCRIPTS):**
   - Developers and AI agents MUST NEVER use automated scripts to generate, update, or edit `.md` documentation files or `AGENTS.md` files. All documentation reviews and line number mappings MUST be performed manually line-by-line using direct editing tools.

---

## 2. GPU Rendering Safety & Color State Isolation Rules

To prevent GPU canvas color bleeding, clipping bugs, or Metal swapchain clear pass leaks:

1. **Clip Isolation:**
   - All rounded rectangle background and stroke rendering MUST go through `theme.DrawRRectBackground(gtx, rect, radius, color)` and `theme.DrawStroke(gtx, path, width, color)`.
   - Drawing calls MUST validate bounds: `if rect.Dx() <= 0 || rect.Dy() <= 0 { return }`.

2. **Radius Clamping:**
   - `clip.UniformRRect` corner radii MUST be clamped to half the component height (`heightPx / 2`). Passing unclamped `RadiusFull` (`9999px`) on small-height rectangles is strictly prohibited as it triggers vector clip mask overflow.

3. **Color State Reset:**
   - Every `Layout` method MUST end with an explicit GPU color state reset:
     ```go
     paint.ColorOp{Color: th.Colors.Background}.Add(gtx.Ops)
     ```
   - In window frame event loops (`app.FrameEvent`), the active GPU color state MUST be reset to `th.Colors.Background` immediately prior to `e.Frame(&ops)`.

---

## Boundaries

- **Always:** Automatically record all new user instructions in the appropriate `AGENTS.md` file immediately upon receipt (check with user if existing instructions conflict).
- **Always:** Any time code is changed such that results from running that code are changed, a test file must be changed as well; 90% code coverage is required (the `scripts/` folder is excluded from this rule).
- **Always:** Keep `components/<component>/AGENTS.md` line numbers, variants, features, demos, and test mappings accurate and up-to-date manually.
- **Ask first:** Modifying public component constructor APIs or changing theme token color definitions in `theme/colors.go`.
- **Never:** Use scripts to generate, update, or edit `.md` documentation or `AGENTS.md` files.
- **Never:** Use raw `paint.FillShape` with unclamped corner radii or bypass `theme.DrawRRectBackground` / `theme.DrawStroke`.

---

## Workspace References
- Component Implementations: `components/<component>/<component>.go`
- Component Specifications & Line Mappings: `components/<component>/AGENTS.md`
- Interactive 37-Component Gallery Showcase: `demo/demo.go` & `demo/cmd/main.go`
- Color Schemes & Theme Tokens: `theme/colors.go` & `theme/theme.go`
