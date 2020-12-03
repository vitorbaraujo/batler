package xcodebuild

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vitorbaraujo/batler/configuration"
)

// Client contains methods for cleaning, building and testing an Xcode application.
type Client struct {
	workspace, scheme, buildDir, xcodeDir   string
	cleanEnabled, buildEnabled, testEnabled bool
}

// Option is a configuration option for the Xcodebuild client.
type Option func(*Client)

// NewClient creates a new Xcodebuild client.
func NewClient(config *configuration.Configuration, opts ...Option) (*Client, error) {
	c := &Client{
		workspace: config.Workspace,
		scheme:    config.Scheme,
		buildDir:  config.BuildDir,
		xcodeDir:  config.XcodePath,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// WithClean enables cleaning the iOS application.
func WithClean() Option {
	return func(c *Client) {
		c.cleanEnabled = true
	}
}

// WithBuild enables building the iOS application.
func WithBuild() Option {
	return func(c *Client) {
		c.buildEnabled = true
	}
}

// WithTest enables running tests the iOS application.
func WithTest() Option {
	return func(c *Client) {
		c.testEnabled = true
	}
}

// Run runs `xcodebuild` with the actions enabled using options.
func (c *Client) Run() error {
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
