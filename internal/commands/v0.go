package commands

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/nguyendkn/augment-smartless/internal/core"
	"github.com/nguyendkn/augment-smartless/internal/prompts"
	"github.com/spf13/cobra"
)

// NewV0Command creates the v0 subcommand
func NewV0Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "v0",
		Short: "Execute v0 workflow using embedded prompt",
		Long:  "Execute the v0 workflow using the embedded v0 prompt content for Vercel's AI-powered assistant",
		Run:   runV0Command,
	}

	return cmd
}

func runV0Command(cmd *cobra.Command, args []string) {
	// Display ASCII art banner
	banner := figure.NewFigure("V0", "big", true)
	banner.Print()

	fmt.Println("🚀 Vercel v0 Assistant")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println("📋 AI-powered web development with Next.js and React")
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Println()

	// Get the v0 prompt
	v0Prompt := prompts.GetV0Prompt()

	// Save the prompt to .augment-guidelines
	if err := core.SaveGuidelines(v0Prompt); err != nil {
		fmt.Printf("Error saving guidelines: %v\n", err)
		return
	}

	fmt.Println("✅ v0 prompt has been saved to .augment-guidelines")
	fmt.Println("📝 The prompt content is now ready for use with your AI assistant")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Copy the content from .augment-guidelines")
	fmt.Println("  2. Paste it into your v0 conversation")
	fmt.Println("  3. Start building with Next.js and React!")
	fmt.Println()
	fmt.Println("Features included:")
	fmt.Println("  • Next.js App Router support")
	fmt.Println("  • shadcn/ui components")
	fmt.Println("  • Tailwind CSS styling")
	fmt.Println("  • TypeScript support")
	fmt.Println("  • Responsive design patterns")
}
