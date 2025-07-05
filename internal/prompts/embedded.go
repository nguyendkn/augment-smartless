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
		"v0": {
			Name:        "v0",
			Description: "Vercel's AI-powered assistant for web development with Next.js and React",
			Content:     getV0Prompt(),
		},
		"lovable": {
			Name:        "lovable",
			Description: "AI editor for creating and modifying web applications with real-time preview",
			Content:     getLovablePrompt(),
		},
	}
}

// GetV0Prompt returns the embedded v0 prompt content (for backward compatibility)
func GetV0Prompt() string {
	return getV0Prompt()
}

// GetLovablePrompt returns the embedded lovable prompt content (for backward compatibility)
func GetLovablePrompt() string {
	return getLovablePrompt()
}
