package main

import (
	"log"
	"os"

	"github.com/vitorbaraujo/batler/configuration"
	"github.com/vitorbaraujo/batler/xcodebuild"

	"github.com/spf13/cobra"
)

var (
	flags = struct {
		projectPath string
	}{}
	rootCmd = &cobra.Command{
		Use:   "batler",
		Short: "Batler is an Xcode test CLI for continuous integration",
		Run: func(cmd *cobra.Command, args []string) {
			runBatler()
		},
	}
)

func init() {
	rootCmd.Flags().StringVar(&flags.projectPath, "project_path", "",
		"An iOS project path that contains a .batler.yml file.")
}

func runBatler() {
	config, err := configuration.FetchConfiguration(flags.projectPath)
	if err != nil {
		log.Fatalf("could not fetch configuration: %v", err)
	}

	if err := xcodebuild.Run(config); err != nil {
		log.Printf("could not run client: %v", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
