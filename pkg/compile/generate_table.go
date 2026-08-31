package compile

import (
	"fmt"
	"strings"

	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
)

func GenerateTable(config cfg.CompileConfig, info ModelInfo) string {
	var b strings.Builder

	typePath := fmt.Sprintf("%s/models/%s", config.TsTypesImportPath, info.KebabName)

	b.WriteString(fmt.Sprintf("import type { %s } from \"%s\";\n", info.Name, typePath))

	b.WriteString(fmt.Sprintf("\nexport interface %sTableProps {\n", info.Name))
	b.WriteString(fmt.Sprintf("  data: %s[];\n", info.Name))
	b.WriteString(fmt.Sprintf("  onRowClick?: (item: %s) => void;\n", info.Name))
	b.WriteString("}\n")

	b.WriteString(fmt.Sprintf("\nexport function %sTable({ data, onRowClick }: %sTableProps) {\n", info.Name, info.Name))
	b.WriteString("  return (\n")
	b.WriteString(fmt.Sprintf("    <table className=\"w-full border-collapse\" data-testid=\"%s-table\">\n", info.CollectionName))
	b.WriteString("      <thead>\n")
	b.WriteString("        <tr className=\"border-b\">\n")

	for _, field := range info.Fields {
		b.WriteString(fmt.Sprintf("          <th className=\"px-4 py-2 text-left text-sm font-medium\">%s</th>\n", field.Label))
	}

	b.WriteString("        </tr>\n")
	b.WriteString("      </thead>\n")
	b.WriteString("      <tbody>\n")
	b.WriteString("        {data.map((item) => (\n")

	primaryKey := "id"
	for _, f := range info.Fields {
		if f.IsPrimaryID {
			primaryKey = f.CamelName
			break
		}
	}

	b.WriteString(fmt.Sprintf("          <tr\n"))
	b.WriteString(fmt.Sprintf("            key={item.%s}\n", primaryKey))
	b.WriteString(fmt.Sprintf("            data-testid=\"%s-row\"\n", info.KebabName))
	b.WriteString("            onClick={() => onRowClick?.(item)}\n")
	b.WriteString("            className=\"border-b hover:bg-muted/50 cursor-pointer\"\n")
	b.WriteString("          >\n")

	for _, field := range info.Fields {
		if field.InputType == "checkbox" {
			b.WriteString(fmt.Sprintf("            <td className=\"px-4 py-2 text-sm\">{item.%s ? \"Yes\" : \"No\"}</td>\n", field.CamelName))
		} else if field.InputType == "date" || field.InputType == "datetime-local" {
			b.WriteString(fmt.Sprintf("            <td className=\"px-4 py-2 text-sm\">{String(item.%s)}</td>\n", field.CamelName))
		} else {
			b.WriteString(fmt.Sprintf("            <td className=\"px-4 py-2 text-sm\">{item.%s}</td>\n", field.CamelName))
		}
	}

	b.WriteString("          </tr>\n")
	b.WriteString("        ))}\n")
	b.WriteString("      </tbody>\n")
	b.WriteString("    </table>\n")
	b.WriteString("  );\n")
	b.WriteString("}\n")

	return b.String()
}
