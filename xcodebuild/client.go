package xcodebuild

import (
	"fmt"
	"os"
	"os/exec"
)

type client struct {
	workspace, scheme, buildDir string
	cleanEnabled, buildEnabled, testEnabled bool
}

type Option func(*client)

func NewClient(workspace, scheme, buildDir string, opts ...Option) *client {
	c := &client{
		workspace: workspace,
		scheme: scheme,
		buildDir: buildDir,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running xcodebuild: %v", err)
	}

	return nil
}