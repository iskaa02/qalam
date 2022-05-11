package styler

import "github.com/iskaa02/qalam/internal/styles"
func (s *Styler) Remove(keys... string)*Styler{
  for _,key:=range keys{
    // hex code
    if key[0] == '#'{
      rgb,err:=styles.Hex2rgb(key[1:])
      if err!= nil{ continue }
      // Find and Remove 38 2 red green blue
      s.styles =removeElement(s.styles,[]uint{38,2,rgb.Red,rgb.Green,rgb.Blue})
    }
    // background hex 
    if key[:3] == "bg#"{
      rgb,err:=styles.Hex2rgb(key[3:])
      if err!= nil{ continue }
      // Find and Remove 38 2 red green blue
      s.styles =removeElement(s.styles,[]uint{48,2,rgb.Red,rgb.Green,rgb.Blue})
    }
    if style,ok:=styles.Styles[key];ok {
      s.styles =removeElement(s.styles,[]uint{style})
    }
  }
  return s
}
func removeElement(styles []uint,elementsToRemove []uint)[]uint{
  if len(elementsToRemove)>1{
    for i:=range styles{
       if styles[i]==elementsToRemove[0]&&styles[i+1]==elementsToRemove[1]{
         return append(styles[:i],styles[i+len(elementsToRemove):]...)
       } 
    } 
  }
  for i,style:= range styles{
    if style==elementsToRemove[0]{
      return append(styles[:i],styles[i+1:]...)
    }
  }
  return styles
}
