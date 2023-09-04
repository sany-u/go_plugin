package hello

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// If no name was given, return an error with a message.
	if name == "" {
		return errors.New("error 001: empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
