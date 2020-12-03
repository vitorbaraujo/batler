package simctl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Creates an arbitrary iOS simulator from existing runtimes and devicetypes.
// This method returns the identifier for the newly created simulator or an error, otherwise.
func CreateSimulator(xcodePath string) (string, error) {
	var err error
	runtimes, err := ListRuntimes(xcodePath)
	if err != nil {
		return "", fmt.Errorf("could not get runtimes: %w", err)
	}

	deviceTypes, err := ListDeviceTypes(xcodePath)
	if err != nil {
		return "", fmt.Errorf("could not get device types")
	}

	for _, runtime := range runtimes {
		for _, deviceType := range deviceTypes.Types {

			cmd := exec.Command("xcrun", "simctl", "create",
				createDeviceName(runtime.Identifier, deviceType.Identifier), deviceType.Identifier, runtime.Identifier)
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", xcodePath))

			output, err := cmd.Output()
			if err != nil {
				continue
			}

			simulatorId := strings.TrimRight(string(output), "\r\n")
			return simulatorId, nil
		}
	}

	return "", fmt.Errorf("could not create simulator from device types and runtimes")
}

func createDeviceName(runtimeId, deviceTypeId string) string {
	runtimeIdParts := strings.Split(runtimeId, ".")
	osVersion := runtimeIdParts[len(runtimeIdParts)-1]

	deviceIdParts := strings.Split(deviceTypeId, ".")
	deviceIdName := deviceIdParts[len(deviceIdParts)-1]

	return fmt.Sprintf("%s-%s", deviceIdName, osVersion)
}
