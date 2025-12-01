package stacks

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var stackFilePattern = regexp.MustCompile(`^Pulumi\.(.+)\.yaml$`)

// Detect finds all Pulumi stack names in the given directory by looking for
// files matching the pattern "Pulumi.<stack-name>.yaml"
func Detect(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var stacks []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		matches := stackFilePattern.FindStringSubmatch(entry.Name())
		if len(matches) == 2 {
			stacks = append(stacks, matches[1])
		}
	}

	return stacks, nil
}

// DetectInCurrentDir finds all Pulumi stack names in the current directory
func DetectInCurrentDir() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return Detect(dir)
}

// FilterStacks returns stacks that start with the given prefix
func FilterStacks(stacks []string, prefix string) []string {
	if prefix == "" {
		return stacks
	}

	var filtered []string
	for _, s := range stacks {
		if strings.HasPrefix(s, prefix) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

// HasPulumiProject checks if the directory contains a Pulumi.yaml file
func HasPulumiProject(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "Pulumi.yaml"))
	return err == nil
}
