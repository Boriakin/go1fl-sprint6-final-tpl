package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConverterMorse(message string) (string, error) {
	if message == "" {
		return "", errors.New("message is empty")
	}

	if isMorse(message) {
		return morse.ToText(message), nil
	}

	return morse.ToMorse(message), nil
}

func isMorse(m string) bool {
	for _, r := range m {
		switch r {
		case '.', '-', ' ':
			continue
		default:
			return false
		}
	}
	return len(m) > 0
}
