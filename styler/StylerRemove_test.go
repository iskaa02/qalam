package styler

import (
	"testing"

	"github.com/stretchr/testify/require"
)
func TestRemoveElement(t *testing.T){
  styles:= []uint{1,4,38,2,90,90,20}

  actual:=removeElement(styles,[]uint{1})
  expected:=[]uint{4,38,2,90,90,20}

  require.Equal(t,expected,actual)

  actual=removeElement(actual,[]uint{38,2,90,90,20})
  expected=[]uint{4}
  
  require.Equal(t,expected,actual)
}
