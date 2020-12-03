package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const configName = ".batler.yml"

// Configuration contains configuration parameters for the CLI.
type Configuration struct {
	BuildDir          string `yaml:"build_dir"`
	Scheme            string
	xcodeVersion      string `yaml:"xcode_version"`
	xcodeDeveloperDir string `yaml:"xcode_developer_dir"`
	XcodePath         string
	Workspace         string
}

// FetchConfiguration retrieves configurations from the config file inside projectPath.
func FetchConfiguration(projectPath string) (*Configuration, error) {
	configPath := filepath.Join(projectPath, configName)
	if !fileExists(configPath) {
		return nil, fmt.Errorf("configuration file does not exist at %v", configPath)
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	c := Configuration{}
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, fmt.Errorf("could not parse yaml file: %w", err)
	}

	c.Workspace = filepath.Join(projectPath, c.Workspace)
	c.BuildDir = filepath.Join(projectPath, c.BuildDir)

	if err := c.isValid(); err != nil {
		return nil, fmt.Errorf("configuration is not valid: %w", err)
	}

	c.XcodePath, err = getXcodePath(c.xcodeVersion, c.xcodeDeveloperDir)
	if err != nil {
		return nil, fmt.Errorf("could not get xcode path: %w", err)
	}

	return &c, nil
}

func (c *Configuration) isValid() error {
	if c.BuildDir == "" {
		return errors.New("missing build_dir")
	}

	if c.Scheme == "" {
		return errors.New("missing scheme")
	}

	if c.Workspace == "" {
		return errors.New("missing workspace")
	}

	if c.xcodeVersion != "" && c.xcodeDeveloperDir != "" {
		return errors.New("cannot set xcode_version and xcode_developer_dir at the same time")
	}

	return nil
}

func getXcodePath(xcodeVersion, xcodeDeveloperDir string) (string, error) {
	if xcodeVersion == "" && xcodeDeveloperDir == "" {
		cmd := exec.Command("xcode-select", "-p")

		defaultXcodePath, err := cmd.Output()
		if err != nil {
			return "", errors.New("cannot fetch default xcode path using `xcode-select -p`")
		}

		return strings.TrimRight(string(defaultXcodePath), "\r\n"), nil
	}

	if xcodeVersion != "" {
		return fmt.Sprintf("/Applications/Xcode-%s.app/Contents/Developer", xcodeVersion), nil
	}

	return xcodeDeveloperDir, nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
