package qalam

import (
	"fmt"
)

const (
	openingTagRegex = `\[[a-z\.#0-9 ]+\]`
	emojiTagRegex   = `\[:[a-z_0-9]+:\]`
)

func Printf(f string, a ...interface{}) {
	fmt.Print(Sprintf(f, a...))
}

func Sprintf(f string, a ...interface{}) string {
	s := fmt.Sprintf(f, a...)
	s = findEmoji(s)
	return format(&formattedTxt{t: s})
}
