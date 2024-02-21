package iteration

import (
	"strings"
)

func Repeat(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatUpper(character string, repeatCount int) string {
	string := Repeat(character, repeatCount)
	return strings.ToUpper(string)
}

func RepeatLower(character string, repeatCount int) string {
	string := Repeat(character, repeatCount)
	return strings.ToLower(string)
}

func RepeatTrim(character string, repeatCount int) string {
	string := Repeat(character, repeatCount)
	return strings.Trim(string, "o")
}
