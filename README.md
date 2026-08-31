# plugin-morphe-crud-react-shadcn

Generates React form, table, and detail component shells from Morphe model definitions. **Thin template layer** — all HTTP, types, validation, and casing conversion are handled by upstream plugins.

## data-testid (Playwright / .up alignment)

Generated components include `data-testid` attributes using a fixed naming convention so they work with **plugin-uicheck-playwright** and `.up` UI Page specs. When writing `.up` element keys, use the same names so generated Playwright tests match. See [docs/TESTID_CONVENTION.md](docs/TESTID_CONVENTION.md) for the full convention.

## What it generates

Given Morphe model files (`.mod`), this plugin generates per model:

1. **Form component** (`{model}-form.tsx`) — `react-hook-form` + `zodResolver` form with fields derived from the model, Zod validation, and an `onSubmit` callback
2. **Table component** (`{model}-table.tsx`) — typed table rendering all model fields with row click support
3. **Detail component** (`{model}-detail.tsx`) — typed read-only field display

## Derivation rules

Uses the same model analysis as `plugin-morphe-crud-go-gin`:

- **Primary identifier** (`identifiers.primary`) — excluded from form fields (auto-generated or externally managed)
- **Secondary identifiers** (e.g., `code`) — included as regular form/display fields
- **ForOne relationships** — synthesizes FK fields (e.g., `OrganizationID`) included in forms, tables, and detail views; also derives filter parameters
- **ForOnePoly relationships** — synthesizes both ID and Type FK fields
- **Field types** — mapped to HTML input types (`String`→`text`, `Integer`→`number`, `Boolean`→`checkbox`, `Date`→`date`, `Time`→`datetime-local`, `Protected`→`password`)
- **Sealed/AutoIncrement fields** — excluded from forms

## What remains hand-written

1. **Data fetching** — use `plugin-openapi-sdk-ts` Client to fetch data, pass as props
2. **Form submission** — wire `onSubmit` to SDK client calls with bridge converters
3. **Routing/navigation** — React Router, Next.js, etc.
4. **Layout/page composition** — combine generated components into pages
5. **State management** — React Query, SWR, Zustand, etc.
6. **Error handling UI** — beyond basic form validation errors

## Configuration

```yaml
config:
  "@kalo-build/plugin-morphe-crud-react-shadcn":
    tsTypesImportPath: "@/generated/types"
    zodSchemasImportPath: "@/generated/schemas"
    excludeModels:
      - User
```

| Key | Type | Required | Description |
|-----|------|----------|-------------|
| `tsTypesImportPath` | string | yes | Base import path for TypeScript interfaces from `plugin-morphe-ts-types` |
| `zodSchemasImportPath` | string | yes | Base import path for Zod schemas from `plugin-morphe-zod-types` |
| `excludeModels` | string[] | no | Model names to skip |

Import paths are resolved as `{basePath}/models/{kebab-name}` in the generated components.

## Pipeline context

This plugin reads Morphe models directly (single input: `KA:MO1:YAML1`). The upstream plugin artifacts are referenced via configurable import paths — the plugin does **not** need to read those files at generation time.

```
Morphe Registry (models)
    │
    ├──► plugin-morphe-ts-types ──► TypeScript interfaces
    ├──► plugin-morphe-zod-types ──► Zod schemas
    ├──► plugin-morphe-ts-zod-bridge ──► wire ↔ app converters
    ├──► plugin-morphe-crud-go-gin ──► Go API handlers
    │       └──► plugin-openapi-sdk-ts ──► typed HTTP client
    │
    └──► plugin-morphe-crud-react-shadcn (this plugin)
              imports from ▲ via configurable paths
```

## kalo.yaml example

```yaml
stores:
  - name: "KA_MO_YAML"
    type: "localFileSystem"
    path: "./morphe"
  - name: "KA_REACT_COMPONENTS"
    type: "localFileSystem"
    path: "./src/generated/components"

plugins:
  "@kalo-build/plugin-morphe-crud-react-shadcn":
    version: "v1.0.0"
    input:
      format: "KA:MO1:YAML1"
      store: "KA_MO_YAML"
    output:
      format: "KA:MO1:REACT1"
      store: "KA_REACT_COMPONENTS"
    config:
      tsTypesImportPath: "@/generated/types"
      zodSchemasImportPath: "@/generated/schemas"
```

## Building

```bash
# Run tests
go test -v ./...

# Build WASI binary
GOOS=wasip1 GOARCH=wasm go build -o dist/plugin.wasm ./cmd/plugin/

# Build native (for local testing)
go build -o dist/plugin ./cmd/plugin/
```

## Project structure

```
├── cmd/plugin/main.go              # WASM entrypoint
├── pkg/
│   ├── compile/
│   │   ├── cfg/config.go           # Configuration structs
│   │   ├── compile.go              # Main pipeline
│   │   ├── model_info.go           # Model analysis + derivation
│   │   ├── generate_form.go        # Form component generation
│   │   ├── generate_table.go       # Table component generation
│   │   └── generate_detail.go      # Detail component generation
│   └── naming/naming.go            # Casing utilities
├── testdata/
│   ├── registry/minimal/models/    # Test Morphe models
│   └── ground-truth/compile-minimal/  # Expected generated output
├── plugin.yaml                     # Plugin manifest
└── README.md
```
