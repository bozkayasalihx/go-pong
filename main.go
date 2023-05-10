package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	HEIGHT = 30
	WIDTH  = 30
)

var (
	myPostionX  = 1
	myPositionY = HEIGHT / 2

	aiPositionX = WIDTH - 2
	aiPositionY = HEIGHT / 2
)

func main() {
	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

  printBorders()
  data := make([]byte, 0, 64)
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
      
      for {
        ev := termbox.ParseEvent(data);
        if ev.N == 0 { break }
        curev := ev
        copy(data, data[curev.N:])
        data = data[:len(data)-curev.N]
      }

		}
    printBorders()
    logic()
	}
}


func logic() {

}

func printBorders() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if i == 0 || j == 0 || j == WIDTH-1 || i == HEIGHT-1 {
				tbprint(j, i, termbox.ColorGreen, "#")
			}
			if j == myPostionX && i == myPositionY {
				tbprint(j, i, termbox.ColorBlue, "|")
			}
			if j == aiPositionX && i == aiPositionY {
				tbprint(j, i, termbox.ColorRed, "|")
			}

		}
	}

	tbprint(0, HEIGHT+2, termbox.ColorBlue, "<ESC> for the kill")
	termbox.Flush()

}

func tbprint(x, y int, fg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, termbox.ColorDefault)
		x++
	}
}
