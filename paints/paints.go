package paints

import (
  "strings"
)



var Wall string = `|||||
                   |
                   |`


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

