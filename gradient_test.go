package qalam

import (
	"fmt"
	"testing"
)

func TestApply(t *testing.T) {
	// newGrad := colorgrad.Rainbow()
	// fmt.Println(applyGradient("short", newGrad))
	fmt.Println(Rainbow("bold", "i").Apply("A Very Long Rainbow Gradinet"))
	fmt.Println(Summer().Apply("A kinda long Vice Gradient"))
	fmt.Println(Vice().Apply("A Summer Hello veyyyyyyyy lllllong  rainbow"))
	newGrad, _ := NewGradient("#FEAC5E", "#C779D0", "#4BC0C8")
	fmt.Println(newGrad.Mutline(` $$$$$$\            $$\                         
$$  __$$\           $$ |                        
$$ /  $$ | $$$$$$\  $$ | $$$$$$\  $$$$$$\$$$$\  
$$ |  $$ | \____$$\ $$ | \____$$\ $$  _$$  _$$\ 
$$ |  $$ | $$$$$$$ |$$ | $$$$$$$ |$$ / $$ / $$ |
$$ $$\$$ |$$  __$$ |$$ |$$  __$$ |$$ | $$ | $$ |
\$$$$$$ / \$$$$$$$ |$$ |\$$$$$$$ |$$ | $$ | $$ |
 \___$$$\  \_______|\__| \_______|\__| \__| \__|
     \___|`))
}
