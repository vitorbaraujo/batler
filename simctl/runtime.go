package simctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Runtimes struct {
	Runtimes []*Runtime `json:"runtimes"`
}

type Runtime struct {
	Name       string
	Identifier string
	Available  bool `json:"isAvailable"`
}

func ListRuntimes() ([]*Runtime, error) {
	cmd := exec.Command("xcrun", "simctl", "list", "runtimes", "--json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running simctl list: %w", err)
	}

	return parseRuntimesOutput(output)
}

func parseRuntimesOutput(output []byte) ([]*Runtime, error) {
	var dt *Runtimes

	if err := json.Unmarshal(output, &dt); err != nil {
		return nil, fmt.Errorf("parsing runtimes output: %w", err)
	}

	return dt.Runtimes, nil
}
