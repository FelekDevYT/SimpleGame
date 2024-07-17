package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	initWindowOfGame()
	setNewLog("Game has successfully created")

	setNewLog("Game has going to update part")

	for !rl.WindowShouldClose() {
		update()
	}

	setNewLog("Game has been closed")

	quit()
}
