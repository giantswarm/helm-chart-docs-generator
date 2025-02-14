package git

import (
	"fmt"
	"os/exec"
)

// CloneRepositoryShallow will clone repository in a given directory.
func CloneRepositoryShallow(user string, repo string, tag string, destDir string) error {
	{
		cmd := exec.Command("git", "clone", "-b", tag, "--depth", "1", fmt.Sprintf("https://github.com/%s/%s.git", user, repo), destDir) // nolint: gosec
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to clone the repo %q with %w", repo, err)
		}
	}

	return nil
}
