package compile

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/kalo-build/morphe-go/pkg/registry"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/compile/cfg"
	"github.com/kalo-build/plugin-morphe-crud-react-shadcn/pkg/naming"
)

func MorpheToReactCRUD(config cfg.CompileConfig) error {
	r, rErr := registry.LoadMorpheRegistry(registry.LoadMorpheRegistryHooks{}, config.MorpheLoadRegistryConfig)
	if rErr != nil {
		return fmt.Errorf("failed to load morphe registry: %w", rErr)
	}

	excludeSet := make(map[string]bool)
	for _, name := range config.ExcludeModels {
		excludeSet[name] = true
	}

	allModels := r.GetAllModels()
	modelNames := make([]string, 0, len(allModels))
	for name := range allModels {
		modelNames = append(modelNames, name)
	}
	sort.Strings(modelNames)

	var modelInfos []ModelInfo
	for _, name := range modelNames {
		if excludeSet[name] {
			continue
		}
		model := allModels[name]
		info := ExtractModelInfo(model)
		modelInfos = append(modelInfos, info)
	}

	if len(modelInfos) == 0 {
		return fmt.Errorf("no models to generate (all excluded or none found)")
	}

	if err := os.MkdirAll(config.OutputDirPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	for _, info := range modelInfos {
		kebab := naming.ToKebabCase(info.Name)

		form := GenerateForm(config, info)
		if err := writeTsxFile(config.OutputDirPath, kebab+"-form.tsx", form); err != nil {
			return fmt.Errorf("failed to write form for %s: %w", info.Name, err)
		}

		table := GenerateTable(config, info)
		if err := writeTsxFile(config.OutputDirPath, kebab+"-table.tsx", table); err != nil {
			return fmt.Errorf("failed to write table for %s: %w", info.Name, err)
		}

		detail := GenerateDetail(config, info)
		if err := writeTsxFile(config.OutputDirPath, kebab+"-detail.tsx", detail); err != nil {
			return fmt.Errorf("failed to write detail for %s: %w", info.Name, err)
		}
	}

	return nil
}

func writeTsxFile(dirPath string, fileName string, content string) error {
	filePath := filepath.Join(dirPath, fileName)
	return os.WriteFile(filePath, []byte(content), 0644)
}
