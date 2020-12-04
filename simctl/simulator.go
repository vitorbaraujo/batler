package simctl

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const simulatorCacheFilename = "/tmp/batler-simulator"

// CreateSimulator creates an arbitrary iOS simulator from existing runtimes and devicetypes.
// This method returns the identifier for the newly created simulator or an error, otherwise.
func CreateSimulator(xcodePath string) (string, error) {
	cachedSimulatorID, err := getSimulatorFromCache(xcodePath)
	if err == nil {
		return cachedSimulatorID, nil
	}

	runtimes, err := ListRuntimes(xcodePath)
	if err != nil {
		return "", fmt.Errorf("could not get runtimes: %w", err)
	}

	deviceTypes, err := ListDeviceTypes(xcodePath)
	if err != nil {
		return "", fmt.Errorf("could not get device types")
	}

	for _, runtime := range runtimes {
		for _, deviceType := range deviceTypes {
			deviceName := createDeviceName(runtime.Identifier, deviceType.Identifier)
			cmd := exec.Command("xcrun", "simctl", "create",
				deviceName, deviceType.Identifier, runtime.Identifier)
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", xcodePath))

			output, err := cmd.Output()
			if err != nil {
				continue
			}

			simulatorID := strings.TrimRight(string(output), "\r\n")

			if err := saveSimulatorOnCache(simulatorID); err != nil {
				// TODO (vitor.araujo): replace fmt by logging
				fmt.Printf("failed to cache simulator id: %v\n", err)
			}

			return simulatorID, nil
		}
	}

	return "", fmt.Errorf("could not create simulator from device types and runtimes")
}

func saveSimulatorOnCache(simulatorID string) error {
	return ioutil.WriteFile(simulatorCacheFilename, []byte(simulatorID), 0o600)
}

func getSimulatorFromCache(xcodePath string) (string, error) {
	output, err := ioutil.ReadFile(simulatorCacheFilename)
	if err != nil {
		return "", fmt.Errorf("could not read cache file: %w", err)
	}

	simulatorID := strings.TrimRight(string(output), "\r\n")
	if !simulatorExists(simulatorID, xcodePath) {
		return "", fmt.Errorf("cached simulator with ID=%s does not exist", simulatorID)
	}

	return simulatorID, nil
}

func simulatorExists(deviceID, xcodePath string) bool {
	cmd := exec.Command("xcrun", "simctl", "list", "devices", deviceID)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", xcodePath))

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.Contains(string(output), deviceID)
}

func createDeviceName(runtimeID, deviceTypeID string) string {
	runtimeIDParts := strings.Split(runtimeID, ".")
	osVersion := runtimeIDParts[len(runtimeIDParts)-1]

	deviceIDParts := strings.Split(deviceTypeID, ".")
	deviceIDName := deviceIDParts[len(deviceIDParts)-1]

	return fmt.Sprintf("%s-%s", deviceIDName, osVersion)
}
