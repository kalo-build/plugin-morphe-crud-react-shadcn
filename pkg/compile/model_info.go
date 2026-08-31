package compile

import (
	"sort"

	"github.com/kalo-build/morphe-go/pkg/yaml"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/naming"
)

type ModelInfo struct {
	Name           string
	KebabName      string
	CollectionName string
	PrimaryIDField string
	Fields         []FieldInfo
	Filters        []FilterInfo
}

type FieldInfo struct {
	Name        string
	CamelName   string
	Label       string
	FieldType   yaml.ModelFieldType
	InputType   string
	IsOptional  bool
	IsFK        bool
	IsPrimaryID bool
}

type FilterInfo struct {
	RelationName string
	CamelName    string
	Label        string
}

func ExtractModelInfo(model yaml.Model) ModelInfo {
	info := ModelInfo{
		Name:           model.Name,
		KebabName:      naming.ToKebabCase(model.Name),
		CollectionName: naming.CollectionName(model.Name),
	}

	for idName, id := range model.Identifiers {
		if idName == "primary" && len(id.Fields) > 0 {
			info.PrimaryIDField = id.Fields[0]
		}
	}

	fieldNames := sortedKeys(model.Fields)
	for _, fieldName := range fieldNames {
		field := model.Fields[fieldName]
		isOptional := false
		for _, attr := range field.Attributes {
			if attr == "optional" {
				isOptional = true
				break
			}
		}
		info.Fields = append(info.Fields, FieldInfo{
			Name:        fieldName,
			CamelName:   naming.ToCamelCase(fieldName),
			Label:       naming.ToLabel(fieldName),
			FieldType:   field.Type,
			InputType:   morpheTypeToInputType(field.Type),
			IsOptional:  isOptional,
			IsPrimaryID: fieldName == info.PrimaryIDField,
		})
	}

	relNames := sortedKeys(model.Related)
	for _, relName := range relNames {
		rel := model.Related[relName]
		switch rel.Type {
		case "ForOne":
			fkName := relName + "ID"
			info.Fields = append(info.Fields, FieldInfo{
				Name:       fkName,
				CamelName:  naming.ToCamelCase(fkName),
				Label:      naming.ToLabel(relName) + " ID",
				InputType:  "text",
				IsOptional: true,
				IsFK:       true,
			})
			info.Filters = append(info.Filters, FilterInfo{
				RelationName: relName,
				CamelName:    naming.ToCamelCase(relName) + "Id",
				Label:        naming.ToLabel(relName),
			})
		case "ForOnePoly":
			through := relName
			if rel.Through != "" {
				through = rel.Through
			}
			info.Fields = append(info.Fields, FieldInfo{
				Name:       through + "ID",
				CamelName:  naming.ToCamelCase(through + "ID"),
				Label:      naming.ToLabel(through) + " ID",
				InputType:  "text",
				IsOptional: true,
				IsFK:       true,
			})
			info.Fields = append(info.Fields, FieldInfo{
				Name:       through + "Type",
				CamelName:  naming.ToCamelCase(through) + "Type",
				Label:      naming.ToLabel(through) + " Type",
				InputType:  "text",
				IsOptional: true,
				IsFK:       true,
			})
			info.Filters = append(info.Filters, FilterInfo{
				RelationName: through,
				CamelName:    naming.ToCamelCase(through + "ID"),
				Label:        naming.ToLabel(through),
			})
		}
	}

	return info
}

func (m ModelInfo) FormFields() []FieldInfo {
	var fields []FieldInfo
	for _, f := range m.Fields {
		if f.IsPrimaryID {
			continue
		}
		if f.FieldType == "Sealed" || f.FieldType == "AutoIncrement" {
			continue
		}
		fields = append(fields, f)
	}
	return fields
}

func morpheTypeToInputType(fieldType yaml.ModelFieldType) string {
	switch fieldType {
	case "Integer", "AutoIncrement", "Float":
		return "number"
	case "Boolean":
		return "checkbox"
	case "Time":
		return "datetime-local"
	case "Date":
		return "date"
	case "Protected":
		return "password"
	default:
		return "text"
	}
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
