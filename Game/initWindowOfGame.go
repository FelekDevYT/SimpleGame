package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"strconv"
)

func initWindowOfGame() {

	setNewLog("Setting up logger")
	file, err := os.Create("changeLog.log")
	if err != nil {
		println(err.Error())
		setNewLog("Error of setting up logger")
		os.Exit(1)
	}
	defer file.Close()
	setNewLog("Logger has created/opened")

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)
	setNewLog("Window has successfully initialized")

	rl.SetTargetFPS(TARGET_FPS)
	setNewLog("FramesPerSecond set to " + strconv.Itoa(int(TARGET_FPS)))

	createMap()
	setNewLog("Map successfully created")

}
