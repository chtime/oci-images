package cmd

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

type ByteSize uint64

var memoryAmount ByteSize

// memoryCmd represents the memory command
var memoryCmd = &cobra.Command{
	Use:   "memory <amount>",
	Short: "Consumes memory indefinitely. The memory amount should be a human readable format, e.g. 50Mi, 10G",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := memoryAmount.Set(args[0]); err != nil {
			slog.Error("invalid memory amount", "error", err)
			return
		}

		slog.Info("allocating memory", "size", memoryAmount.String())

		buf := make([]byte, memoryAmount)
		// Touch every page to force physical memory allocation
		// OS uses lazy allocation - pages aren't committed until written
		pageSize := 4096
		for i := 0; i < len(buf); i += pageSize {
			buf[i] = 1
		}
		slog.Debug("memory allocated")

		// can't block this single goroutine since a deadlock is automatically panicked upon
		const maxDuration time.Duration = 1<<63 - 1
		slog.Warn("sleeping virtually indefinitely", "duration", maxDuration)
		time.Sleep(maxDuration)

		slog.Info("releasing memory block", "size", len(buf))
	},
}

func init() {
	eatCmd.AddCommand(memoryCmd)
}

func (b *ByteSize) String() string {
	return humanize.IBytes(uint64(*b))
}

func (b *ByteSize) Set(s string) error {
	parsed, err := humanize.ParseBytes(s)
	if err != nil {
		return fmt.Errorf("invalid byte size: %s", s)
	}
	*b = ByteSize(parsed)
	return nil
}

func (b *ByteSize) Type() string {
	return "bytes"
}
