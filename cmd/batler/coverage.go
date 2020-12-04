package main

import (
	"fmt"
	"path/filepath"

	"github.com/vitorbaraujo/batler/configuration"
	"github.com/vitorbaraujo/batler/slather"

	"github.com/spf13/cobra"
)

var (
	coverageFlags = struct {
		html        bool
		outputDir   string
		projectPath string
	}{}
	coverageCmd = &cobra.Command{
		Use:   "coverage",
		Short: "Retrieve coverage from previously executed tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCoverage()
		},
	}
)

func init() {
	coverageCmd.Flags().BoolVar(&coverageFlags.html, "html", false, "Output coverage report to HTML")
	coverageCmd.Flags().StringVarP(&coverageFlags.projectPath, "project_path", "p", "",
		"The project path for the iOS application. This path must contain a .batler.yml file.")
	coverageCmd.Flags().StringVarP(&coverageFlags.outputDir, "output_dir", "o", "",
		"Directory to output report files")

	rootCmd.AddCommand(coverageCmd)
}

func runCoverage() error {
	config, err := configuration.FetchConfiguration(coverageFlags.projectPath)
	if err != nil {
		return fmt.Errorf("could not fetch configuration: %w", err)
	}

	if coverageFlags.outputDir == "" {
		coverageFlags.outputDir = filepath.Join(coverageFlags.projectPath, "coverage_out")
	}

	return slather.Run(&slather.Config{
		IgnoredFiles:    config.Coverage.IgnoredFiles,
		HTMLOutput:      coverageFlags.html,
		OutputDirectory: coverageFlags.outputDir,
		Scheme:          config.Scheme,
		DerivedDataPath: config.BuildDir,
		XcodeProject:    config.XcodeProject,
		XcodeWorkspace:  config.Workspace,
	})
}
