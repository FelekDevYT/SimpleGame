package gamePart

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"strconv"
)

func input() {
	if rl.IsKeyPressed(rl.KeyW) {
		PLAYER_Y -= CUBE_SIZE
		SetNewLog("player change position to UP")
	} else if rl.IsKeyPressed(rl.KeyD) {
		PLAYER_X += CUBE_SIZE
		SetNewLog("player change position to RIGHT")
	} else if rl.IsKeyPressed(rl.KeyA) {
		PLAYER_X -= CUBE_SIZE
		SetNewLog("player change position to LEFT")
	} else if rl.IsKeyPressed(rl.KeyS) {
		PLAYER_Y += CUBE_SIZE
		SetNewLog("player change position to DOWN")
	} else if rl.IsKeyPressed(rl.KeyE) {
		if isOpenedInventory {
			isOpenedInventory = false
		} else {
			isOpenedInventory = true
		}
	} else if rl.IsKeyPressed(rl.KeyOne) {
		current_slot = 0
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeyTwo) {
		current_slot = 1
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeyThree) {
		current_slot = 2
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeyFour) {
		current_slot = 3
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeyFive) {
		current_slot = 4
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeySix) {
		current_slot = 5
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsKeyPressed(rl.KeySeven) {
		current_slot = 6
		SetNewLog("Set current slot to " + strconv.Itoa(int(current_slot)))
	} else if rl.IsMouseButtonPressed(rl.MouseLeftButton) { //153 481
		if isOpenedMenu {
			if rl.GetMouseX() >= SCREEN_WIDTH/4+15 && rl.GetMouseX() <= SCREEN_WIDTH/4+15+500 && rl.GetMouseY() >= 300 && rl.GetMouseY() <= 400 && rl.GetMouseY() <= 400 {
				rl.CloseWindow()
				os.Exit(0)
			}
		}
		if isOpenedInventory {
			for i := 0; i != 7; i += 1 {
				if rl.GetMouseX() >= int32(360+(50*i)) &&
					rl.GetMouseX() <= int32(410+(50*i)) &&
					rl.GetMouseY() >= SCREEN_HEIGHT-380 &&
					rl.GetMouseY() <= SCREEN_HEIGHT-330 &&
					rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					inventory[current_slot] = rune(i)
					current_slot++
					SetNewLog("Setting up of new block in inventory")
				}
			}
		} else {
			x := rl.GetMouseX()
			y := rl.GetMouseY()

			cube_x := x / CUBE_SIZE
			cube_y := y / CUBE_SIZE

			MAP[cube_x][cube_y].ID = inventory[current_slot]

			SetNewLog("Set new block at: " + strconv.Itoa(int(cube_x)) + ", " + strconv.Itoa(int(cube_y)))

			setEnable(&api, 0)
			api.isBlockSet.cords.x = cube_x
			api.isBlockSet.cords.y = cube_y
		}
	} else if rl.IsKeyPressed(rl.KeyM) {
		if isOpenedMenu {
			isOpenedMenu = false
		} else {
			isOpenedInventory = false
			isOpenedMenu = true
		}
	} else if rl.IsKeyPressed(rl.KeyP) {
		SaveCurrentLevel()
	} else if rl.IsKeyPressed(rl.KeyF2) {
		OpenSaveOfLevel()
	}

	if current_slot == 0 {
		renderCubeInCurrentPositionOfInventory(380, SCREEN_HEIGHT-80, getInventoryColor(inventory[0]))
	} else if current_slot == 1 {
		renderCubeInCurrentPositionOfInventory(430, SCREEN_HEIGHT-80, getInventoryColor(inventory[1]))
	} else if current_slot == 2 {
		renderCubeInCurrentPositionOfInventory(480, SCREEN_HEIGHT-80, getInventoryColor(inventory[2]))
	} else if current_slot == 3 {
		renderCubeInCurrentPositionOfInventory(530, SCREEN_HEIGHT-80, getInventoryColor(inventory[3]))
	} else if current_slot == 4 {
		renderCubeInCurrentPositionOfInventory(580, SCREEN_HEIGHT-80, getInventoryColor(inventory[4]))
	} else if current_slot == 5 {
		renderCubeInCurrentPositionOfInventory(630, SCREEN_HEIGHT-80, getInventoryColor(inventory[5]))
	} else if current_slot == 6 {
		renderCubeInCurrentPositionOfInventory(680, SCREEN_HEIGHT-80, getInventoryColor(inventory[6]))
	}

}
