package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hello returns a map that associates each of the named person.
// with greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to accociated names with message
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you,%v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
