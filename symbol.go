package qalam

import (
	"regexp"
	"strconv"
	"strings"
)

// findEmoji replace emoji tags with their respective emoji
// this is a [:vampire:] -> this is a 🧛
func findEmoji(s string) string {
	emojiT := regexp.MustCompile(emojiTagRegex)
	return string(emojiT.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		key := string(b)
		// taking off the [::]
		key = key[2 : len(key)-2]
		unicode, ok := emojisMap[key]
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

// unquoteCodePoint takes unicode string and return rune
func unquoteCodePoint(s string) (rune, error) {
	// 16 specifies hex encoding
	// 32 is size in bits of the rune type
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\u"), 16, 32)
	return rune(r), err
}
