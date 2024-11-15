package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
	"os"
	"strconv"
)

func InitWindowOfGame(wn string) {
	WorldName = wn

	SetNewLog("Setting up logger")
	file, _ := os.Create("changeLog.log")

	defer file.Close()
	SetNewLog("Logger has created/opened")

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE)
	SetNewLog("Window has successfully initialized")

	SetNewLog("Setting up window icon...")
	img := rl.LoadImage("assets/logo.png")
	rl.SetWindowIcon(*img)
	SetNewLog("Successfully loaded logo.png")

	rl.SetTargetFPS(TARGET_FPS)
	SetNewLog("FramesPerSecond set to " + strconv.Itoa(int(TARGET_FPS)))

	if !IsAlredyGenerated {
		CreateMap()
		SetNewLog("Map successfully created")
	} else {
		SetNewLog("World already opened")
	}

	SetNewLog("Setting up Lua Interpreter...")
	LuaInt = lua.NewState()
	LuaInt.PreloadModule("api", Loader)

	SetNewLog("Setting default value to EVENTS")
	clearAndSetDisable(&api)

	api.WorldEvent.isWorldAlreadyOpened = true
}
