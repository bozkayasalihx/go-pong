package paints

import (
	"strings"

	"github.com/nsf/termbox-go"
)



var wall string = `|||||
                   |
                   |`

var font = map[rune][][]rune{
  'p': asArray(wall),
}


//NOTE: 
//{["||||"], ["|"], ["|"]}

func asArray(chars string) [][]rune {
  result := [][]rune{}
  line := []rune{}
  str := strings.TrimPrefix(chars, "\n")
  for _, c := range str {
    if c == '\n' {
      result = append(result, line)
      line = []rune{} 
    }else {
      line = append(line, c)
    }
  }
  return result
}


func DrawChar(x,y int, char rune, color termbox.Attribute) (ux int, uy int) {
  uy = y
  v := font[char]
  if v == nil {
    panic("font not found")
  }

  for _, cSlice := range v {
    ux = x
    for _, c := range cSlice {
      termbox.SetCell(ux, uy, c, color, termbox.ColorDefault )
      ux++
    }
    uy++
  }

  return ux, uy
}

