package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"helloApp/Game/gamePart"
	"helloApp/Game/logoPart"
	"os"
)

func LoadScreen() {
	rl.InitWindow(gamePart.SCREEN_WIDTH, gamePart.SCREEN_HEIGHT, gamePart.SCREEN_TITLE)
	rl.SetTargetFPS(gamePart.TARGET_FPS)

	gamePart.CreateMap()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		gamePart.DrawMap()
		rl.DrawText("Game by FelsStudio 0.0.1", 15, 100, 90, rl.Black)

		logoPart.UpdateLogoScreen()

		switch logoPart.Input() {
		case 1:
			rl.CloseWindow()
			gameScreen()
			break
		case 2:
			rl.CloseWindow()
			os.Exit(0)
		}

		rl.EndDrawing()
	}

	gamePart.MAP = [10000][10000]rune{}

}

func gameScreen() {
	gamePart.InitWindowOfGame()
	gamePart.SetNewLog("Game has successfully created")

	gamePart.SetNewLog("Game has going to update part")

	for !rl.WindowShouldClose() {
		gamePart.Update()
	}

	gamePart.SetNewLog("Game has been closed")

	gamePart.Quit()
}
