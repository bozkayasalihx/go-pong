package paints

import (
	"strings"
	"github.com/nsf/termbox-go"
)

//NOTE:
//{["||||"], ["|"], ["|"]}

var wall string = `| 
                   |
                   |`

func asArray(chars string) [][]rune {
  var result [][]rune
  var line  []rune
  str := strings.TrimPrefix(chars, "\n")
  for _, c := range str {
    if c == '\n' {
      result = append(result, line)
      //NOTE: gained performans gains while each iteration 
      line = line[:0]
    }else {
      line = append(line, c)
    }
  }
  return result
}


func DrawChar(x,y int,color termbox.Attribute) (ux int, uy int) {
  uy = y
  v := asArray(wall)
  if v == nil {
    panic("font not found")
  }

  for _, cSlice := range v {
    ux = x
    for _, c := range cSlice {
      termbox.SetCell(ux, uy, c, color, termbox.ColorDefault)
      ux++
    }
    uy++
  }

  return ux, uy
}
