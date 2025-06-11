package cmd

func Execute() error {
	rootCmd := rootCommand()

	return rootCmd.Execute()
}
