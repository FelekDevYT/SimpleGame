package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

func initWindowOfGame() {

	file, err := os.Create("changeLog.log")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	setNewLog("Logger has created/opened")

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)
	setNewLog("Window has successfully initialized")

	rl.SetTargetFPS(TARGET_FPS)
	setNewLog("FPS successfully set to 2048")

	createMap()
	setNewLog("Map successfully created")

}
