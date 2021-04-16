package fault

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

const (
	errorExample = "error example 1234"
)

func TestHandleError(t *testing.T) {
	err := errors.New(errorExample)

	var str bytes.Buffer
	log.SetOutput(&str)

	HandleError(err)

	if !strings.Contains(str.String(), errorExample) {
		t.Errorf("Error output incorrect, got: %s, may contain: %s.", str.String(), errorExample)
	}
}
