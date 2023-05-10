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
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
        break loop;
			}
      tbprint(0, HEIGHT +10, termbox.ColorGreen, "w")
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
