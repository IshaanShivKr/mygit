package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo(path string) error {
	gitDir := filepath.Join(path, ".mygit")

	dirs := []string{
		gitDir,
		filepath.Join(gitDir, "objects"),
		filepath.Join(gitDir, "refs"),
		filepath.Join(gitDir, "refs", "heads"),
		filepath.Join(gitDir, "refs", "tags"),
	}

	for _, dirPath := range dirs {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}

	headFile := filepath.Join(gitDir, "HEAD")
	headContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(headFile, headContents, 0644); err != nil {
		return fmt.Errorf("failed to write HEAD file: %w", err)
	}

	return nil
}