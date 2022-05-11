package styles


var Styles map[string]uint = map[string]uint{
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
	"s":             9,
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

func escapeANSI(s string) string {
	escapeCode := "\033"
	return escapeCode + s + escapeCode + "[0m"
}
