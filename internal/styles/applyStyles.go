package styles

import (
	"fmt"
	"strconv"
	"strings"
)

// getStyles return ANSI style codes each by their key
func GetStyles(AllKeys string) []uint {
        keys := strings.Split(AllKeys, " ")
        codes := []uint{}
        for _, key := range keys {
                if len(key) < 3 {
                        if v, ok := Styles[key]; ok {
                                codes = append(codes, v)
                        }
                        continue
                }
                if key[:1] == "#" {
                        c, err := Hex2rgb(key[1:])
                        if err != nil {
                                fmt.Println("cant decode hex")
                                continue
                        }
                        codes = append(codes, 38, 2, c.Red, c.Green, c.Blue)
                }
                if strings.Contains(key, "bg#") {
                        c, err := Hex2rgb(key[3:])
                        if err != nil {
                                fmt.Println("cant decode hex")
                                continue
                        }
                        codes = append(codes, 48, 2, c.Red, c.Green, c.Blue)
                }
                if v, ok := Styles[key]; ok {
                        codes = append(codes, v)
                }
        }
        return codes
}

// applyStyles takes string and the styles should be applied to it and return formatted string.
func ApplyStyles(s string, styles []uint) string {
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

func Style(txt, c string) string {
	return ApplyStyles(txt, GetStyles(c))
}
