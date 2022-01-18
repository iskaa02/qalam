package qalam

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type rgb struct {
	red   uint
	green uint
	blue  uint
}

var styles map[string]uint = map[string]uint{
	"reset":         0,
	"none":          0,
	"bold":          1,
	"b":             1,
	"gray":          2,
	"italic":        3,
	"i":             3,
	"underline":     4,
	"u":             4,
	"blink":         5,
	"inverted":      7,
	"strikethrough": 9,
	"black":         30,
	"bg.black":      40,
	"red":           31,
	"bg.red":        41,
	"green":         32,
	"bg.green":      42,
	"yellow":        33,
	"bg.yellow":     43,
	"blue":          34,
	"bg.blue":       44,
	"magenta":       35,
	"bg.magenta":    45,
	"cyan":          36,
	"bg.cyan":       46,
	"white":         37,
	"bg.white":      47,
}

// hex2rgb change hex string to rgb
func hex2rgb(hex string) (rgb, error) {
	var RGB rgb
	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return rgb{}, err
	}
	RGB = rgb{
		red:   uint(values >> 16),
		green: uint((values >> 8) & 0xFF),
		blue:  uint(values & 0xFF),
	}

	return RGB, nil
}

// getStyles return ANSI style codes each by their key
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

// applyStyles takes string and the styles should be applied to it and return formatted string.
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

// escapeANSI add ansi escape code to the string dependent on the OS that runs on
// some OS have different escape codes for UNIX systems it's \033 while on windows it's ESC[.
func escapeANSI(s string) string {
	escapeCode := "\033"
	if runtime.GOOS == "windows" {
		escapeCode = "ESC["
	}
	return escapeCode + s + escapeCode + "[0m"
}
