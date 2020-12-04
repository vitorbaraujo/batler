package slather

import (
	"fmt"
	"os"
	"os/exec"
)

// Config contains configuration parameters for the Slather CLI.
// See https://github.com/SlatherOrg/slather.
type Config struct {
	HTMLOutput      bool
	IgnoredFiles    []string
	OutputDirectory string
	Scheme          string
	DerivedDataPath string
	XcodeProject    string
	XcodeWorkspace  string
}

// Run runs the slather CLI using the given configuration.
func Run(config *Config) error {
	args := []string{
		"coverage",
		"--output-directory",
		fmt.Sprintf("\"%s\"", config.OutputDirectory),
		"--workspace",
		fmt.Sprintf("\"%s\"", config.XcodeWorkspace),
		"--scheme",
		config.Scheme,
		"--build-directory",
		config.DerivedDataPath,
	}

	for _, file := range config.IgnoredFiles {
		args = append(args, []string{
			"--ignore",
			file,
		}...)
	}

	if config.HTMLOutput {
		args = append(args, "--html")
	}

	args = append(args, config.XcodeProject)

	cmd := exec.Command("slather", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("coult not run slather: %w", err)
	}

	return nil
}
