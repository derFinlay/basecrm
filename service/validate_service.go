package service

import (
	"errors"
	"regexp"
	"strings"
)

// Password validation Errros
var ErrInvalidPasswordLength = errors.New("Invalid password length")
var ErrLowercaseMissing = errors.New("Lowercase letter missing in password")
var ErrUppercaseMissing = errors.New("Uppercase letter missing in password")
var ErrDigitMissing = errors.New("Digit missing in password")
var ErrSpecialMissing = errors.New("Special character missing in password")

// Email validation Errros
var ErrInvalidEmail = errors.New("Invalid email")

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}

	substrings := strings.Split(email, "@")

	if len(substrings) != 2 {
		return ErrInvalidEmail
	}

	lastPart := substrings[1]

	if strings.Count(lastPart, ".") != 1 {
		return ErrInvalidEmail
	}

	return nil
}

var lowercaseRegex = regexp.MustCompile(".*[a-z].*")
var uppercaseRegex = regexp.MustCompile(".*[A-Z].*")
var digitsRegex = regexp.MustCompile(".*\\d.*")
var specialRegex = regexp.MustCompile(".*\\W.*")

func ValidatePassword(password string) (bool, []error) {
	errors := make([]error, 0)
	if len(password) < 8 || len(password) > 64 {
		errors = append(errors, ErrInvalidPasswordLength)
	}

	if !lowercaseRegex.Match([]byte(password)) {
		errors = append(errors, ErrLowercaseMissing)
	}

	if !uppercaseRegex.Match([]byte(password)) {
		errors = append(errors, ErrUppercaseMissing)
	}

	if !digitsRegex.Match([]byte(password)) {
		errors = append(errors, ErrDigitMissing)
	}

	if !specialRegex.Match([]byte(password)) {
		errors = append(errors, ErrSpecialMissing)
	}
	if len(errors) > 0 {
		return false, errors
	}
	return true, errors
}

func ConverEmailToSafeEmail(email string) (string, error) {
	output := strings.ToLower(email)
	if strings.HasSuffix(output, "@gmail.com") {
		output = strings.Replace(output, ".", "", -1)
	}

	return strings.ToLower(email), nil
}
