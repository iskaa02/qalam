package gradient

import (
	"image/color"

	"github.com/mazznoer/colorgrad"
)

const (
	BlendHcl BlendMode = iota
	BlendHsv
	BlendLab
	BlendLinearRgb
	BlendLuv
	BlendRgb
	BlendOklab
)

type GradBuilder struct {
	builder *colorgrad.GradientBuilder
	styles  []string
}

// NewGradientBuilder return Gradient Builder for more customization
func NewGradientBuilder() *GradBuilder {
	return &GradBuilder{
		builder: colorgrad.NewGradient(),
	}
}

// Mode change the belnding mode for the gradient default is BlendRgb
// available modes: BlendRgb,BlendLinearRgb,BlendHcl,BlendHsv,BlendLab,BlendLuv,BlendOklab
func (g *GradBuilder) Mode(m BlendMode) *GradBuilder {
	g.builder = g.builder.Mode(colorgrad.BlendMode(m))
	return g
}

// Colors accept anything that implement color.Color interface.
func (g *GradBuilder) Colors(c ...color.Color) *GradBuilder {
	g.builder = g.builder.Colors(c...)
	return g
}

// ANSIstyle accept ansi styles provided
func (g *GradBuilder) ANSIstyle(styles ...string) *GradBuilder {
	g.styles = styles
	return g
}

// HtmlColors accepts any valid css color including
// named colors, hexadecimal (#rgb, #rgba, #rrggbb, #rrggbbaa), rgb(), rgba(), hsl(), hsla(), hwb(), and hsv().
func (g *GradBuilder) HtmlColors(c ...string) *GradBuilder {
	g.builder = g.builder.HtmlColors(c...)
	return g
}

// Changes the ratio between colors, default: [0,1]
func (g *GradBuilder) Domain(d ...float64) *GradBuilder {
	g.builder = g.builder.Domain(d...)
	return g
}

func (g *GradBuilder) Build(c ...string) (Gradient, error) {
	grad, err := g.builder.Build()
	return Gradient{gradient: grad, styles: g.styles}, err
}
