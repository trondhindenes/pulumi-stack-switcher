package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/trond/pulumi-stack-switcher/internal/stacks"
)

var version = "dev"

// SetVersion sets the version string for the CLI
func SetVersion(v string) {
	version = v
	rootCmd.Version = v
}

var rootCmd = &cobra.Command{
	Use:   "pulumi-stack-switcher [stack-name]",
	Short: "Switch between Pulumi stacks with ease",
	Long: `A lightweight CLI tool that makes it easy to switch between Pulumi stacks.

It automatically detects available stacks by scanning for Pulumi.<stack-name>.yaml
files in the current directory and provides shell completion for quick switching.

Examples:
  pulumi-stack-switcher dev        # Switch to the 'dev' stack
  pulumi-stack-switcher production # Switch to the 'production' stack
  pulumi-stack-switcher            # List available stacks`,
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: completeStacks,
	RunE:              run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	availableStacks, err := stacks.DetectInCurrentDir()
	if err != nil {
		return fmt.Errorf("failed to detect stacks: %w", err)
	}

	if len(availableStacks) == 0 {
		return fmt.Errorf("no Pulumi stacks found in current directory (looking for Pulumi.*.yaml files)")
	}

	// If no stack specified, list available stacks
	if len(args) == 0 {
		fmt.Println("Available stacks:")
		for _, s := range availableStacks {
			fmt.Printf("  %s\n", s)
		}
		return nil
	}

	stackName := args[0]

	// Verify the stack exists
	found := false
	for _, s := range availableStacks {
		if s == stackName {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("stack '%s' not found. Available stacks: %v", stackName, availableStacks)
	}

	// Execute pulumi stack select
	pulumiCmd := exec.Command("pulumi", "stack", "select", stackName)
	pulumiCmd.Stdout = os.Stdout
	pulumiCmd.Stderr = os.Stderr
	pulumiCmd.Stdin = os.Stdin

	if err := pulumiCmd.Run(); err != nil {
		return fmt.Errorf("failed to switch stack: %w", err)
	}

	fmt.Printf("Switched to stack '%s'\n", stackName)
	return nil
}

// completeStacks provides shell completion for stack names
func completeStacks(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Only complete the first argument
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	availableStacks, err := stacks.DetectInCurrentDir()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	// Filter stacks by prefix if user has typed something
	filtered := stacks.FilterStacks(availableStacks, toComplete)

	return filtered, cobra.ShellCompDirectiveNoFileComp
}
