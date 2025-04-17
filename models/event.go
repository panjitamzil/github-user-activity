package models

import "time"

// Event represents a GitHub event
type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

// Repo represents a GitHub repository
type Repo struct {
	Name string `json:"name"`
}

// Payload represents the payload of a GitHub event
type Payload struct {
	Action string `json:"action"`
	Size   int    `json:"size"`
}
