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
			if rl.GetMouseX() >= 360 && rl.GetMouseX() <= 410 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 0
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 410 && rl.GetMouseX() <= 460 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 1
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 460 && rl.GetMouseX() <= 510 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 2
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 510 && rl.GetMouseX() <= 560 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 3
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 560 && rl.GetMouseX() <= 610 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 4
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 610 && rl.GetMouseX() <= 660 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 5
				current_slot++
				SetNewLog("setting up of new block in inventory")
			} else if rl.GetMouseX() >= 660 && rl.GetMouseX() <= 710 && rl.GetMouseY() >= SCREEN_HEIGHT-380 && rl.GetMouseY() <= SCREEN_HEIGHT-330 {
				inventory[current_slot] = 6
				current_slot++
				SetNewLog("setting up of new block in inventory")
			}
		} else {
			x := rl.GetMouseX()
			y := rl.GetMouseY()

			cube_x := x / CUBE_SIZE
			cube_y := y / CUBE_SIZE

			MAP[cube_x][cube_y] = inventory[current_slot]

			SetNewLog("Set new block at: " + strconv.Itoa(int(cube_x)) + ", " + strconv.Itoa(int(cube_y)))
		}
	} else if rl.IsKeyPressed(rl.KeyM) {
		if isOpenedMenu {
			isOpenedMenu = false
		} else {
			isOpenedInventory = false
			isOpenedMenu = true
		}
	} else if rl.IsKeyPressed(rl.KeyP) {
		saveCurrentLevel("level#1.txt")
	} else if rl.IsKeyPressed(rl.KeyF2) {
		openSaveOfLevel("level#1.txt")
	}

	if current_slot == 0 {
		renderCubeInCurrentPositionOnInventory(380, SCREEN_HEIGHT-80, getInventoryColor(inventory[0]))
	} else if current_slot == 1 {
		renderCubeInCurrentPositionOnInventory(430, SCREEN_HEIGHT-80, getInventoryColor(inventory[1]))
	} else if current_slot == 2 {
		renderCubeInCurrentPositionOnInventory(480, SCREEN_HEIGHT-80, getInventoryColor(inventory[2]))
	} else if current_slot == 3 {
		renderCubeInCurrentPositionOnInventory(530, SCREEN_HEIGHT-80, getInventoryColor(inventory[3]))
	} else if current_slot == 4 {
		renderCubeInCurrentPositionOnInventory(580, SCREEN_HEIGHT-80, getInventoryColor(inventory[4]))
	} else if current_slot == 5 {
		renderCubeInCurrentPositionOnInventory(630, SCREEN_HEIGHT-80, getInventoryColor(inventory[5]))
	} else if current_slot == 6 {
		renderCubeInCurrentPositionOnInventory(680, SCREEN_HEIGHT-80, getInventoryColor(inventory[6]))
	}

}
