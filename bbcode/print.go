package bbcode

import (
	"fmt"
)

const (
	// openingTagRegex REGEX for finding bbcode opening tags
	openingTagRegex = `\[[a-z\.#0-9 ]+\]`
	// emojiTagRegex REGEX for finding emoji tags
	emojiTagRegex = `\[:[a-z_0-9]+:\]`
)

// Printf takes a BBcode string and Print the formatted text.
func Printf(f string, a ...interface{}) {
	fmt.Print(Sprintf(f, a...))
}

// Sprintf takes a BBcode string and return the formatted text.
func Sprintf(f string, a ...interface{}) string {
	s := fmt.Sprintf(f, a...)
	return format(&formattedTxt{t: s})
}

// Style format text with provided codes.
// takes 2 strings first is the txt the will be styled and,
// the second is the code of the styles (multiple codes should be separated with space).
