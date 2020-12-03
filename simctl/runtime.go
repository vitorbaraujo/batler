package simctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Runtimes is the representation for the runtimes list from `simctl`.
type runtimes struct {
	Runtimes []*Runtime `json:"runtimes"`
}

// Runtime represents a simulator runtime.
type Runtime struct {
	Name       string
	Identifier string
	Available  bool `json:"isAvailable"`
}

// ListRuntimes lists all existing runtimes from the current Xcode version.
func ListRuntimes(xcodePath string) ([]*Runtime, error) {
	cmd := exec.Command("xcrun", "simctl", "list", "runtimes", "--json")
	cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", xcodePath))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running simctl list: %w", err)
	}

	return parseRuntimesOutput(output)
}

func parseRuntimesOutput(output []byte) ([]*Runtime, error) {
	var dt *runtimes
	if err := json.Unmarshal(output, &dt); err != nil {
		return nil, fmt.Errorf("parsing runtimes output: %w", err)
	}

	return dt.Runtimes, nil
}
