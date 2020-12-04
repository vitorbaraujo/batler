package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vitorbaraujo/batler/configuration"
	"github.com/vitorbaraujo/batler/xcodebuild"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for an iOS application",
	Run: func(cmd *cobra.Command, args []string) {
		runTests()
	},
}

func init() {
	testCmd.Flags().StringVarP(&flags.projectPath, "project_path", "p", "",
		"The project path for the iOS application. This path must contain a .batler.yml file.")

	rootCmd.AddCommand(testCmd)
}

func runTests() {
	config, err := configuration.FetchConfiguration(flags.projectPath)
	if err != nil {
		log.Fatalf("could not fetch configuration: %v", err)
	}

	if err := xcodebuild.Run(config); err != nil {
		log.Printf("could not run client: %v", err)
	}
}
