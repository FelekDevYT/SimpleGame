package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

func Update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.SkyBlue)

	if current_slot >= 7 {
		current_slot = 0
	}

	DrawMap()

	drawInventory()

	playerReffer()

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

	if isOpenedMenu {
		rl.DrawRectangle(SCREEN_WIDTH/4+15, 300, 500, 100, rl.Black)
		rl.DrawText("Quit from the game", SCREEN_WIDTH/4+90, 330, 40, rl.White)
	}

	rl.EndDrawing()

	//UPDATE LUA

	for _, folder := range ModPaths {
		LuaInt.DoFile(folder + "\\run.lua")
	}

	//CLEARING UP CHASE OF EVENTS

	clearAndSetDisable(&api)
}

func playerReffer() {
	if PLAYER_X <= 0-CUBE_SIZE {
		PLAYER_X = PLAYER_X + CUBE_SIZE
	} else if PLAYER_X >= int(SCREEN_WIDTH) {
		PLAYER_X = PLAYER_X - CUBE_SIZE
	} else if PLAYER_Y >= int(SCREEN_HEIGHT) {
		PLAYER_Y = PLAYER_Y - CUBE_SIZE
	} else if PLAYER_Y <= 0+CUBE_SIZE {
		PLAYER_Y = PLAYER_Y + CUBE_SIZE
	}
}
