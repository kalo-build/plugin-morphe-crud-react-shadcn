package compile

import (
	"fmt"
	"strings"

	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/naming"
)

func GenerateDetail(config cfg.CompileConfig, info ModelInfo) string {
	var b strings.Builder

	typePath := fmt.Sprintf("%s/models/%s", config.TsTypesImportPath, info.KebabName)

	b.WriteString(fmt.Sprintf("import type { %s } from \"%s\";\n", info.Name, typePath))

	b.WriteString(fmt.Sprintf("\nexport interface %sDetailProps {\n", info.Name))
	b.WriteString(fmt.Sprintf("  data: %s;\n", info.Name))
	b.WriteString("}\n")

	b.WriteString(fmt.Sprintf("\nexport function %sDetail({ data }: %sDetailProps) {\n", info.Name, info.Name))
	b.WriteString("  return (\n")
	b.WriteString(fmt.Sprintf("    <dl className=\"space-y-4\" data-testid=\"%s-detail\">\n", info.KebabName))

	for _, field := range info.Fields {
		fieldKebab := naming.ToKebabCase(field.CamelName)
		testID := fmt.Sprintf("%s-%s", info.KebabName, fieldKebab)
		b.WriteString("      <div>\n")
		b.WriteString(fmt.Sprintf("        <dt className=\"text-sm font-medium text-muted-foreground\">%s</dt>\n", field.Label))
		if field.InputType == "checkbox" {
			b.WriteString(fmt.Sprintf("        <dd className=\"text-sm\" data-testid=\"%s\">{data.%s ? \"Yes\" : \"No\"}</dd>\n", testID, field.CamelName))
		} else if field.InputType == "date" || field.InputType == "datetime-local" {
			b.WriteString(fmt.Sprintf("        <dd className=\"text-sm\" data-testid=\"%s\">{String(data.%s)}</dd>\n", testID, field.CamelName))
		} else {
			b.WriteString(fmt.Sprintf("        <dd className=\"text-sm\" data-testid=\"%s\">{data.%s}</dd>\n", testID, field.CamelName))
		}
		b.WriteString("      </div>\n")
	}

	b.WriteString("    </dl>\n")
	b.WriteString("  );\n")
	b.WriteString("}\n")

	return b.String()
}
