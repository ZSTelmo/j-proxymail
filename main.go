package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"j-proxymail/models"
	"j-proxymail/tools"

	// ...existing imports...

	"github.com/TwiN/go-color"
	"github.com/mcnijman/go-emailaddress"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Run the application
func run(logger zerolog.Logger) error {
	emailJSON, err := parseJSONArgs()
	if err != nil {
		return err
	}

	logger.Info().Msg("Destination: " + emailJSON.Destination)
	logger.Info().Msg("Subject: " + emailJSON.Subject)
	logger.Info().Msg("Body: " + emailJSON.Body)

	text := []byte(emailJSON.Destination)
	validateHost := false

	emails := emailaddress.FindWithIcannSuffix(text, validateHost)

	for _, cleanEmail := range emails {
		fmt.Printf(color.Colorize(color.Green, "Email: %v \n"), cleanEmail)
		err := tools.SendEmail(fmt.Sprint(cleanEmail), emailJSON.Subject, emailJSON.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsJSON checks if a string is valid JSON
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsValidEmailJSON checks if a string is valid JSON with required email fields
func IsValidEmailJSON(str string) bool {
	var emailData models.EmailJSON
	err := json.Unmarshal([]byte(str), &emailData)
	if err != nil {
		return false
	}

	// Check if all required fields are present and not empty
	return emailData.Destination != "" &&
		emailData.Subject != "" &&
		emailData.Body != ""
}

func parseJSONArgs() (*models.EmailJSON, error) {
	jsonStr := flag.String("json", "", "JSON string containing email details")
	flag.Parse()

	if *jsonStr == "" {
		return nil, fmt.Errorf("JSON argument is required")
	}

	if !IsJSON(*jsonStr) {
		return nil, fmt.Errorf("invalid JSON format")
	}

	if !IsValidEmailJSON(*jsonStr) {
		return nil, fmt.Errorf("JSON missing required email fields")
	}

	var emailData models.EmailJSON
	err := json.Unmarshal([]byte(*jsonStr), &emailData)
	if err != nil {
		return nil, err
	}

	return &emailData, nil
}

// Main function
func main() {

	file, err := tools.StartLog()

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create log file")
	}
	defer file.Close()

	// Set up the logger to write to the file
	logger := zerolog.New(file).With().Timestamp().Logger()

	// Log some messages
	//logger.Info().Msg("This is an info message")
	//logger.Error().Msg("This is an error message")

	// You can also log to both the console and a file if you want
	//consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	//multi := zerolog.MultiLevelWriter(consoleWriter, file)
	//logger = zerolog.New(multi).With().Timestamp().Logger()

	// This log will go to both console and the log file
	//logger.Info().Msg("This message goes to both the console and the file")

	error := run(logger)
	if error != nil {
		log.Panic().Err(err).Msg("Unable to Run Application")
	}
}
