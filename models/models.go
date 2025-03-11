package models

// EmailJSON represents the structure of an email JSON payload
type EmailJSON struct {
	Destination string `json:"destination"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}
