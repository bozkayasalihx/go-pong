package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	HEIGHT = 30
	WIDTH  = 30
)

var def = termbox.ColorDefault

type Game struct {
  myPositionY int 
  myPositionX int 
  aiPositionX int 
  aiPositionY int 
  curNavigationIn   string
}


func (g *Game) navigationPrint(x,y int) {
  for _, val := range g.curNavigationIn{
    termbox.SetCell(x,y, val, termbox.ColorBlue, termbox.ColorDefault)  
    x++
  }
}

func (g *Game) tbPrint(x,y int, msg string, attr ...termbox.Attribute) {
  var color termbox.Attribute
  if len(attr) > 1 {
    log.Fatalf("want one termbox.Attribute but got %v", attr)
    os.Exit(2);
  }
  if len(attr) == 1 {
    color = attr[0]
  }else {
    color = termbox.ColorBlue;
  }
  for _, val := range msg {
    termbox.SetCell(x,y, val, color, termbox.ColorDefault)
    x++
  }
}


func (g *Game) draw(){
  termbox.Clear(def, def)
  for i:=0;i<HEIGHT;i++ {
    for j:=0;j<WIDTH;j++ {
      if j== 0 || j == WIDTH-1 || i == 0 || i== HEIGHT-1 {
        g.tbPrint(j, i, "#")
      }else if(j == g.myPositionX && i == g.myPositionY) {
        g.tbPrint(j,i, "|", termbox.ColorGreen)
      }else {
        g.tbPrint(j,i, " ")
      }
    }
  }
  termbox.Flush()
}



func main() {
	err := termbox.Init()
  game := &Game{
    myPositionY: HEIGHT /2,   
    myPositionX: 1, 
    aiPositionY: HEIGHT /2,
    aiPositionX: WIDTH -1,
  }

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

  data := make([]byte, 0, 64) 
  game.draw()
loop:
	for {
    if cap(data)-len(data) < 32 {
      newData  := make([]byte, len(data), len(data)+32)
      copy(newData, data)
      data = newData
    }
    lenData := len(data)

    d := data[lenData:lenData+32] 

		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
      data = data[:lenData+ev.N]
      c := fmt.Sprintf("%q", data)
      if c == `"q"` {
        break loop
      }

      game.curNavigationIn =c;
      for {
        ev := termbox.ParseEvent(data);
        if ev.N == 0 { break }
        curev := ev
        copy(data, data[curev.N:])
        data = data[:len(data)-curev.N]
      }
      game.draw()
		}
	}
}
