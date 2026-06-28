package blogpost

import (
	"fmt"
	"os/exec"
	"strings"
)

func CommitAndPush(repoRoot, filePath, title string) error {
	rel, err := relativeToRepo(repoRoot, filePath)
	if err != nil {
		return err
	}

	if err := runGit(repoRoot, "add", rel); err != nil {
		return fmt.Errorf("git add: %w", err)
	}

	msg := fmt.Sprintf("Add post: %s", title)
	if err := runGit(repoRoot, "commit", "-m", msg); err != nil {
		return fmt.Errorf("git commit: %w", err)
	}

	if err := runGit(repoRoot, "push"); err != nil {
		return fmt.Errorf("git push: %w", err)
	}
	return nil
}

func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

func relativeToRepo(repoRoot, filePath string) (string, error) {
	if strings.HasPrefix(filePath, repoRoot) {
		rel := strings.TrimPrefix(filePath, repoRoot)
		return strings.TrimPrefix(rel, "/"), nil
	}
	return filePath, nil
}
