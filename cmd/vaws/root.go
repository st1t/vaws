package vaws

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vaws",
	Short:   "The vaws command was created to simplify the display of AWS resources.",
	Long:    `The vaws command was created to simplify the display of AWS resources.`,
	Version: "0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("aws-profile", "p", "", "-p my-aws")
	rootCmd.PersistentFlags().IntP("sort-position", "s", 1, "-s 1")
}
