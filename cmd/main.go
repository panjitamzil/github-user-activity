package main

import (
	"fmt"
	"os"

	"github-user-activity/internal/api"
	"github-user-activity/internal/formatter"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: Could not load .env file, relying on environment variables:", err)
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable is not set")
		os.Exit(1)
	}

	events, err := api.GetUserActivity(username, token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	activities := formatter.FormatActivities(events)
	fmt.Println("Output:")
	for _, activity := range activities {
		fmt.Println("-", activity)
	}
}
