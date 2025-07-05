package prompts

import (
	"fmt"
)

// PromptInfo represents information about an available prompt
type PromptInfo struct {
	Name        string
	Description string
	Content     string
}

// GetPrompt returns the content of a specific prompt by name
func GetPrompt(name string) (string, error) {
	prompts := GetAvailablePrompts()
	if prompt, exists := prompts[name]; exists {
		return prompt.Content, nil
	}
	return "", fmt.Errorf("prompt '%s' not found", name)
}

// GetAvailablePrompts returns a map of all available prompts
func GetAvailablePrompts() map[string]PromptInfo {
	return map[string]PromptInfo{
		"coding": {
			Name:        "coding",
			Description: "Software development assistant prompt with systematic research-driven development approaches",
			Content:     getCodingPrompt(),
		},
		"coder": {
			Name:        "coder",
			Description: "Alternative coding prompt (alias for coding)",
			Content:     getCodingPrompt(),
		},
		"designer": {
			Name:        "designer",
			Description: "UI/UX design assistant prompt",
			Content:     getDesignerPrompt(),
		},
	}
}

// GetCodingPrompt returns the embedded coding prompt content (for backward compatibility)
func GetCodingPrompt() string {
	return getCodingPrompt()
}
