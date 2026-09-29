package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "ci" {
		fmt.Println("Usage: main.go ci")
		os.Exit(1)
	}

	changedFilesPath := os.Getenv("CHANGED_FILES_PATH")
	if changedFilesPath == "" {
		changedFilesPath = "changed_files.json"
	}

	data, err := os.ReadFile(changedFilesPath)
	if err != nil {
		fmt.Printf("Error reading changed files: %v\n", err)
		os.Exit(1)
	}

	var changes map[string]interface{}
	if err := json.Unmarshal(data, &changes); err != nil {
		fmt.Printf("Error parsing changed files: %v\n", err)
		os.Exit(1)
	}

	env := os.Getenv("PUBLICATION_ENV")
	fmt.Printf("Registry automation running in %s environment\n", env)
	fmt.Printf("Changed files: %s\n", string(data))
	fmt.Println("Processing connector updates...")
	fmt.Println("Done.")
}
