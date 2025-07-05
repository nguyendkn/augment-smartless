package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	GuidelinesFileName = ".augment-guidelines"
	GitignoreFileName  = ".gitignore"
)

// SaveGuidelines saves the provided content to .augment-guidelines file
func SaveGuidelines(content string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	outputPath := filepath.Join(cwd, GuidelinesFileName)

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write guidelines file: %w", err)
	}

	fmt.Printf("Guidelines saved to: %s\n", outputPath)

	// Update .gitignore
	if err := updateGitignore(cwd); err != nil {
		return fmt.Errorf("failed to update .gitignore: %w", err)
	}

	return nil
}

// updateGitignore adds .augment-guidelines to .gitignore if not present
func updateGitignore(cwd string) error {
	gitignorePath := filepath.Join(cwd, GitignoreFileName)

	// Check if .gitignore exists
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		// Create .gitignore with .augment-guidelines
		if err := os.WriteFile(gitignorePath, []byte(GuidelinesFileName+"\n"), 0644); err != nil {
			return fmt.Errorf("failed to create .gitignore: %w", err)
		}
		fmt.Println("Created .gitignore and added .augment-guidelines")
		return nil
	}

	// Read existing .gitignore
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		return fmt.Errorf("failed to read .gitignore: %w", err)
	}

	gitignoreContent := string(content)

	// Check if .augment-guidelines is already present
	if strings.Contains(gitignoreContent, GuidelinesFileName) {
		return nil // Already present
	}

	// Append .augment-guidelines to .gitignore
	file, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for appending: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString("\n" + GuidelinesFileName + "\n"); err != nil {
		return fmt.Errorf("failed to append to .gitignore: %w", err)
	}

	fmt.Println("Added .augment-guidelines to .gitignore")
	return nil
}
