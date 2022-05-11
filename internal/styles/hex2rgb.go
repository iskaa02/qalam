package styles

import "strconv"

type Rgb struct {
	Red   uint
	Green uint
	Blue  uint
}
// hex2rgb change hex string to rgb
func Hex2rgb(hex string) (Rgb, error) {
	var RGB Rgb
	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return Rgb{}, err
	}
	RGB = Rgb{ Red:   uint(values >> 16),
		Green: uint((values >> 8) & 0xFF),
		Blue:  uint(values & 0xFF),
	}

	return RGB, nil
}
