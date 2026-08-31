package cfg

import (
	rcfg "github.com/kalo-build/morphe-go/pkg/registry/cfg"
)

type CompileConfig struct {
	rcfg.MorpheLoadRegistryConfig

	OutputDirPath        string
	TsTypesImportPath    string
	ZodSchemasImportPath string
	ExcludeModels        []string
}

func (c *CompileConfig) Validate() error {
	return nil
}
