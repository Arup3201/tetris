package main

import (
	"log"

	"github.com/Arup3201/tetris/gui"
)

func main() {
	windowWidth, windowHeight := 480, 960
	game := gui.CreateGame(windowWidth, windowHeight, 0, 0)

	if err := game.Run("Tetris"); err != nil {
		log.Fatal(err)
	}
}
