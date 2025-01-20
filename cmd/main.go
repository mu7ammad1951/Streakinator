package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	outputFilePath := "./data/data.txt"

	// GMT +1 timezone offset 3600 seconds - 1 hour
	gmtTimeLoc := time.FixedZone("GMT", 3600)

	// Get current time and convert it to GMT+1
	currentTime := time.Now().In(gmtTimeLoc).Format("Mon, 02 Jan 2006 15:04:05")

	fmt.Printf("[%v - GMT +1] Waking up! Writing current time and date to %v\n", currentTime, outputFilePath)

	// Making sure the directory we'll write the currentTime value to exists, if not: create it
	// Setting the filemode to 0755 - owner can read/write/execute, others can read/execute
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		fmt.Printf("[%v - GMT +1] failed to locate directory, creating it now\n", currentTime)
		err := os.Mkdir("./data", 0755)
		if err != nil {
			fmt.Printf("[%v - GMT +1] failed to create directory: %v\n", currentTime, err)
			return
		}
	}

	// Writing the current time to the data.txt file, it uses the 0644 filemode to be read/writeable
	// for the owner and readable by others
	err := os.WriteFile(outputFilePath, []byte(fmt.Sprintf("%v", currentTime)), 0644)
	if err != nil {
		fmt.Printf("[%v - GMT +1] error occurred: %v\n", currentTime, err)
		return
	}

	fmt.Printf("[%v - GMT +1] %v has been updated\n", currentTime, outputFilePath)

	// TODO: Everything is done, create a new commit and push to the associated git repository specified in the .env file
}
