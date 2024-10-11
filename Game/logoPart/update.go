package logoPart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"helloApp/Game/gamePart"
)

func UpdateLogoScreen() {
	rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 300, 500, 100, rl.Black)
	rl.DrawText("Play", gamePart.SCREEN_WIDTH/4+220, 330, 40, rl.White)

	rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 450, 500, 100, rl.Black)
	rl.DrawText("About", gamePart.SCREEN_WIDTH/4+220, 480, 40, rl.White)

	rl.DrawRectangle(gamePart.SCREEN_WIDTH/4+15, 600, 500, 100, rl.Black)
	rl.DrawText("Quit", gamePart.SCREEN_WIDTH/4+220, 630, 40, rl.White)
}
