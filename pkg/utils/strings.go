package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PadLeft(str string, length int, padStr string) string {
	if len(str) >= length {
		return str
	}
	
	return padStr[:length-len(str)] + str
}

func TrimInput(text string, limit int) string {
	if len(text) > limit {
		return text[:limit]
	}

	return text
}

// ToTitle converts string to Title format.
func ToTitle(str string) string {
	title := cases.Title(language.English)
	return title.String(str)
}
