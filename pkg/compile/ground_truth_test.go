package compile_test

import (
	"path/filepath"
	"runtime"
	"testing"

	rcfg "github.com/kalo-build/morphe-go/pkg/registry/cfg"
	"github.com/kalo-build/go-util/assertfile"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
	"github.com/stretchr/testify/suite"
)

type GroundTruthTestSuite struct {
	assertfile.FileSuite

	TestDataPath    string
	GroundTruthPath string
}

func TestGroundTruthTestSuite(t *testing.T) {
	suite.Run(t, new(GroundTruthTestSuite))
}

func (suite *GroundTruthTestSuite) SetupTest() {
	_, filename, _, _ := runtime.Caller(0)
	suite.TestDataPath = filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
	suite.GroundTruthPath = filepath.Join(suite.TestDataPath, "ground-truth", "compile-minimal")
}

func (suite *GroundTruthTestSuite) TestGenerateAll() {
	registryPath := filepath.Join(suite.TestDataPath, "registry", "minimal")
	outputDir := suite.T().TempDir()

	config := cfg.CompileConfig{
		MorpheLoadRegistryConfig: rcfg.MorpheLoadRegistryConfig{
			RegistryEnumsDirPath:      filepath.Join(registryPath, "enums"),
			RegistryModelsDirPath:     filepath.Join(registryPath, "models"),
			RegistryStructuresDirPath: filepath.Join(registryPath, "structures"),
			RegistryEntitiesDirPath:   filepath.Join(registryPath, "entities"),
		},
		OutputDirPath:        outputDir,
		TsTypesImportPath:    "@/generated/types",
		ZodSchemasImportPath: "@/generated/schemas",
	}

	err := compile.MorpheToReactCRUD(config)
	suite.NoError(err)

	files := []string{
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

	for _, f := range files {
		actualPath := filepath.Join(outputDir, f)
		expectedPath := filepath.Join(suite.GroundTruthPath, f)
		suite.FileExists(actualPath)
		suite.FileEquals(actualPath, expectedPath)
	}
}
