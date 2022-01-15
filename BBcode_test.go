package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitTags(t *testing.T) {
	args := []splittedTxt{
		{t: "normal case", code: "red"},
		{t: "with no Formating"},
		{t: "with nested tag [red]nested[/red]", code: "b"},
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
	splits := splitTags(testTxt)
	fmt.Println(splits, args)
	require.Equal(t, args, splits)
}

func TestFormat(t *testing.T) {}
