package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/martinlindhe/inputbox"
	"github.com/tawesoft/golib/v2/dialog"
	"helloApp/Game/gamePart"
	"helloApp/Game/logoPart"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getDirectories(dirPath string) []string {
	var folders []string

	entries, _ := ioutil.ReadDir(dirPath)

	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, filepath.Join(dirPath, entry.Name()))
		}
	}

	return folders
}

/*
	ModsButton = gui.Button(rl.Rectangle{float32(gamePart.SCREEN_WIDTH/4 + 15 + 500 + 20), 600, 100, 100}, "Mods")

	PlayButton = gui.Button(rl.Rectangle{float32(gamePart.SCREEN_WIDTH/4 + 15), 300, 500, 100}, "Play")
	AboutButton = gui.Button(rl.Rectangle{float32(gamePart.SCREEN_WIDTH/4 + 15), 450, 500, 100}, "About")
	QuitButton = gui.Button(rl.Rectangle{float32(gamePart.SCREEN_WIDTH)/4 + 14, 600, 500, 100}, "Quit")
*/

func LoadScreen() {

	folders := getDirectories("mods")

	a := 0
	for _, folder := range folders {
		gamePart.ModPaths[a] = folder
	}

	gamePart.SetNewLog("Setting up mods folders")

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
			gamePart.MAP = [10000][10000]rune{}
			rl.CloseWindow()
			selectWorldType()
		case 2:
			rl.CloseWindow()
			os.Exit(0)
		case 3:
			gamePart.MAP = [10000][10000]rune{}
			rl.CloseWindow()
			aboutPart()
		case 4:
			gamePart.MAP = [10000][10000]rune{}
			rl.CloseWindow()
			ModsPagePart()
		}

		rl.EndDrawing()
	}

	gamePart.MAP = [10000][10000]rune{}

}

func ModsPagePart() {
	gamePart.InitWindowOfGame("world")

	gamePart.CreateMap()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		gamePart.DrawMap()

		rl.DrawRectangle(550, 100, 500, 600, rl.White)

		ModsPage()

		rl.EndDrawing()
	}
}

func selectWorldType() {
	gamePart.InitWindowOfGame("world")

	gamePart.CreateMap()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		gamePart.DrawMap()

		rl.DrawText("Select the option", 180, 100, 90, rl.Black)

		rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 300, 500, 100, rl.Black)
		rl.DrawText("New world", gamePart.SCREEN_WIDTH/4+220, 330, 40, rl.White)

		rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 450, 500, 100, rl.Black)
		rl.DrawText("Open world", gamePart.SCREEN_WIDTH/4+220, 480, 40, rl.White)

		rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 600, 500, 100, rl.Black)
		rl.DrawText("Back", gamePart.SCREEN_WIDTH/4+220, 630, 40, rl.White)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			if rl.GetMouseX() >= gamePart.SCREEN_WIDTH/4+15 && rl.GetMouseX() <= gamePart.SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 300 && rl.GetMouseY() <= 400 && rl.GetMouseY() <= 400 {
				//CREATING NEW LEVEL
				name, ok := inputbox.InputBox("Select world name", "Enter valid world name", "new world")

				if ok {
					rl.CloseWindow()
					gameScreen(name)
				} else {
					dialog.Alert("Enter valid name and press 'ok' button")
				}
			} else if rl.GetMouseX() >= gamePart.SCREEN_WIDTH/4+15 && rl.GetMouseX() <= gamePart.SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 450 && rl.GetMouseY() <= 550 {
				//OPENING LEVEL
				name, ok := inputbox.InputBox("Enter world name", "Enter valid world name", "world#1")

				if ok {
					gamePart.IsAlredyGenerated = true
					gamePart.WorldName = name
					gamePart.OpenSaveOfLevel()
					rl.CloseWindow()
					gameScreen(name)
				} else {
					dialog.Alert("Enter valid name and press 'ok' button")
				}
			} else if rl.GetMouseX() >= gamePart.SCREEN_WIDTH/4+15 && rl.GetMouseX() <= gamePart.SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 600 && rl.GetMouseY() <= 700 {
				gamePart.MAP = [10000][10000]rune{}
				rl.CloseWindow()
				LoadScreen()
			}
		}

		rl.EndDrawing()
	}
}

// name, ok := inputbox.InputBox("Select world name", "Enter valid world name", "new world")
//
// if ok {
// gamePart.MAP = [10000][10000]rune{}
// rl.CloseWindow()
// gameScreen(name)
// } else {
// dialog.Alert("Enter valid name and press 'ok' button")
// }
func aboutPart() {
	gamePart.InitWindowOfGame("world")

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

func gameScreen(wn string) {

	gamePart.InitWindowOfGame(wn)
	gamePart.SetNewLog("Game has successfully created")

	gamePart.SetNewLog("Game has going to update part")

	defer gamePart.LuaInt.Close()

	for !rl.WindowShouldClose() {
		gamePart.Update()
	}

	gamePart.SetNewLog("Game has been closed")

	gamePart.Quit()
}
