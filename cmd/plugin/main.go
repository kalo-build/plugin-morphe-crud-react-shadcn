package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	rcfg "github.com/kalo-build/morphe-go/pkg/registry/cfg"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
)

type StoreConfig struct {
	ID        uint32 `json:"id"`
	Type      string `json:"type"`
	MountPath string `json:"mountPath,omitempty"`
}

type PluginConfig struct {
	Stores     map[string]StoreConfig `json:"stores,omitempty"`
	InputPath  string                 `json:"inputPath,omitempty"`
	OutputPath string                 `json:"outputPath,omitempty"`
	Config     PluginConfigFields     `json:"config"`
	Verbose    bool                   `json:"verbose,omitempty"`
}

type PluginConfigFields struct {
	TsTypesImportPath    string   `json:"tsTypesImportPath"`
	ZodSchemasImportPath string   `json:"zodSchemasImportPath"`
	ExcludeModels        []string `json:"excludeModels,omitempty"`
}

const (
	ErrMissingConfig      = 3
	ErrInvalidConfig      = 4
	ErrInputPathRequired  = 12
	ErrOutputPathRequired = 13
	ErrCompileFailed      = 1
)

func logInfo(verbose bool, format string, args ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stdout, format+"\n", args...)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: plugin-morphe-crud-react-shadcn <config>")
		os.Exit(ErrMissingConfig)
	}

	rawConfig := os.Args[1]
	var pluginConfig PluginConfig
	if err := json.Unmarshal([]byte(rawConfig), &pluginConfig); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing config JSON:", err)
		os.Exit(ErrInvalidConfig)
	}

	var inputPath, outputPath string

	if pluginConfig.Stores != nil {
		for _, store := range pluginConfig.Stores {
			switch store.MountPath {
			case "/input":
				inputPath = "/input"
			case "/output":
				outputPath = "/output"
			}
		}
	}

	if inputPath == "" && pluginConfig.InputPath != "" {
		inputPath = pluginConfig.InputPath
	}
	if outputPath == "" && pluginConfig.OutputPath != "" {
		outputPath = pluginConfig.OutputPath
	}

	if inputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: Input path is required (Morphe registry directory)")
		os.Exit(ErrInputPathRequired)
	}
	if outputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: Output path is required")
		os.Exit(ErrOutputPathRequired)
	}

	if inputPath[0] != '/' {
		if abs, err := filepath.Abs(inputPath); err == nil {
			inputPath = abs
		}
	}
	if outputPath[0] != '/' {
		if abs, err := filepath.Abs(outputPath); err == nil {
			outputPath = abs
		}
	}

	if pluginConfig.Config.TsTypesImportPath == "" {
		fmt.Fprintln(os.Stderr, "Error: tsTypesImportPath is required in config")
		os.Exit(ErrInvalidConfig)
	}
	if pluginConfig.Config.ZodSchemasImportPath == "" {
		fmt.Fprintln(os.Stderr, "Error: zodSchemasImportPath is required in config")
		os.Exit(ErrInvalidConfig)
	}

	compileConfig := cfg.CompileConfig{
		MorpheLoadRegistryConfig: rcfg.MorpheLoadRegistryConfig{
			RegistryEnumsDirPath:      filepath.Join(inputPath, "enums"),
			RegistryModelsDirPath:     filepath.Join(inputPath, "models"),
			RegistryStructuresDirPath: filepath.Join(inputPath, "structures"),
			RegistryEntitiesDirPath:   filepath.Join(inputPath, "entities"),
		},
		OutputDirPath:        outputPath,
		TsTypesImportPath:    pluginConfig.Config.TsTypesImportPath,
		ZodSchemasImportPath: pluginConfig.Config.ZodSchemasImportPath,
		ExcludeModels:        pluginConfig.Config.ExcludeModels,
	}

	logInfo(pluginConfig.Verbose, "Reading Morphe registry from: '%s'", inputPath)
	logInfo(pluginConfig.Verbose, "Generating React components to: '%s'", outputPath)

	if err := compile.MorpheToReactCRUD(compileConfig); err != nil {
		fmt.Fprintln(os.Stderr, "Compilation failed:", err)
		os.Exit(ErrCompileFailed)
	}

	logInfo(pluginConfig.Verbose, "React components generated successfully")
	os.Exit(0)
}
