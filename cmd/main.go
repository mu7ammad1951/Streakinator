package main

import (
	"fmt"
	"log"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
)

const timeFormat = "Mon, 02 Jan 2006 15:04:05"
const repoPath = "./repo"
const outputFilePath = repoPath + "/data/date.txt"

func main() {
	// Attempt to fetch environment variables
	repositoryUrl, token, username, email := getEnvironmentVariables()

	// Check if the application is running in GitHub Actions (GITHUB_TOKEN provided)
	if token == "" {
		// If the application is running locally; load the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Re-fetch the environment variables from .env
		repositoryUrl, token, username, email = getEnvironmentVariables()
	}

	// Fail fast if any variables are missing
	if repositoryUrl == "" || token == "" || username == "" || email == "" {
		log.Fatal("Error: Missing required environment variables.\nEnsure S_GITHUB_REPOSITORY_URL, S_GITHUB_TOKEN, S_GITHUB_USERNAME, and S_GITHUB_EMAIL are set.")
	}

	// Clone the repository
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		},
	})
	if err != nil {
		log.Fatalf("Error cloning repository into path %s: %v\n", repoPath, err)
	}

	// Load time location and file path
	loc, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		log.Fatal("Error loading location:", err)
	}

	// Log startup message
	fmt.Printf("Waking up! Writing current time and date to %v\n", outputFilePath)

	// Ensure the data directory exists
	err = os.MkdirAll("./data", 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}

	// Open the file for reading and writing
	file, err := os.OpenFile(outputFilePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Write the current time to the file
	newContent := time.Now().In(loc).Format(timeFormat) + "\n"
	_, err = file.WriteString(newContent)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	// Log file update message
	fmt.Printf("%v has been updated\n", outputFilePath)

	// Open the repository again
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal("Error opening repository:", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatal("Error getting worktree:", err)
	}

	// Stage the changes
	_, err = w.Add(".")
	if err != nil {
		log.Fatal("Error adding changes:", err)
	}

	// Check if there are any unstaged changes in the working tree
	status, err := w.Status()
	if err != nil {
		log.Fatal("Error checking status:", err)
	}

	if status.IsClean() {
		fmt.Println("No changes detected, skipping commit.")
		return
	}

	// Commit the changes
	commitMessage := fmt.Sprintf("[Streakinator] Updated date.txt to the current date: %s", newContent)
	commit, err := w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal("Error committing changes:", err)
		return
	}

	// Print the commit hash
	fmt.Println("Commit hash:", commit)

	// Push the changes to the remote repository
	for i := 0; i < 3; i++ {
		err = repo.Push(&git.PushOptions{
			Auth: &http.BasicAuth{
				Username: username,
				Password: token,
			},
		})
		if err == nil {
			fmt.Println("Changes pushed successfully!")
			break
		}
		fmt.Printf("Error pushing changes (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	if err != nil {
		log.Fatal("Failed to push changes after 3 attempts:", err)
	}

	defer func() {
		err := os.RemoveAll(repoPath)
		if err != nil {
			log.Fatalf("Error cleaning up repoPath %s: %v\n", repoPath, err)
		}
	}()

	fmt.Println("Done with cleaning up the pulled repository! Time to sleep... Zzzzz...")
}

func getEnvironmentVariables() (string, string, string, string) {
	return os.Getenv("S_GITHUB_REPOSITORY_URL"),
		os.Getenv("S_GITHUB_TOKEN"),
		os.Getenv("S_GITHUB_USERNAME"),
		os.Getenv("S_GITHUB_EMAIL")
}
