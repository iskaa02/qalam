package main

import (
	"fmt"
	"regexp"
)

func findEmoji(s string) string {
	emojiT := regexp.MustCompile(emojiTag)
	return string(emojiT.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		key := string(b)
		key = key[2 : len(key)-2]
		unicode := emojisMap[key]
		fmt.Println(key)
		return []byte(unicode)
	}))
}
