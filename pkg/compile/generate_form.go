package compile

import (
	"fmt"
	"strings"

	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/naming"
)

func GenerateForm(config cfg.CompileConfig, info ModelInfo) string {
	var b strings.Builder

	typePath := fmt.Sprintf("%s/models/%s", config.TsTypesImportPath, info.KebabName)
	schemaPath := fmt.Sprintf("%s/models/%s", config.ZodSchemasImportPath, info.KebabName)

	b.WriteString("\"use client\";\n\n")
	b.WriteString("import { useForm } from \"react-hook-form\";\n")
	b.WriteString("import { zodResolver } from \"@hookform/resolvers/zod\";\n")
	b.WriteString(fmt.Sprintf("import type { %s } from \"%s\";\n", info.Name, typePath))
	b.WriteString(fmt.Sprintf("import { %sSchema } from \"%s\";\n", info.Name, schemaPath))

	b.WriteString(fmt.Sprintf("\nexport interface %sFormProps {\n", info.Name))
	b.WriteString(fmt.Sprintf("  defaultValues?: Partial<%s>;\n", info.Name))
	b.WriteString(fmt.Sprintf("  onSubmit: (data: %s) => void | Promise<void>;\n", info.Name))
	b.WriteString("  disabled?: boolean;\n")
	b.WriteString("}\n")

	b.WriteString(fmt.Sprintf("\nexport function %sForm({ defaultValues, onSubmit, disabled }: %sFormProps) {\n", info.Name, info.Name))
	b.WriteString("  const {\n")
	b.WriteString("    register,\n")
	b.WriteString("    handleSubmit,\n")
	b.WriteString("    formState: { errors, isSubmitting },\n")
	b.WriteString(fmt.Sprintf("  } = useForm<%s>({\n", info.Name))
	b.WriteString(fmt.Sprintf("    resolver: zodResolver(%sSchema),\n", info.Name))
	b.WriteString("    defaultValues,\n")
	b.WriteString("  });\n")

	b.WriteString("\n  return (\n")
	b.WriteString(fmt.Sprintf("    <form onSubmit={handleSubmit(onSubmit)} className=\"space-y-4\" data-testid=\"%s-form\">\n", info.KebabName))

	for _, field := range info.FormFields() {
		writeFormField(&b, info.KebabName, field)
	}

	b.WriteString("      <button\n")
	b.WriteString("        type=\"submit\"\n")
	b.WriteString("        data-testid=\"submit-button\"\n")
	b.WriteString("        disabled={disabled || isSubmitting}\n")
	b.WriteString("        className=\"rounded-md bg-primary px-4 py-2 text-primary-foreground\"\n")
	b.WriteString("      >\n")
	b.WriteString("        {isSubmitting ? \"Saving...\" : \"Save\"}\n")
	b.WriteString("      </button>\n")

	b.WriteString("    </form>\n")
	b.WriteString("  );\n")
	b.WriteString("}\n")

	return b.String()
}

func writeFormField(b *strings.Builder, modelKebab string, field FieldInfo) {
	if field.InputType == "checkbox" {
		writeCheckboxField(b, modelKebab, field)
		return
	}

	fieldKebab := naming.ToKebabCase(field.CamelName)
	testID := fmt.Sprintf("%s-%s-input", modelKebab, fieldKebab)

	registerOpts := ""
	if field.InputType == "number" {
		registerOpts = ", { valueAsNumber: true }"
	}

	b.WriteString("      <div>\n")
	b.WriteString(fmt.Sprintf("        <label htmlFor=\"%s\" className=\"block text-sm font-medium\">\n", field.CamelName))
	b.WriteString(fmt.Sprintf("          %s\n", field.Label))
	b.WriteString("        </label>\n")
	b.WriteString(fmt.Sprintf("        <input\n"))
	b.WriteString(fmt.Sprintf("          id=\"%s\"\n", field.CamelName))
	b.WriteString(fmt.Sprintf("          type=\"%s\"\n", field.InputType))
	b.WriteString(fmt.Sprintf("          data-testid=\"%s\"\n", testID))
	b.WriteString(fmt.Sprintf("          {...register(\"%s\"%s)}\n", field.CamelName, registerOpts))
	b.WriteString("          disabled={disabled || isSubmitting}\n")
	b.WriteString("          className=\"mt-1 block w-full rounded-md border border-input px-3 py-2\"\n")
	b.WriteString("        />\n")
	b.WriteString(fmt.Sprintf("        {errors.%s && (\n", field.CamelName))
	b.WriteString(fmt.Sprintf("          <p className=\"mt-1 text-sm text-destructive\">{errors.%s.message}</p>\n", field.CamelName))
	b.WriteString("        )}\n")
	b.WriteString("      </div>\n")
}

func writeCheckboxField(b *strings.Builder, modelKebab string, field FieldInfo) {
	fieldKebab := naming.ToKebabCase(field.CamelName)
	testID := fmt.Sprintf("%s-%s-checkbox", modelKebab, fieldKebab)

	b.WriteString("      <div className=\"flex items-center gap-2\">\n")
	b.WriteString(fmt.Sprintf("        <input\n"))
	b.WriteString(fmt.Sprintf("          id=\"%s\"\n", field.CamelName))
	b.WriteString("          type=\"checkbox\"\n")
	b.WriteString(fmt.Sprintf("          data-testid=\"%s\"\n", testID))
	b.WriteString(fmt.Sprintf("          {...register(\"%s\")}\n", field.CamelName))
	b.WriteString("          disabled={disabled || isSubmitting}\n")
	b.WriteString("          className=\"rounded border border-input\"\n")
	b.WriteString("        />\n")
	b.WriteString(fmt.Sprintf("        <label htmlFor=\"%s\" className=\"text-sm font-medium\">\n", field.CamelName))
	b.WriteString(fmt.Sprintf("          %s\n", field.Label))
	b.WriteString("        </label>\n")
	b.WriteString("      </div>\n")
}
