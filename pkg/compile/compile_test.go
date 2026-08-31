package compile_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	rcfg "github.com/kalo-build/morphe-go/pkg/registry/cfg"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
}

func minimalRegistryConfig(registryPath string) rcfg.MorpheLoadRegistryConfig {
	return rcfg.MorpheLoadRegistryConfig{
		RegistryEnumsDirPath:      filepath.Join(registryPath, "enums"),
		RegistryModelsDirPath:     filepath.Join(registryPath, "models"),
		RegistryStructuresDirPath: filepath.Join(registryPath, "structures"),
		RegistryEntitiesDirPath:   filepath.Join(registryPath, "entities"),
	}
}

func TestMorpheToReactCRUD(t *testing.T) {
	registryPath := filepath.Join(testdataDir(), "registry", "minimal")
	outputDir := t.TempDir()

	config := cfg.CompileConfig{
		MorpheLoadRegistryConfig: minimalRegistryConfig(registryPath),
		OutputDirPath:            outputDir,
		TsTypesImportPath:        "@/generated/types",
		ZodSchemasImportPath:     "@/generated/schemas",
	}

	err := compile.MorpheToReactCRUD(config)
	require.NoError(t, err)

	expectedFiles := []string{
		"organization-form.tsx",
		"organization-table.tsx",
		"organization-detail.tsx",
		"project-form.tsx",
		"project-table.tsx",
		"project-detail.tsx",
		"task-form.tsx",
		"task-table.tsx",
		"task-detail.tsx",
	}

	for _, f := range expectedFiles {
		path := filepath.Join(outputDir, f)
		_, err := os.Stat(path)
		assert.NoError(t, err, "expected file %s to exist", f)
	}

	formContent, err := os.ReadFile(filepath.Join(outputDir, "organization-form.tsx"))
	require.NoError(t, err)
	formStr := string(formContent)

	assert.Contains(t, formStr, "OrganizationForm")
	assert.Contains(t, formStr, "useForm")
	assert.Contains(t, formStr, "zodResolver")
	assert.Contains(t, formStr, "OrganizationSchema")
	assert.Contains(t, formStr, "@/generated/types/models/organization")
	assert.Contains(t, formStr, "@/generated/schemas/models/organization")
	assert.Contains(t, formStr, "register(\"code\")")
	assert.Contains(t, formStr, "register(\"name\")")
	assert.NotContains(t, formStr, "register(\"id\")")

	tableContent, err := os.ReadFile(filepath.Join(outputDir, "project-table.tsx"))
	require.NoError(t, err)
	tableStr := string(tableContent)

	assert.Contains(t, tableStr, "ProjectTable")
	assert.Contains(t, tableStr, "onRowClick")
	assert.Contains(t, tableStr, "item.code")
	assert.Contains(t, tableStr, "Organization ID")

	detailContent, err := os.ReadFile(filepath.Join(outputDir, "task-detail.tsx"))
	require.NoError(t, err)
	detailStr := string(detailContent)

	assert.Contains(t, detailStr, "TaskDetail")
	assert.Contains(t, detailStr, "data.title")
	assert.Contains(t, detailStr, "data.status")
	assert.Contains(t, detailStr, "Project ID")
}

func TestMorpheToReactCRUD_ExcludeModels(t *testing.T) {
	registryPath := filepath.Join(testdataDir(), "registry", "minimal")
	outputDir := t.TempDir()

	config := cfg.CompileConfig{
		MorpheLoadRegistryConfig: minimalRegistryConfig(registryPath),
		OutputDirPath:            outputDir,
		TsTypesImportPath:        "@/generated/types",
		ZodSchemasImportPath:     "@/generated/schemas",
		ExcludeModels:            []string{"Task"},
	}

	err := compile.MorpheToReactCRUD(config)
	require.NoError(t, err)

	_, err = os.Stat(filepath.Join(outputDir, "organization-form.tsx"))
	assert.NoError(t, err)

	_, err = os.Stat(filepath.Join(outputDir, "task-form.tsx"))
	assert.True(t, os.IsNotExist(err), "task components should not be generated when excluded")
}
