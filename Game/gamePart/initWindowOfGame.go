package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"strconv"
)

func InitWindowOfGame() {

	SetNewLog("Setting up logger")
	file, err := os.Create("changeLog.log")
	if err != nil {
		println(err.Error())
		SetNewLog("Error of setting up logger")
		os.Exit(1)
	}
	defer file.Close()
	SetNewLog("Logger has created/opened")

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)
	SetNewLog("Window has successfully initialized")

	SetNewLog("CameraManager has initialized")

	rl.SetTargetFPS(TARGET_FPS)
	SetNewLog("FramesPerSecond set to " + strconv.Itoa(int(TARGET_FPS)))

	CreateMap()
	SetNewLog("Map successfully created")

}
