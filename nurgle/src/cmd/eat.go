package cmd

import (
	"github.com/spf13/cobra"
)

var eatCmd = &cobra.Command{
	Use:   "eat",
	Short: "Consume system resources.",
	Long:  `Consume system resources such as memory or CPU.`,
}

func init() {
	rootCmd.AddCommand(eatCmd)
}
