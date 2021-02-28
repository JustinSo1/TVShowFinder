package internal

import "log"

// HandleError outputs error into terminal
func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
