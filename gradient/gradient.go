package gradient

import (
	"fmt"
	"strings"

	qalam "github.com/iskaa02/qalam/internal/styles"
	"github.com/mazznoer/colorgrad"
)

type BlendMode int

// NewGradient take any valid css color
func NewGradient(c ...string) (Gradient, error) {
	return NewGradientBuilder().HtmlColors(c...).Build()
}

// Summer pre-defiend gradient [#fdbb2d,#22c1c3]
func Summer(otherStyle ...string) Gradient {
	v, _ := NewGradientBuilder().
		HtmlColors("#fdbb2d", "#22c1c3").
		ANSIstyle(otherStyle...).
		Build()
	return v
}

// Vice Pre-defiend gradient [#5ee7df,#b490ca]
func Vice(otherStyle ...string) Gradient {
	v, _ := NewGradientBuilder().
		HtmlColors("#5ee7df", "#b490ca").
		ANSIstyle(otherStyle...).
		Build()
	return v
}

// Rainbow Pre-defiend gradient Rainbow
func Rainbow(otherStyle ...string) Gradient {
	return Gradient{
		gradient: colorgrad.Rainbow(),
		styles:   otherStyle,
	}
}

type Gradient struct {
	gradient colorgrad.Gradient
	styles   []string
}

// Apply Gradinet to provided string
func (g Gradient) Apply(s string) string {
	return applyGradient(s, g.gradient, g.styles...)
}

// Apply Gradient to string and print it
func (g Gradient) Print(s string) {
	fmt.Print(g.Apply(s))
}

// Mutline Apply gradient to mutline string
func (g Gradient) Mutline(s string) string {
	result := []string{}
	max := 0
	lines := strings.Split(s, "\n")
	// space should not be skipped if otherstyles includes underline or strikethrough
	shouldSkipSpace := true
	for _, s := range lines {
		if doesIncludeStyle(s, "u", "underline", "strikethrough", "s") {
			shouldSkipSpace = true
		}
		if shouldSkipSpace && max < len(s) {
			max = len(s)
		} else if max < len(strings.ReplaceAll(s, " ", ""))-1 {
			max = len(strings.ReplaceAll(s, " ", "")) - 1
		}
	}

	if max <= 0 {
		return s
	}
	for _, s := range lines {
		res := ""
		i := float64(0)
		diff := 1 / float64(max)
		for _, char := range s {
			if char == ' ' && shouldSkipSpace {
				styles := strings.Join(g.styles, " ")
				res += qalam.Style(string(char), styles)
				continue
			}
			hex := g.gradient.At(i).Hex()
			styles := strings.Join(g.styles, " ") + " " + hex
			res += qalam.Style(string(char), styles)
			i += diff
		}
		result = append(result, res)
	}
	return strings.Join(result, "\n")
}

func applyGradient(s string, grad colorgrad.Gradient, otherStyle ...string) string {
	res := ""
	max := 0
	// space should not be skipped if otherstyles includes underline or strikethrough
	shouldSkipSpace := doesIncludeStyle(res, "u", "underline", "strikethrough", "s")
	if shouldSkipSpace {
		max = len(s)
	} else {
		max = len(strings.ReplaceAll(s, " ", "")) - 1
	}
	if max <= 0 {
		return s
	}
	i := float64(0)
	diff := 1 / float64(max)
	for _, char := range s {
		if char == ' ' && shouldSkipSpace {
			styles := strings.Join(otherStyle, " ")
			res += qalam.Style(string(char), styles)
			continue
		}
		hex := grad.At(i).Hex()
		styles := strings.Join(otherStyle, " ") + " " + hex
		res += qalam.Style(string(char), styles)
		i += diff
	}
	return res
}

func doesIncludeStyle(styles string, includedStyle ...string) bool {
	brokedStyle := strings.Split(styles, " ")
	for _, style := range includedStyle {
		for _, s := range brokedStyle {
			if s == style {
				return true
			}
		}
	}
	return false
}
