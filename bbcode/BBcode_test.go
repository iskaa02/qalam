package bbcode

import (
	"fmt"
	"testing"

	"github.com/iskaa02/qalam/internal/styles"
	"github.com/stretchr/testify/require"
)

func TestSplitTags(t *testing.T) {
	args := []splittedTxt{
		{t: "normal case", code: "red"},
		{t: "with no Formating"},
		{t: "with nested tag [red]nested[/red] a", code: "b"},
		{t: "[red]with unclosed tag [?]and with unknown chars[/?]"},
		{t: "with multiple codes", code: "i b red bg.white strikethrough"},
	}
	testTxt := ""
	for _, arg := range args {
		if arg.code == "" {
			testTxt += arg.t
			continue
		}
		testTxt += fmt.Sprintf("[%v]%v[/%v]", arg.code, arg.t, arg.code)
	}
	fmt.Println(testTxt)
	splits := splitTags(testTxt)
	require.Equal(t, args, splits)
}

func TestFormat(t *testing.T) {
	expected := "[red] red only [bold]red and bold[/bold][i]red and italic[/i][/red]"
	actual :=styles.ApplyStyles(" red only ", []uint{styles.Styles["red"]})
	actual +=styles.ApplyStyles("red and bold", []uint{styles.Styles["red"], styles.Styles["bold"]})
	actual += styles.ApplyStyles("red and italic", []uint{styles.Styles["red"], styles.Styles["i"]})
	expected = format(&formattedTxt{t: expected})
	require.Equal(t, expected, actual)
}
