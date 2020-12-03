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

// Configuration contains configuration parameters for the Xcode application.
type Configuration struct {
	BuildDir  string
	Scheme    string
	XcodePath string
	Workspace string
}

type configurationFile struct {
	BuildDir          string `yaml:"build_dir"`
	Scheme            string `yaml:"scheme"`
	XcodeVersion      string `yaml:"xcode_version"`
	XcodeDeveloperDir string `yaml:"xcode_developer_dir"`
	Workspace         string `yaml:"workspace"`
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

	return createConfiguration(yamlFile, projectPath)
}

func createConfiguration(fileContent []byte, projectPath string) (*Configuration, error) {
	configFile := &configurationFile{}
	if err := yaml.Unmarshal(fileContent, configFile); err != nil {
		return nil, fmt.Errorf("could not parse yaml file: %w", err)
	}

	config := &Configuration{
		BuildDir:  filepath.Join(projectPath, configFile.BuildDir),
		Scheme:    configFile.Scheme,
		Workspace: filepath.Join(projectPath, configFile.Workspace),
	}

	if err := configFile.isValid(); err != nil {
		return nil, fmt.Errorf("configuration is not valid: %w", err)
	}

	var err error
	config.XcodePath, err = getXcodePath(configFile.XcodeVersion, configFile.XcodeDeveloperDir)
	if err != nil {
		return nil, fmt.Errorf("could not get xcode path: %w", err)
	}

	return config, nil
}

func (c configurationFile) isValid() error {
	if c.BuildDir == "" {
		return errors.New("missing build_dir")
	}

	if c.Scheme == "" {
		return errors.New("missing scheme")
	}

	if c.Workspace == "" {
		return errors.New("missing workspace")
	}

	if c.XcodeVersion != "" && c.XcodeDeveloperDir != "" {
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
