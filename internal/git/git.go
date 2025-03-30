package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"rdr/git-archiver/internal/config"
)

func Is_git_installed() error {
	out, err := exec.Command("git", "-v").Output()

	if err != nil {
		return fmt.Errorf("git not found: %s", err)
	}

	output := string(out[:])
	log.Println(output, "is installed!")
	return nil
}

func Sync(repo config.RepoConfig) error {
	var repoExists bool = directoryExists(repo.LocalPath) && directoryExists(filepath.Join(repo.LocalPath, ".git"))

	var cmd *exec.Cmd

	if repoExists {
		// repository exists, perform a pull
		log.Println("pulling repo:", repo.URL)
		cmd = exec.Command("git", "-C", repo.LocalPath, "pull", "--all")
	} else {
		// repository doesn't exist, perform a clone
		log.Print("clonging repo:", repo.URL)

		// create the local path directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(repo.LocalPath), 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory for repo: %w", err)
		}
		cmd = exec.Command("git", "clone", repo.URL, repo.LocalPath)
	}

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to pull/clone repo: %s", cmd.Err)
	}

	return nil
}

// directoryExists checks if a directory exists
func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}