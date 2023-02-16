package fault

import (
	"log"
)

// HandleFatalError logs error string if it exists and stops program to run.
func HandleFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HandleError logs error string if it exists.
func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
