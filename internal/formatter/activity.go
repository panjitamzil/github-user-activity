package formatter

import (
	"fmt"

	"github-user-activity/models"
)

// FormatEvent formats a single event into a readable string
func FormatEvent(event models.Event) string {
	switch event.Type {
	case "PushEvent":
		return fmt.Sprintf("Pushed %d commits to %s", event.Payload.Size, event.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s", event.Repo.Name)
	case "IssueEvent":
		return fmt.Sprintf("Opened a new issue in %s", event.Repo.Name)
	default:
		return fmt.Sprintf("Performed %s on %s", event.Type, event.Repo.Name)
	}
}

// FormatActivities formats a list of events
func FormatActivities(events []models.Event) []string {
	var activities []string
	for _, event := range events {
		activities = append(activities, FormatEvent(event))
	}
	return activities
}
