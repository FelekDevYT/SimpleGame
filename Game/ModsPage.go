package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"helloApp/Game/gamePart"
)

var mods [1024]string

var listViewScrollIndex int32 = 0
var listViewActive int32 = -1

func loadMods() {
	folders := getDirectories("mods")

	a := 0
	for _, folder := range folders {
		mods[a] = folder
		a += 1
	}
}

func parseListOfMods() string {
	var arr string

	for _, mod := range mods {
		arr = arr + mod + ";"
	}
	return arr
}

func ModsPage() {
	loadMods()

	listViewActive = gui.ListView(rl.Rectangle{100, 100, 350, 600}, parseListOfMods(), &listViewScrollIndex, listViewActive)

	mod := gamePart.Mod{}

	if listViewActive != -1 {
		mod = gamePart.GetModInfo(mods[listViewActive])
	}

	rl.DrawText("ID: "+mod.ID+
		"\nVersion: "+mod.Version+
		"\nCreator: "+mod.Creator+
		"\nDescription: "+mod.Description, 560, 110, 30, rl.Black)
}
