package styler

import (
	"fmt"

	"github.com/iskaa02/qalam/internal/styles"
)

type Styler struct {
  styles []uint
}
func (s *Styler) Bold()*Styler{
  s.styles = append(s.styles, 1)
  return s
}
func (s *Styler) Dim()*Styler{
  s.styles = append(s.styles, 2)
  return s
}
func (s *Styler) Italic()*Styler{
  s.styles = append(s.styles, 3)
  return s
}
func (s *Styler) Underline()*Styler{
  s.styles = append(s.styles, 4)
  return s
}
func (s *Styler) Blink()*Styler{
  s.styles = append(s.styles, 5)
  return s
}
// Flip foreground an background colors
func (s *Styler) Invert()*Styler{
  s.styles = append(s.styles, 7)
  return s
}
func (s *Styler) Strikethrough()*Styler{
  s.styles = append(s.styles, 9)
  return s
}
// foreground Colors
func (s *Styler) Red()*Styler{
  s.styles = append(s.styles, 31)
  return s
}
func (s *Styler) Green()*Styler{
  s.styles = append(s.styles, 32)
  return s
}
func (s *Styler) Yellow()*Styler{
  s.styles = append(s.styles, 33)
  return s
}
func (s *Styler) Blue()*Styler{
  s.styles = append(s.styles, 34)
  return s
}
func (s *Styler) Magenta()*Styler{
  s.styles = append(s.styles, 35)
  return s
}
func (s *Styler) Cyan()*Styler{
  s.styles = append(s.styles, 36)
  return s
}

func (s *Styler) Hex(hex string)*Styler{
  if(hex[0]=='#'){
    hex = hex[1:]
  }
  c,err:=styles.Hex2rgb(hex)
  if err!= nil{
    return s
  }
  s.styles = append(s.styles, 38, 2, c.Red, c.Green, c.Blue)
  return s
}
// Background Color

func (s *Styler) RedBackground()*Styler{
  s.styles = append(s.styles, 41)
  return s
}
func (s *Styler) GreenBackground()*Styler{
  s.styles = append(s.styles, 42)
  return s
}
func (s *Styler) YellowBackground()*Styler{
  s.styles = append(s.styles, 43)
  return s
}
func (s *Styler) BlueBackground()*Styler{
  s.styles = append(s.styles, 44)
  return s
}
func (s *Styler) MagentaBackground()*Styler{
  s.styles = append(s.styles, 45)
  return s
}
func (s *Styler) CyanBackground()*Styler{
  s.styles = append(s.styles, 46)
  return s
}
func (s *Styler) HexBackground(hex string)*Styler{
  if(hex[0]=='#'){
    hex = hex[1:]
  }
  c,err:= styles.Hex2rgb(hex)
  if err!= nil{
    return s
  }
  s.styles = append(s.styles, 48, 2, c.Red, c.Green, c.Blue)
  return s
}

func (s *Styler) Printf(format string,a ...interface{}){
  fmt.Printf(styles.ApplyStyles(format,s.styles),a...)
}

func (s *Styler) Sprintf(format string,a ...interface{})string{
  return fmt.Sprintf(styles.ApplyStyles(format,s.styles),a...)
}

