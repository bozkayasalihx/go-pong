package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	HEIGHT = 20
	WIDTH  = 60
  FPS    = 4
)

var def = termbox.ColorDefault

type Coord struct {
  X int  
  Y int
}

type Game struct {
  myCoord  *Coord
  aiCoord  *Coord
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

func (g *Game) handleNavigation() {
  curNav := g.curNavigationIn;  
  //NOTE: it should not pass the top border
  if curNav == `"w"` && g.myCoord.Y>1{
    g.myCoord.Y--;
    //NOTE: ALSO should not pass the bottm border too
  }else if curNav == `"s"` &&  g.myCoord.Y < HEIGHT-2 {
    g.myCoord.Y++;
  }
}


func (g *Game) draw() {
  termbox.Clear(def, def)
  g.handleNavigation()
  for i:=0;i<HEIGHT;i++ {
    for j:=0;j<WIDTH;j++ {
      if j== 0 || j == WIDTH-1 || i == 0 || i== HEIGHT-1 {
        g.tbPrint(j, i, "#")
      }else if (j == WIDTH/2) {
        g.tbPrint(j,i, ":", termbox.ColorYellow)
      }else if(j == g.aiCoord.X && i == g.aiCoord.Y) {
        g.tbPrint(j,i, "|", termbox.ColorRed)
      }else if(j == g.myCoord.X && i == g.myCoord.Y) {
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
    myCoord: &Coord{X: 1, Y: HEIGHT/2},
    aiCoord: &Coord{X: WIDTH-2, Y: HEIGHT/2},
  }

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

  data := make([]byte, 0, 64) 
  game.draw()
  // go func() {
  //   for {
  //     if game.aiCoord.Y == HEIGHT-2 {
  //       game.aiCoord.Y = 0 
  //     }
  //     game.aiCoord.Y++
  //     game.draw()
  //     time.Sleep(20*time.Millisecond)
  //   }
  // }()
mainloop:
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
        break mainloop
      }

      game.curNavigationIn = c;
      for {
        ev := termbox.ParseEvent(data);
        if ev.N == 0 { break }
        curev := ev
        copy(data, data[curev.N:])
        data = data[:len(data)-curev.N]
      }
      game.draw()
		}
    time.Sleep(20*time.Millisecond)
	}
}
