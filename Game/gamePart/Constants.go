package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

const (
	SCREEN_WIDTH  rune = 1150 //1150
	SCREEN_HEIGHT rune = 800
	SCREEN_TITLE       = "Game"
	TARGET_FPS         = 120
	CUBE_SIZE          = 50
)

var (
	api      GameAPIEvents
	PLAYER_X = 100
	PLAYER_Y = 100
	MAP      [10000][10000]rune
	colors   = map[string]rl.Color{
		"GRASS":       rl.NewColor(40, 180, 99, 255),
		"SKY":         rl.SkyBlue,
		"STONE":       rl.DarkGray,
		"SNOW":        rl.White,
		"IRON":        rl.NewColor(132, 141, 148, 255),
		"GOLD":        rl.NewColor(255, 240, 15, 255),
		"DIAMOND":     rl.NewColor(54, 171, 186, 255),
		"PLAYER":      rl.Orange,
		"TRANSPARENT": rl.Black,
	}
	isOpenedInventory bool    = false
	isOpenedMenu      bool    = false
	current_slot              = 0
	inventory         [7]rune = [7]rune{
		7, 7, 7, 7, 7, 7, 7,
	}
	iron              []rune = []rune{0, 5, 0, 5}
	gold              []rune = []rune{0, 9, 0, 9}
	diamond           []rune = []rune{0, 12, 0, 12}
	WorldName                = "newWorld"
	IsAlredyGenerated        = false
	LuaInt                   = lua.NewState()
	ModPaths                 = [1024]string{}
)
