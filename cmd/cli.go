package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/nguyendkn/augment-smartless/internal/commands"
	"github.com/nguyendkn/augment-smartless/internal/core"
	"github.com/nguyendkn/augment-smartless/internal/prompts"
	"github.com/spf13/cobra"
)

func main() {
	// Handle special case for version flag before Cobra processing
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("augment version 1.0.0")
		os.Exit(0)
	}

	// Handle direct prompt access (using embedded prompts)
	if len(os.Args) == 2 && !isSubcommand(os.Args[1]) && os.Args[1] != "--help" && os.Args[1] != "-h" {
		// Display ASCII art banner
		banner := figure.NewFigure("AUGMENT CODE CLI", "big", true)
		banner.Print()
		fmt.Println("🔧 AUGMENT CODE CLI v1.0.0 - Development Workflow Assistant")
		fmt.Println("════════════════════════════════════════════════════")
		fmt.Println()

		promptName := os.Args[1]
		fmt.Printf("Loading embedded prompt: %s\n", promptName)

		// Get the prompt content from embedded prompts
		promptContent, err := prompts.GetPrompt(promptName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Println("\nAvailable prompts:")
			availablePrompts := prompts.GetAvailablePrompts()
			for name, info := range availablePrompts {
				fmt.Printf("  - %s: %s\n", name, info.Description)
			}
			os.Exit(1)
		}

		// Save the prompt content
		if err := core.SaveGuidelines(promptContent); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving guidelines: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ %s prompt saved to .augment-guidelines\n", promptName)
		return
	}

	rootCmd := &cobra.Command{
		Use:   "augment",
		Short: "Augment CLI - Global command-line interface",
		Long: `AUGMENT CODE CLI v1.0.0 - A powerful command-line interface for development workflows

Usage: augment [options] [prompt_name]

Options:
  -h, --help     Show this help message
  --version      Show version information

Examples:
  augment --help     Show this help
  augment v0         Load embedded v0 prompt to .augment-guidelines
  augment lovable    Load embedded lovable prompt to .augment-guidelines

Available embedded prompts:
  - v0: Vercel's AI-powered assistant for web development with Next.js and React
  - lovable: AI editor for creating and modifying web applications with real-time preview`,
		Run: func(cmd *cobra.Command, args []string) {
			// Display ASCII art banner
			banner := figure.NewFigure("AUGMENT CODE CLI", "big", true)
			banner.Print()
			fmt.Println("🔧 AUGMENT CODE CLI v1.0.0 - Development Workflow Assistant")
			fmt.Println("════════════════════════════════════════════════════")
			fmt.Println()
			fmt.Println("No arguments provided. Try: augment --help")
		},
	}

	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	// Add subcommands
	rootCmd.AddCommand(commands.NewV0Command())
	rootCmd.AddCommand(commands.NewLovableCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// isSubcommand checks if the given argument is a known subcommand
func isSubcommand(arg string) bool {
	subcommands := []string{"v0", "lovable", "help", "completion"}
	for _, cmd := range subcommands {
		if arg == cmd {
			return true
		}
	}
	return false
}
