package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty input")
	}

	if strings.HasPrefix(input, ".") || strings.HasPrefix(input, "-") {
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	} else {
		result := morse.ToMorse(input)
		if result == "" {
			return "", errors.New("invalid text")
		}
		return result, nil
	}
}