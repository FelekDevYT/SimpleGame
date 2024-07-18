package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var camera rl.Camera2D

func setUpCameraManager() {
	camera.Target = rl.Vector2{float32(PLAYER_X + 20), float32(PLAYER_Y + 20)}
	camera.Offset = rl.Vector2{float32(SCREEN_WIDTH / 2), float32(SCREEN_HEIGHT / 2)}
	camera.Rotation = 0
	camera.Zoom = 1
}
