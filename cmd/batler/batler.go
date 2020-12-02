package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"batler/xcodebuild"
)

func main() {
	projectPath := flag.String("p", "", "Project path to run batler")
	flag.Parse()

	config, err := fetchConfiguration(*projectPath)
	if err != nil {
		log.Fatalf("could not fetch configuration: %v", err)
	}

	client := xcodebuild.NewClient(config.Workspace, config.Scheme, config.BuildDir,
		xcodebuild.WithClean(), xcodebuild.WithBuild(), xcodebuild.WithTest())

	if err := client.Run(); err != nil {
		log.Printf("could not run client: %v", err)
	}
}

const configName = ".batler.yml"

type Configuration struct {
	Workspace string
	Scheme string
	BuildDir string `yaml:"build_dir"`
}

func fetchConfiguration(projectPath string) (*Configuration, error) {
	configPath := filepath.Join(projectPath, configName)
	if !fileExists(configPath) {
		return nil, fmt.Errorf("configuration file does not exist at %v", configPath)
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	c := Configuration{}
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, fmt.Errorf("could not parse yaml file: %w", err)
	}

	c.Workspace = filepath.Join(projectPath, c.Workspace)
	// TODO (vitor.araujo): also join projectPath with c.BuildDir

	return &c, nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
