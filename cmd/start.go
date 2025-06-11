package cmd

import (
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Short: "Start the service",
	}

	start.AddCommand(ingestorCommand())
	start.AddCommand(consumerCommand())

	return start
}
