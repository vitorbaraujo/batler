package simctl

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type DeviceTypes struct {
	Types []*DeviceType `json:"devicetypes"`
}

type DeviceType struct {
	Name          string
	Identifier    string
	ProductFamily string
}

func ListDeviceTypes() ([]*DeviceType, error) {
	cmd := exec.Command("xcrun", "simctl", "list", "devicetypes", "--json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("running simctl list: %w", err)
	}

	return parseDeviceTypesOutput(output)
}

func parseDeviceTypesOutput(output []byte) ([]*DeviceType, error) {
	var dt *DeviceTypes

	if err := json.Unmarshal(output, &dt); err != nil {
		return nil, fmt.Errorf("parsing devicetypes output: %w", err)
	}

	return dt.Types, nil
}
