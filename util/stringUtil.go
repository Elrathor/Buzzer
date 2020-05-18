package util

import (
	"regexp"
)

var regex *regexp.Regexp

func RemoveNonAlphanumericCharacters(input string) (output string) {
	if regex == nil {
		regex = regexp.MustCompile("[^a-zA-Z0-9 _]+")
	}

	output = regex.ReplaceAllString(input, "")
	return
}
