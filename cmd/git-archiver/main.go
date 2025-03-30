package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"rdr/git-archiver/internal/config"
	"rdr/git-archiver/internal/git"
)

func main() {
	if git.Is_git_installed() != nil {
		log.Fatal("Git command not found! Make sure git is installed properly.")
		os.Exit(1)
	}

	// Load configuration from file
	config, err := config.ReadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Print loaded configuration details
	fmt.Printf("Loaded configuration with %d repositories\n", len(config.Repositories))
	for i, repo := range config.Repositories {
		fmt.Printf("Repository %d:\n", i+1)
		fmt.Printf("  URL: %s\n", repo.URL)
		fmt.Printf("  Update Interval: %v\n", repo.Interval)
		fmt.Printf("  Local Path: %s\n", repo.LocalPath)
		fmt.Println()
	}

	// Start syncing repositories in parallel
	for _, repo := range config.Repositories {
		go func() {
			ticker := time.NewTicker(repo.Interval)
			// Ensure the ticker is stopped when we're done (though here it runs forever).
			defer ticker.Stop()

			// loop forever, waiting for the ticker to send a signal
			for {
				<-ticker.C

				git.Sync(repo)
				fmt.Println("done!:", repo.URL)
			}
		}()
		
		err := git.Sync(repo)
		if err != nil {
			fmt.Println("error cloning:", err)
		}
	}

	// prevent main func from exiting by blocking forever
	select {}
}
