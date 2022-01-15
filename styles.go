package main

import "strconv"

type rgb struct {
	red   uint
	green uint
	blue  uint
}

var emojisMap map[string]string = map[string]string{
	"vampire": "u/2344",
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
	"blink":         5,
	"inverted":      7,
	"strikethrough": 9,
	"white":         37,
	"bg.white":      47,
	"black":         30,
	"bg.black":      40,
	"red":           31,
	"bg.red":        41,
}

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
