package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	openingTagRegex = `\[[a-z\.#0-9 ]+\]`
	emojiTag        = `\[:[a-z_0-9]+:\]`
)

type formattedTxt struct {
	t              string
	inheritedCodes []uint
}

func Printf(f string, a ...interface{}) {
	s := findEmoji(fmt.Sprintf(f, a...))
	fmt.Print(format(&formattedTxt{t: s}))
}
func applyStyles(s string, styles []uint) string {
	styleCodes := []string{"["}
	for i, code := range styles {
		styleCode := strconv.Itoa(int(code))
		if i != len(styles)-1 {
			styleCode += ";"
		}
		styleCodes = append(styleCodes, styleCode)
	}
	styleCodes = append(styleCodes, "m")
	return escapeANSI(strings.Join(styleCodes, "") + s)
}
func escapeANSI(s string) string {
	return "\033" + s + "\033" + "[0m"
}
func getStyles(AllKeys string) []uint {
	keys := strings.Split(AllKeys, " ")
	codes := []uint{}
	for _, key := range keys {
		if len(key) < 3 {
			if v, ok := styles[key]; ok {
				codes = append(codes, v)
			}
			continue
		}
		if key[:1] == "#" {
			c, err := hex2rgb(key[1:])
			if err != nil {
				fmt.Println("cant decode hex")
				continue
			}
			codes = append(codes, 38, 2, c.red, c.green, c.blue)
		}
		if strings.Contains(key, "bg#") {
			c, err := hex2rgb(key[3:])
			if err != nil {
				fmt.Println("cant decode hex")
				continue
			}
			codes = append(codes, 48, 2, c.red, c.green, c.blue)

		}

		if v, ok := styles[key]; ok {
			codes = append(codes, v)
		}
	}
	return codes
}
