package cmd

import (
	"log/slog"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

var numberProcesses int

var cpuCmd = &cobra.Command{
	Use:   "cpu <processes>",
	Short: "Consumes CPU resources across the specified number of parallel processes.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		numberProcesses, err = strconv.Atoi(args[0])
		if err != nil || numberProcesses < 1 {
			slog.Error("invalid number of processes", "error", err)
			return
		}

		var waitGroup sync.WaitGroup
		slog.Info("starting cpu consumption", "processes", numberProcesses)
		for i := range numberProcesses {
			waitGroup.Add(1)
			go func() {
				slog.Debug("starting goroutine", "n", i)
				defer waitGroup.Done()
				primes()
			}()
		}

		// wait for all goroutines to finish
		waitGroup.Wait()
		slog.Info("finished prime calculations")
	},
}

func init() {
	eatCmd.AddCommand(cpuCmd)
}

func primes() {
	const maxInt = 1<<63 - 1
	var highestPrime = 0
	for i := range maxInt {
		if isPrime(i) {
			highestPrime = max(highestPrime, i)
		}
	}
	slog.Info("last prime number found", "prime", highestPrime)
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
