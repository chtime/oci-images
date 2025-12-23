package cmd

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var afterDelay int

var poisonCmd = &cobra.Command{
	Use:   "poison <delay>",
	Short: "Crashes.",
	Long:  `Causes a process crash after the specified delay in seconds.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		afterDelay, err = strconv.Atoi(args[0])
		if err != nil || afterDelay < 0 {
			slog.Error("invalid delay", "error", err)
			return
		}

		slog.Info("spreading poison", "delay", afterDelay)
		time.Sleep(time.Duration(afterDelay) * time.Second)

		panic("Gellerpox")
	},
}

func init() {
	rootCmd.AddCommand(poisonCmd)
}
