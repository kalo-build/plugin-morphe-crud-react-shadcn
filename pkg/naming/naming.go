package naming

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	return splitAndJoin(s, "_")
}

func ToKebabCase(s string) string {
	return splitAndJoin(s, "-")
}

func ToCamelCase(s string) string {
	words := splitPascal(s)
	if len(words) == 0 {
		return s
	}
	words[0] = strings.ToLower(words[0])
	return strings.Join(words, "")
}

func ToLabel(s string) string {
	words := splitPascal(s)
	return strings.Join(words, " ")
}

func Pluralize(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") || strings.HasSuffix(s, "z") {
		return s + "es"
	}
	if strings.HasSuffix(s, "y") && len(s) > 1 {
		prev := s[len(s)-2]
		if prev != 'a' && prev != 'e' && prev != 'i' && prev != 'o' && prev != 'u' {
			return s[:len(s)-1] + "ies"
		}
	}
	return s + "s"
}

func CollectionName(modelName string) string {
	return Pluralize(ToKebabCase(modelName))
}

func splitAndJoin(s string, sep string) string {
	words := splitPascal(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, sep)
}

func splitPascal(s string) []string {
	if s == "" {
		return nil
	}

	var words []string
	wordStart := 0
	runes := []rune(s)

	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			if unicode.IsLower(runes[i-1]) {
				words = append(words, string(runes[wordStart:i]))
				wordStart = i
			} else if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				words = append(words, string(runes[wordStart:i]))
				wordStart = i
			}
		}
	}
	words = append(words, string(runes[wordStart:]))
	return words
}
