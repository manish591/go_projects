package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var fileNames []string

var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "Word, line, character, and byte count.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, filename := range fileNames {
			readBytes(filename)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&fileNames, "filename", "c", []string{}, "Please enter valid filenames")
}
