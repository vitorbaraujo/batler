package xcodebuild

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vitorbaraujo/batler/simctl"

	"github.com/vitorbaraujo/batler/configuration"
)

// Run runs `xcodebuild` with the given configuration.
func Run(config *configuration.Configuration) error {
	if config.Destination == "" {
		destination, err := simctl.CreateSimulator(config.XcodePath)
		if err != nil {
			return fmt.Errorf("could not create destination: %w", err)
		}
		config.Destination = formatDestinationOutput(destination)
	}

	args := []string{
		"-workspace",
		config.Workspace,
		"-scheme",
		config.Scheme,
		"-derivedDataPath",
		config.BuildDir,
		"-destination",
		config.Destination,
	}

	if config.CleanBuild {
		args = append(args, "clean")
	}

	args = append(args, "build")
	args = append(args, "test")

	cmd := exec.Command("xcodebuild", args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", config.XcodePath))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running xcodebuild: %w", err)
	}

	return nil
}

func formatDestinationOutput(simulatorID string) string {
	return fmt.Sprintf("platform=iOS Simulator,id=%s", simulatorID)
}
