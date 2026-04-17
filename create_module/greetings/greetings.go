package greetings

import (
	"fmt"
	"time"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail %v! Well met!",
	}

	randNum := r.Intn(len(formats))
	return formats[randNum]
}
