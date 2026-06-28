package blogpost

import (
	"fmt"
	"os"
	"path/filepath"
)

func ResolveRepoRoot(explicit string) (string, error) {
	if explicit != "" {
		return validateRepoRoot(explicit)
	}

	if env := os.Getenv("BLOG_REPO"); env != "" {
		return validateRepoRoot(env)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if root, ok := findRepoRoot(cwd); ok {
		return root, nil
	}

	return "", fmt.Errorf("nie znaleziono repozytorium Hugo (brak hugo.toml); użyj --repo lub BLOG_REPO")
}

func findRepoRoot(start string) (string, bool) {
	dir := start
	for {
		if fileExists(filepath.Join(dir, "hugo.toml")) || fileExists(filepath.Join(dir, "config.toml")) {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

func validateRepoRoot(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if _, ok := findRepoRoot(abs); !ok {
		return "", fmt.Errorf("%s nie wygląda na katalog Hugo", abs)
	}
	return abs, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
