package tools

import (
	"os"
)

// StartLog creates a new log file or opens an existing log file for appending
func StartLog() (*os.File, error) {

	// Specify the log file name
	logFileName := "/var/log/j-proxymail.log"

	// Check if the log file exists
	_, err := os.Stat(logFileName)
	var file *os.File

	if err == nil {
		// If file exists, open it for appending
		file, err = os.OpenFile(logFileName, os.O_APPEND|os.O_RDWR, 0644)
		return file, err
	} else {
		// If file does not exist, create a new file
		file, err = os.Create(logFileName)
		return file, err
	}

	return nil, nil
}
