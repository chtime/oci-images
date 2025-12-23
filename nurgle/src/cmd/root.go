package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nurgle",
	Short: "An entity to bring despair over its surroundings.",
	Long: `
Runs different scenarios that will test the limits of the environment.
  - eat memory <amount>: consumes <amount> memory; amount can be human-readable formats like 50Mi.
  - eat cpu <n>: consumes any cpu cycles available for <n> cpus.
  - poison <delay>: will cause this process to crash after <delay> seconds.
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// set up logging
		verbose, _ := cmd.Flags().GetBool("verbose")

		level := slog.LevelInfo
		if verbose {
			level = slog.LevelDebug
		}

		opts := &slog.HandlerOptions{
			Level: level,
		}
		logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
		slog.SetDefault(logger)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")
}
