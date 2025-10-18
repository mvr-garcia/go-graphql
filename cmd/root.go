package cmd

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use:   "courses-api",
	Short: "Courses API",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
