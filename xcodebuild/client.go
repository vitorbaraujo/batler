package xcodebuild

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vitorbaraujo/batler/configuration"
)

type client struct {
	workspace, scheme, buildDir, xcodeDir string
	cleanEnabled, buildEnabled, testEnabled bool
}

type Option func(*client)

func NewClient(config *configuration.Configuration, opts ...Option) (*client, error) {
	xcodeDir, err := config.XcodeDir()
	if err != nil {
		return nil, fmt.Errorf("fetching xcode developer directory: %w", err)
	}

	c := &client{
		workspace: config.Workspace,
		scheme: config.Scheme,
		buildDir: config.BuildDir,
		xcodeDir: xcodeDir,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func WithClean() Option {
	return func(c *client) {
		c.cleanEnabled = true
	}
}

func WithBuild() Option {
	return func(c *client) {
		c.buildEnabled = true
	}
}

func WithTest() Option {
	return func(c *client) {
		c.testEnabled = true
	}
}

func (c *client) Run() error {
	args := []string{
		"-workspace",
		c.workspace,
		"-scheme",
		c.scheme,
		"-derivedDataPath",
		c.buildDir,
	}

	if c.cleanEnabled {
		args = append(args, "clean")
	}
	if c.buildEnabled {
		args = append(args, "build")
	}
	if c.testEnabled {
		args = append(args, "test")
	}

	cmd := exec.Command("xcodebuild", args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("DEVELOPER_DIR=%s", c.xcodeDir))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running xcodebuild: %w", err)
	}

	return nil
}