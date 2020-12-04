package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.3.0"
	rootCmd = &cobra.Command{
		Use:     "batler",
		Short:   "Batler is an Xcode test CLI for continuous integration",
		Version: version,
	}
)

func init() {
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
