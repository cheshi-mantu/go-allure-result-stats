package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Step 1: Read the ALLURE_RESULTS environment variable
	allureResults := os.Getenv("ALLURE_RESULTS")
	if allureResults == "" {
		fmt.Println("ALLURE_RESULTS environment variable is not set.")
		return
	}

	// Step 2: Scan the directory using os.ReadDir
	entries, err := os.ReadDir(allureResults)
	if err != nil {
		fmt.Printf("Failed to read directory: %s\n", err)
		return
	}

	// Step 3: Filter out *-result.json files
	pattern := regexp.MustCompile(`.*-result\.json$`)
	statusCounts := make(map[string]int)

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}

		if !pattern.MatchString(entry.Name()) {
			continue // Skip non-matching files
		}

		// Step 4: Read each file and extract the `status` attribute
		filePath := filepath.Join(allureResults, entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Failed to read file %s: %s\n", entry.Name(), err)
			continue
		}

		var result map[string]interface{}
		if err := json.Unmarshal(content, &result); err != nil {
			fmt.Printf("Failed to parse JSON from file %s: %s\n", entry.Name(), err)
			continue
		}

		if status, ok := result["status"].(string); ok {
			// Step 5: Count occurrences of each `status`
			statusCounts[status]++
		}
	}

	// Step 6: Print the status counts
	// fmt.Println("Status counts:")
	// for status, count := range statusCounts {
	// 	fmt.Printf("\"%s\" - %d\n", status, count)
	// }
	// After counting the status occurrences and storing them in the statusCounts map

	for status, count := range statusCounts {
		envVarName := fmt.Sprintf("ALLURE_%s", strings.ToUpper(status))
		envVarValue := fmt.Sprintf("%d", count)

		// Set the environment variable
		err := os.Setenv(envVarName, envVarValue)
		if err != nil {
			fmt.Printf("Failed to set environment variable %s: %s\n", envVarName, err)
		} else {
			fmt.Printf("export %s=%s\n", envVarName, envVarValue)
		}
	}

}
