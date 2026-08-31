package compile

import (
	"testing"

	"github.com/kalo-build/morphe-go/pkg/yaml"
	"github.com/stretchr/testify/assert"
)

func TestExtractModelInfo_Basic(t *testing.T) {
	model := yaml.Model{
		Name: "Organization",
		Fields: map[string]yaml.ModelField{
			"ID":   {Type: "UUID"},
			"Code": {Type: "String"},
			"Name": {Type: "String"},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {Fields: []string{"ID"}},
			"code":    {Fields: []string{"Code"}},
		},
	}

	info := ExtractModelInfo(model)

	assert.Equal(t, "Organization", info.Name)
	assert.Equal(t, "organization", info.KebabName)
	assert.Equal(t, "organizations", info.CollectionName)
	assert.Equal(t, "ID", info.PrimaryIDField)
	assert.Len(t, info.Fields, 3)
	assert.Empty(t, info.Filters)
}

func TestExtractModelInfo_ForOneRelation(t *testing.T) {
	model := yaml.Model{
		Name: "Project",
		Fields: map[string]yaml.ModelField{
			"ID":   {Type: "UUID"},
			"Name": {Type: "String"},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {Fields: []string{"ID"}},
		},
		Related: map[string]yaml.ModelRelation{
			"Organization": {Type: "ForOne"},
		},
	}

	info := ExtractModelInfo(model)

	assert.Len(t, info.Filters, 1)
	assert.Equal(t, "Organization", info.Filters[0].RelationName)
	assert.Equal(t, "organizationId", info.Filters[0].CamelName)

	var fkField *FieldInfo
	for _, f := range info.Fields {
		if f.Name == "OrganizationID" {
			fkField = &f
			break
		}
	}
	assert.NotNil(t, fkField)
	assert.True(t, fkField.IsFK)
	assert.Equal(t, "organizationID", fkField.CamelName)
}

func TestExtractModelInfo_ForOnePolyRelation(t *testing.T) {
	model := yaml.Model{
		Name: "FileReference",
		Fields: map[string]yaml.ModelField{
			"ID":        {Type: "UUID"},
			"FilePath":  {Type: "String"},
			"CreatedAt": {Type: "Time"},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {Fields: []string{"ID"}},
		},
		Related: map[string]yaml.ModelRelation{
			"Anchor": {
				Type: "ForOnePoly",
				For:  []string{"BacklogItem", "Capture"},
			},
		},
	}

	info := ExtractModelInfo(model)

	var anchorID *FieldInfo
	for i := range info.Fields {
		if info.Fields[i].Name == "AnchorID" {
			anchorID = &info.Fields[i]
			break
		}
	}
	assert.NotNil(t, anchorID)
	assert.Equal(t, "anchorID", anchorID.CamelName)
}

func TestFormFields_ExcludesPrimaryID(t *testing.T) {
	model := yaml.Model{
		Name: "Organization",
		Fields: map[string]yaml.ModelField{
			"ID":   {Type: "UUID"},
			"Code": {Type: "String"},
			"Name": {Type: "String"},
		},
		Identifiers: map[string]yaml.ModelIdentifier{
			"primary": {Fields: []string{"ID"}},
		},
	}

	info := ExtractModelInfo(model)
	formFields := info.FormFields()

	for _, f := range formFields {
		assert.NotEqual(t, "ID", f.Name, "primary ID should not be in form fields")
	}
	assert.Len(t, formFields, 2)
}

func TestMorpheTypeToInputType(t *testing.T) {
	tests := []struct {
		morpheType yaml.ModelFieldType
		expected   string
	}{
		{"String", "text"},
		{"UUID", "text"},
		{"Integer", "number"},
		{"Float", "number"},
		{"Boolean", "checkbox"},
		{"Date", "date"},
		{"Time", "datetime-local"},
		{"Protected", "password"},
	}

	for _, tt := range tests {
		t.Run(string(tt.morpheType), func(t *testing.T) {
			assert.Equal(t, tt.expected, morpheTypeToInputType(tt.morpheType))
		})
	}
}
