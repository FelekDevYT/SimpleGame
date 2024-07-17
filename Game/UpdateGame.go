package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.SkyBlue)

	if current_slot >= 7 {
		current_slot = 0
	}

	drawMap()

	drawInventory()

	input()

	if PLAYER_Y >= int(SCREEN_HEIGHT-100) {
		PLAYER_Y = int(SCREEN_HEIGHT - 200)
	}

	renderPlayer()

	if isOpenedInventory {
		rl.DrawRectangle(350, SCREEN_HEIGHT-400, 400, 300, rl.Black)
		for i := 0; i < 7; i++ {
			renderCube(rune(360+(CUBE_SIZE*i)), SCREEN_HEIGHT-380, getInventoryColor(rune(i)))
		}
	}

	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 10, 10, 40, rl.DarkGray)

	rl.EndDrawing()
}
