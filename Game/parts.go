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
		case 2:
			rl.CloseWindow()
			os.Exit(0)
		case 3:
			gamePart.MAP = [10000][10000]rune{}
			rl.CloseWindow()
			aboutPart()
		}

		rl.EndDrawing()
	}

	gamePart.MAP = [10000][10000]rune{}

}

func aboutPart() {
	gamePart.InitWindowOfGame()

	gamePart.CreateMap()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		gamePart.DrawMap()

		rl.DrawText("Game by FelsStudio 0.0.1", 15, 100, 90, rl.Black)
		rl.DrawText("Game has first created in 18.07.2024", 250, 200, 40, rl.Black)
		rl.DrawText("Creators:", 250, 250, 40, rl.Black)
		rl.DrawText("Felek_", 280, 300, 40, rl.Black)
		rl.DrawText("Thanks by use this simple game", 250, 400, 40, rl.Black)

		rl.DrawRectangle(250, 450, 500, 100, rl.Black)
		rl.DrawText("Back to lobby", 350, 480, 40, rl.White)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if rl.GetMouseX() >= 250 && rl.GetMouseX() <= 750 && rl.GetMouseY() >= 450 && rl.GetMouseY() <= 550 {
				gamePart.MAP = [10000][10000]rune{}
				rl.CloseWindow()
				LoadScreen()
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
	os.Exit(0)
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
