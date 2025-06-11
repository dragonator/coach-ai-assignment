package cmd

import (
	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "coach-ai",
		Short: "coach-ai is a transaction ingestion service",
	}

	root.AddCommand(startCommand())

	return root
}
