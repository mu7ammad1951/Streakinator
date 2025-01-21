package main

import (
	"fmt"
	"log"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const timeFormat = "Mon, 02 Jan 2006 15:04:05"
const defaultLoc = "Europe/Paris"
const filePath = "./data/date.txt"

func main() {
	repoPath := "./repo"

	// Always clone the repository
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "https://github.com/EmielD/Streakinator",
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "Streakinator",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	// Log startup message
	fmt.Printf("Waking up! Writing current time and date to %v\n", filePath)

	// Ensure the data directory exists, open the file, and write the current time
	err = os.MkdirAll("./data", 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Load time location and file path, or use defaults
	loc, err := time.LoadLocation(defaultLoc)
	if err != nil {
		log.Fatal("Error loading location:", err)
	}

	currentFormattedTime := time.Now().In(loc).Format(timeFormat)
	_, err = file.WriteString(currentFormattedTime + "\n")
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	// Log file update message
	fmt.Printf("%v has been updated\n", filePath)

	// Stage the changes
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}
	_, err = w.Add(".")
	if err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Commit the changes
	commit, err := w.Commit(fmt.Sprintf("Updated text file with date: %v", currentFormattedTime), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Streakinator",
			Email: "bot@bot.bot",
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Print the commit hash
	fmt.Println("Commit hash:", commit)

	// Push the changes to the remote repository
	err = repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "Streakinator",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		fmt.Println("Error pushing changes:", err)
		return
	}

	fmt.Println("Changes pushed successfully!")
}
