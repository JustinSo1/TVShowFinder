package internal

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
