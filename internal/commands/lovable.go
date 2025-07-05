package commands

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/nguyendkn/augment-smartless/internal/core"
	"github.com/nguyendkn/augment-smartless/internal/prompts"
	"github.com/spf13/cobra"
)

// NewLovableCommand creates the lovable subcommand
func NewLovableCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lovable",
		Short: "Execute Lovable workflow using embedded prompt",
		Long:  "Execute the Lovable workflow using the embedded Lovable prompt content for AI web application editor",
		Run:   runLovableCommand,
	}

	return cmd
}

func runLovableCommand(cmd *cobra.Command, args []string) {
	// Display ASCII art banner
	banner := figure.NewFigure("LOVABLE", "big", true)
	banner.Print()

	fmt.Println("🚀 Lovable AI Editor")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println("📋 AI editor for creating and modifying web applications")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println()

	// Get the lovable prompt
	lovablePrompt := prompts.GetLovablePrompt()

	// Save the prompt to .augment-guidelines
	if err := core.SaveGuidelines(lovablePrompt); err != nil {
		fmt.Printf("Error saving guidelines: %v\n", err)
		return
	}

	fmt.Println("✅ Lovable prompt has been saved to .augment-guidelines")
	fmt.Println("📝 The prompt content is now ready for use with your AI assistant")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Copy the content from .augment-guidelines")
	fmt.Println("  2. Paste it into your Lovable conversation")
	fmt.Println("  3. Start building web applications!")
	fmt.Println()
	fmt.Println("Features included:")
	fmt.Println("  • Real-time code editing")
	fmt.Println("  • React/TypeScript development")
	fmt.Println("  • shadcn/ui components")
	fmt.Println("  • Tailwind CSS styling")
	fmt.Println("  • Live preview functionality")
	fmt.Println("  • File operations (create, rename, delete)")
	fmt.Println("  • Dependency management")
}
