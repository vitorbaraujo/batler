package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/vitorbaraujo/batler/configuration"
	"github.com/vitorbaraujo/batler/slather"

	"github.com/spf13/cobra"
)

type coverageCmdFlags struct {
	html        bool
	outputDir   string
	projectPath string
}

var (
	coverageFlags = coverageCmdFlags{}
	coverageCmd   = &cobra.Command{
		Use:   "coverage",
		Short: "Retrieve coverage from previously executed tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCoverage()
		},
	}
)

func init() {
	coverageCmd.Flags().BoolVar(&coverageFlags.html, "html", false, "Exports the coverage report to HTML")
	coverageCmd.Flags().StringVarP(&coverageFlags.projectPath, "project_path", "p", "",
		"The project path for the iOS application. This path must contain a \".batler.yml\" file. "+
			"Defaults to current directory")
	coverageCmd.Flags().StringVarP(&coverageFlags.outputDir, "output_dir", "o", "",
		"Directory to output report files. "+
			"This option can only be used when exporting the coverage report."+
			"Defaults to \"<project_path>/coverage-out/\"")

	rootCmd.AddCommand(coverageCmd)
}

func runCoverage() error {
	config, err := configuration.FetchConfiguration(coverageFlags.projectPath)
	if err != nil {
		return fmt.Errorf("could not fetch configuration: %w", err)
	}

	switch {
	case !isExportingCoverage(coverageFlags) && coverageFlags.outputDir != "":
		return errors.New("cannot set output_dir without exporting report")
	case isExportingCoverage(coverageFlags) && coverageFlags.outputDir == "":
		coverageFlags.outputDir = filepath.Join(coverageFlags.projectPath, "coverage_out")
	}

	if err = slather.Run(&slather.Config{
		IgnoredFiles:    config.Coverage.IgnoredFiles,
		HTMLOutput:      coverageFlags.html,
		OutputDirectory: coverageFlags.outputDir,
		Scheme:          config.Scheme,
		DerivedDataPath: config.BuildDir,
		XcodeProject:    config.XcodeProject,
		XcodeWorkspace:  config.Workspace,
	}); err != nil {
		return fmt.Errorf("could not run slather: %w", err)
	}

	return nil
}

func isExportingCoverage(flags coverageCmdFlags) bool {
	return flags.html
}
