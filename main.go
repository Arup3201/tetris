package main

import (
	"log"

	"github.com/Arup3201/tetris/gui"
)

func main() {
	gridRows, gridCols := 20, 10
	windowWidth, windowHeight := 640, 480
	game := gui.CreateGame(gridRows, gridCols, windowWidth, windowHeight, 0, 0)

	if err := game.Run("Tetris"); err != nil {
		log.Fatal(err)
	}
}
