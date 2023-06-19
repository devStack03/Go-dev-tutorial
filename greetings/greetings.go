package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Empty name")
	}

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	// A map to associate names with messages
	messages := make(map[string]string);
	// Loop through the received slice of names, calling
	// the Helloes function to get a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with 
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
// Return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}