package commands

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/nguyendkn/augment-smartless/internal/core"
	"github.com/nguyendkn/augment-smartless/internal/prompts"
	"github.com/spf13/cobra"
)

// NewCodingCommand creates the coding subcommand
func NewCodingCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "coding [prompt-name]",
		Short: "Execute coding workflow using embedded prompts",
		Long:  "Execute the coding workflow using embedded prompt content (v0, lovable)",
		Args:  cobra.MaximumNArgs(1),
		Run:   runCodingCommand,
	}

	return cmd
}

func runCodingCommand(cmd *cobra.Command, args []string) {
	// Determine which prompt to use
	promptName := "v0" // default
	if len(args) > 0 {
		promptName = args[0]
	}

	// Display ASCII art banner
	banner := figure.NewFigure("PROMPTS", "big", true)
	banner.Print()

	fmt.Println("🚀 Augment Prompt Assistant")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Printf("📋 Loading %s prompt\n", promptName)
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println()

	// Get the specified prompt
	promptContent, err := prompts.GetPrompt(promptName)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		fmt.Println()
		fmt.Println("Available prompts:")
		availablePrompts := prompts.GetAvailablePrompts()
		for name, info := range availablePrompts {
			fmt.Printf("  • %s - %s\n", name, info.Description)
		}
		return
	}

	// Save the prompt to .augment-guidelines
	if err := core.SaveGuidelines(promptContent); err != nil {
		fmt.Printf("Error saving guidelines: %v\n", err)
		return
	}

	fmt.Printf("✅ %s prompt has been saved to .augment-guidelines\n", promptName)
	fmt.Println("📝 The prompt content is now ready for use with your AI assistant")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Copy the content from .augment-guidelines")
	fmt.Println("  2. Paste it into your AI assistant conversation")
	fmt.Println("  3. Start your coding workflow!")
}
