package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand/v2"
	"os"
	"time"
)

func renderCube(x rune, y rune, color rl.Color) {
	rl.DrawRectangle(x, y, CUBE_SIZE, CUBE_SIZE, color)
	rl.SetLineWidth(3)
	rl.DrawLine(x, y, x+CUBE_SIZE, y, rl.Gray)
	rl.DrawLine(x, y, x, y+CUBE_SIZE, rl.Gray)
	rl.DrawLine(x, y+CUBE_SIZE, x+CUBE_SIZE, y+CUBE_SIZE, rl.Gray)
	rl.DrawLine(x+CUBE_SIZE, y, x+CUBE_SIZE, y+CUBE_SIZE, rl.Gray)
}

func renderCubeInInventory(x rune, y rune, color rl.Color) {
	rl.DrawRectangle(x, y, CUBE_SIZE, CUBE_SIZE, color)
	rl.SetLineWidth(3)
	rl.DrawLine(x, y, x+CUBE_SIZE, y, rl.LightGray)
	rl.DrawLine(x, y, x, y+CUBE_SIZE, rl.LightGray)
	rl.DrawLine(x, y+CUBE_SIZE, x+CUBE_SIZE, y+CUBE_SIZE, rl.LightGray)
	rl.DrawLine(x+CUBE_SIZE, y, x+CUBE_SIZE, y+CUBE_SIZE, rl.LightGray)
}

func renderPlayer() {
	renderCube(rune(PLAYER_X), rune(PLAYER_Y), colors["PLAYER"])
}

func getIndex(idx rune) rune {
	/*
		14 = 1
		13 = 2
		12 = 3
		11 = 4
		10 = 5
		11 = 6
		... = ...
	*/
	var idxs = map[rune]rune{}

	for i := 0; i < int(SCREEN_HEIGHT/CUBE_SIZE); i++ {
		idxs[rune(i)] = rune(int(SCREEN_HEIGHT/CUBE_SIZE) - i)
	}
	return idxs[idx]

}

func getDefaultIndex(idx rune) int {
	return int(idx - 1)
}

func isWorldBorderInLeft(i rune) rune {
	if i == 1 {
		return 0
	} else {
		return i - 1
	}
}

func random(min, max int) int {
	return rand.IntN(max-min) + min
}

func isLeftBorder(i rune) bool {
	if i == 0 {
		return true
	} else {
		return false
	}
}

func drawInventory() {
	rl.DrawRectangle(0, SCREEN_HEIGHT-100, SCREEN_WIDTH, 100, rl.Black)
	startDrawing := SCREEN_WIDTH / 3
	renderCube(startDrawing, SCREEN_HEIGHT-80, getInventoryColor(inventory[0]))
	renderCube(startDrawing+(CUBE_SIZE*1), SCREEN_HEIGHT-80, getInventoryColor(inventory[1]))
	renderCube(startDrawing+(CUBE_SIZE*2), SCREEN_HEIGHT-80, getInventoryColor(inventory[2]))
	renderCube(startDrawing+(CUBE_SIZE*3), SCREEN_HEIGHT-80, getInventoryColor(inventory[3]))
	renderCube(startDrawing+(CUBE_SIZE*4), SCREEN_HEIGHT-80, getInventoryColor(inventory[4]))
	renderCube(startDrawing+(CUBE_SIZE*5), SCREEN_HEIGHT-80, getInventoryColor(inventory[5]))
	renderCube(startDrawing+(CUBE_SIZE*6), SCREEN_HEIGHT-80, getInventoryColor(inventory[6]))
}

func getInventoryColor(color rune) rl.Color {
	switch color {
	case 0:
		return colors["GRASS"]
	case 1:
		return colors["SKY"]
	case 2:
		return colors["STONE"]
	case 3:
		return colors["SNOW"]
	case 4:
		return colors["IRON"]
	case 5:
		return colors["GOLD"]
	case 6:
		return colors["DIAMOND"]
	case 7:
		return colors["TRANSPARENT"]
	default:
		panic("Not found color")
	}
}

func isOre(lvl []rune) bool {
	if random(int(lvl[0]), int(lvl[1])) == random(int(lvl[2]), int(lvl[3])) {
		return true
	}
	return false
}

/*
LOGGER - USE LOGGER FOR INFO OF GAME
*/

func setNewLog(log string) {
	log = "[" + time.Now().Format("15:15:05") + "]" + log + "\n"
	data := []byte(log)
	file, err := os.OpenFile("changeLog.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(data)
}
