package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type BLOCK struct {
	ID  rune
	col rl.Color
}

type MOD_DATA_GENERATES struct {
	blocks [24]BLOCK
}

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
	MAP      [10000][10000]BLOCK
	blocks   = []BLOCK{
		{ID: 0, col: rl.Green},
		{ID: 1, col: rl.SkyBlue},
		{ID: 2, col: rl.DarkGray},
		{ID: 3, col: rl.White},
		{ID: 4, col: rl.NewColor(132, 141, 148, 255)},
		{ID: 5, col: rl.NewColor(255, 240, 15, 255)},
		{ID: 6, col: rl.NewColor(54, 171, 186, 255)},
		//{ID: 7, col: rl.Orange},
		//{ID: 8, col: rl.Black},
	}
	colors = map[string]rl.Color{
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
	iron                []rune = []rune{0, 5, 0, 5}
	gold                []rune = []rune{0, 9, 0, 9}
	diamond             []rune = []rune{0, 12, 0, 12}
	WorldName                  = "newWorld"
	IsAlredyGenerated          = false
	LuaInt                     = lua.NewState()
	ModPaths                   = [1024]string{}
	lastAddedBlockIndex        = 7
	MOD_GENERATED_DATA         = MOD_DATA_GENERATES{}
)
