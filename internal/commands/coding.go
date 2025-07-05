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
		Use:   "coding",
		Short: "Execute coding workflow using embedded prompt",
		Long:  "Execute the coding workflow using the embedded coding.md prompt content",
		Run:   runCodingCommand,
	}

	return cmd
}

func runCodingCommand(cmd *cobra.Command, args []string) {
	// Display ASCII art banner
	banner := figure.NewFigure("CODING", "big", true)
	banner.Print()

	fmt.Println("🚀 Augment Coding Assistant")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println("📋 Systematic Development Workflow Assistant")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println()

	// Get the coding prompt
	codingPrompt := prompts.GetCodingPrompt()

	// Save the prompt to .augment-guidelines
	if err := core.SaveGuidelines(codingPrompt); err != nil {
		fmt.Printf("Error saving guidelines: %v\n", err)
		return
	}

	fmt.Println("✅ Coding guidelines have been saved to .augment-guidelines")
	fmt.Println("📝 The prompt content is now ready for use with your AI assistant")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Copy the content from .augment-guidelines")
	fmt.Println("  2. Paste it into your AI assistant conversation")
	fmt.Println("  3. Start your coding workflow!")
}
