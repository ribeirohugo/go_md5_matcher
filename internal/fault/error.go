package fault

import (
	"log"
)

// Handle Fatal Error logs error string if it exists and stops program to run.
func HandleFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Handle Error logs error string if it exists.
func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
