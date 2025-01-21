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
const defaultLoc = "Europe/Paris"
const repoPath = "./repo"
const outputFilePath = repoPath + "/data/date.txt"

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Clone the repository
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "https://github.com/EmielD/Streakinator",
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "Streakinator",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		fmt.Println(os.Getenv("GITHUB_TOKEN"))
		fmt.Println("Error cloning repository:", err)
		return
	}

	// Load time location and file path
	loc, err := time.LoadLocation(defaultLoc)
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
		fmt.Println("Error opening repository:", err)
		return
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}

	// Stage the changes
	_, err = w.Add(".")
	if err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Check if there are any changes to commit
	status, err := w.Status()
	if err != nil {
		fmt.Println("Error checking status:", err)
		return
	}

	if status.IsClean() {
		fmt.Println("No changes detected, skipping commit.")
		return
	}

	// Commit the changes
	commit, err := w.Commit("Update text file with current date/time", &git.CommitOptions{
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

	err = os.RemoveAll(repoPath)
	if err != nil {
		fmt.Println("Something went wrong cleaning up the pulled repository: ", err)
	}

	fmt.Println("Done with cleaning up the pulled repository! Time to sleep... Zzzzz...")
}
