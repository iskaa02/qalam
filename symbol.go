package qalam

import (
	"regexp"
	"strconv"
	"strings"
)

func findEmoji(s string) string {
	emojiT := regexp.MustCompile(emojiTagRegex)
	return string(emojiT.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		key := string(b)
		key = key[2 : len(key)-2]
		unicode, ok := emojiesMap[key]
		if !ok {
			return b
		}
		r, err := unquoteCodePoint(unicode)
		if err != nil {
			return b
		}

		return []byte(string(r))
	}))
}
func unquoteCodePoint(s string) (rune, error) {
	// 16 specifies hex encoding
	// 32 is size in bits of the rune type
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\u"), 16, 32)
	return rune(r), err
}
