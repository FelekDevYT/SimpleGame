package logoPart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"helloApp/Game/gamePart"
)

func Input() rune {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		if rl.GetMouseX() >= gamePart.SCREEN_WIDTH/4+15 && rl.GetMouseX() <= gamePart.SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 300 && rl.GetMouseY() <= 400 && rl.GetMouseY() <= 400 {
			return 1
		} else if rl.GetMouseX() >= gamePart.SCREEN_WIDTH/4+15 && rl.GetMouseX() <= gamePart.SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 450 && rl.GetMouseY() <= 550 {
			return 2
		}
	}
	return 0
}
