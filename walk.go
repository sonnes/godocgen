package godocgen

import (
	"os"
	"path/filepath"
	"strings"
)

// GetSourcePackages returns all package directories in path
// excludes vendor directories
func GetSourcePackages(path string) ([]string, error) {

	absPath, _ := filepath.Abs(path)

	directories := []string{}

	err := filepath.Walk(
		absPath,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() {
				return nil
			}

			if strings.Contains(path, "/vendor/") ||
				strings.Contains(path, ".git") {
				return nil
			}

			directories = append(directories, path)

			return nil
		},
	)

	return directories, err
}
