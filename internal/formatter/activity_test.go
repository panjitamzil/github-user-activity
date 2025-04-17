package formatter

import (
	"reflect"
	"testing"

	"github-user-activity/models"
)

func TestFormatEvent(t *testing.T) {
	tests := []struct {
		name     string
		event    models.Event
		expected string
	}{
		{
			name:     "PushEvent",
			event:    models.Event{Type: "PushEvent", Repo: models.Repo{Name: "testuser/repo"}, Payload: models.Payload{Size: 3}},
			expected: "Pushed 3 commits to testuser/repo",
		},
		{
			name:     "WatchEvent",
			event:    models.Event{Type: "WatchEvent", Repo: models.Repo{Name: "testuser/repo"}},
			expected: "Starred testuser/repo",
		},
		{
			name:     "IssueEvent",
			event:    models.Event{Type: "IssueEvent", Repo: models.Repo{Name: "testuser/repo"}},
			expected: "Opened a new issue in testuser/repo",
		},
		{
			name:     "UnknownEvent",
			event:    models.Event{Type: "UnknownEvent", Repo: models.Repo{Name: "testuser/repo"}},
			expected: "Performed UnknownEvent on testuser/repo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatEvent(tt.event)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFormatActivities(t *testing.T) {
	events := []models.Event{
		{Type: "PushEvent", Repo: models.Repo{Name: "testuser/repo1"}, Payload: models.Payload{Size: 2}},
		{Type: "WatchEvent", Repo: models.Repo{Name: "testuser/repo2"}},
	}
	expected := []string{
		"Pushed 2 commits to testuser/repo1",
		"Starred testuser/repo2",
	}

	result := FormatActivities(events)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFormatActivities_Empty(t *testing.T) {
	events := []models.Event{}
	result := FormatActivities(events)
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}
}
