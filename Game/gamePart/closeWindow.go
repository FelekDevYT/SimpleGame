package gamePart

import (
	"fmt"
	"os"
)

func Quit() {
	fmt.Println("You are quitting")
	SetNewLog("Game has closed")

	//END LUA

	for _, folder := range ModPaths {
		LuaInt.DoFile(folder + "\\end.lua")
	}

	os.Exit(0)
}
