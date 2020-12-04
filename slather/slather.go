package slather

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

// Config contains configuration parameters for the Slather CLI.
// See https://github.com/SlatherOrg/slather.
type Config struct {
	HTMLOutput      bool
	IgnoredFiles    []string
	OutputDirectory string
	Scheme          string
	DerivedDataPath string
	XcodeProject    string
	XcodeWorkspace  string
}

// Run runs the slather CLI using the given configuration.
func Run(config *Config) error {
	args := []string{
		"coverage",
		"--output-directory",
		config.OutputDirectory,
		"--workspace",
		config.XcodeWorkspace,
		"--scheme",
		config.Scheme,
		"--build-directory",
		config.DerivedDataPath,
	}

	for _, file := range config.IgnoredFiles {
		args = append(args, []string{
			"--ignore",
			file,
		}...)
	}

	if config.HTMLOutput {
		args = append(args, "--html")
	}

	args = append(args, config.XcodeProject)

	cmd := exec.Command("slather", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if config.HTMLOutput {
		cmd.Stdout = nil
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("coult not run slather: %w", err)
	}

	if config.HTMLOutput {
		outputPath, err := filepath.Abs(filepath.Join(config.OutputDirectory, "index.html"))
		if err != nil {
			return fmt.Errorf("cannot find output directory: %w", err)
		}
		fmt.Printf("Coverage output can be found at: %s\n", outputPath)

		coverage, err := getCoverageFromHTML(config.OutputDirectory)
		if err != nil {
			return fmt.Errorf("could not determine coverage from HTML report: %w", err)
		}

		fmt.Printf("Test Coverage: %s\n", coverage)
	}

	return nil
}

// Slather does not display the coverage when outputting to html,
// so we retrieve the coverage from the generated html report and display it.
func getCoverageFromHTML(outputDir string) (string, error) {
	indexPath := fmt.Sprintf("%s/index.html", outputDir)

	content, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return "", fmt.Errorf("reading HTML output file: %w", err)
	}

	regex := regexp.MustCompile(`<span>Total\ Coverage\ :\ <\/span><span\ class=\"(.*)\"\ id=\"total_coverage\">(.*)<\/span>`)
	matches := regex.FindAllStringSubmatch(string(content), -1)

	if len(matches) == 0 {
		return "", fmt.Errorf("could not parse HTML output")
	}

	// first match (0), second regex match group (2, zero-indexed)
	coverage := matches[0][2]
	return coverage, nil
}
