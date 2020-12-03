package simctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type deviceTypes struct {
	Types []*DeviceType `json:"devicetypes"`
}

// DeviceType represents a simulator device type.
type DeviceType struct {
	Name          string
	Identifier    string
	ProductFamily string
}

// ListDeviceTypes lists all existing device types for the current Xcode version.
func ListDeviceTypes(xcodePath string) ([]*DeviceType, error) {
	cmd := exec.Command("xcrun", "simctl", "list", "devicetypes", "--json")
	cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", xcodePath))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running simctl list: %w", err)
	}

	return parseDeviceTypesOutput(output)
}

func parseDeviceTypesOutput(output []byte) ([]*DeviceType, error) {
	var dt *deviceTypes
	if err := json.Unmarshal(output, &dt); err != nil {
		return nil, fmt.Errorf("parsing devicetypes output: %w", err)
	}

	return dt.Types, nil
}
