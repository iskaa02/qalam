package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplyStyles(t *testing.T) {
	styles := []uint{1, 31, 3}
	txt := "Bold, Red And Italic"
	s := applyStyles(txt, styles)
	stylesString := ""
	for i, v := range styles {
		if i != 0 {
			stylesString += ";"
		}
		stylesString += strconv.Itoa(int(v))
	}
	require.Equal(t, s, fmt.Sprintf("\033[%sm%s\033[0m", stylesString, txt))
}
func TestGetStyles(t *testing.T) {
	keys := "b i strikethrough #000000 bg#ffffff"
	styles := getStyles(keys)
	require.Equal(t, []uint{1, 3, 9, 38, 2, 0, 0, 0, 48, 2, 255, 255, 255}, styles)
}
