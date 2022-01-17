package qalam

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnquoteCodePoint(t *testing.T) {
	r, err := unquoteCodePoint("1f9db")
	require.NoError(t, err)
	require.Equal(t, '🧛', r)
}
func TestFindEmoji(t *testing.T) {
	actual := "this is a vampire [:vampire:]"
	expected := "this is a vampire 🧛"
	actual = findEmoji(actual)
	require.Equal(t, actual, expected)
}
