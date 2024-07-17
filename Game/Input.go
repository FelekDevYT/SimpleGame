package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	if rl.IsKeyPressed(rl.KeyW) {
		PLAYER_Y -= CUBE_SIZE
		setNewLog("player change position to UP")
	} else if rl.IsKeyPressed(rl.KeyD) {
		PLAYER_X += CUBE_SIZE
		setNewLog("player change position to RIGHT")
	} else if rl.IsKeyPressed(rl.KeyA) {
		PLAYER_X -= CUBE_SIZE
		setNewLog("player change position to LEFT")
	} else if rl.IsKeyPressed(rl.KeyS) {
		PLAYER_Y += CUBE_SIZE
		setNewLog("player change position to DOWN")
	} else if rl.IsKeyPressed(rl.KeyE) {
		if isOpenedInventory {
			isOpenedInventory = false
		} else {
			isOpenedInventory = true
		}
	} else if rl.IsKeyPressed(rl.KeyOne) {
		current_slot = 0
	} else if rl.IsKeyPressed(rl.KeyTwo) {
		current_slot = 1
	} else if rl.IsKeyPressed(rl.KeyThree) {
		current_slot = 2
	} else if rl.IsKeyPressed(rl.KeyFour) {
		current_slot = 3
	} else if rl.IsKeyPressed(rl.KeyFive) {
		current_slot = 4
	} else if rl.IsKeyPressed(rl.KeySix) {
		current_slot = 5
	} else if rl.IsKeyPressed(rl.KeySeven) {
		current_slot = 6
	} else if rl.IsMouseButtonPressed(rl.MouseLeftButton) { //153 481
		if isOpenedInventory {
			if rl.GetMouseX() >= 360 && rl.GetMouseX() <= 410 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 0
				current_slot++
			} else if rl.GetMouseX() >= 410 && rl.GetMouseX() <= 460 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 1
				current_slot++
			} else if rl.GetMouseX() >= 460 && rl.GetMouseX() <= 510 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 2
				current_slot++
			} else if rl.GetMouseX() >= 510 && rl.GetMouseX() <= 560 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 3
				current_slot++
			} else if rl.GetMouseX() >= 560 && rl.GetMouseX() <= 610 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 4
				current_slot++
			} else if rl.GetMouseX() >= 610 && rl.GetMouseX() <= 660 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 5
				current_slot++
			} else if rl.GetMouseX() >= 660 && rl.GetMouseX() <= 710 && rl.GetMouseY() >= SCREEN_HEIGHT-380 {
				inventory[current_slot] = 6
				current_slot++
			}
		} else {
			x := rl.GetMouseX()
			y := rl.GetMouseY()

			cube_x := x / CUBE_SIZE
			cube_y := y / CUBE_SIZE

			MAP[cube_x][cube_y] = inventory[current_slot]

			setNewLog("Set new block at an pos")
		}
	}

	if current_slot == 0 {
		renderCubeInInventory(380, SCREEN_HEIGHT-80, getInventoryColor(inventory[0]))
	} else if current_slot == 1 {
		renderCubeInInventory(430, SCREEN_HEIGHT-80, getInventoryColor(inventory[1]))
	} else if current_slot == 2 {
		renderCubeInInventory(480, SCREEN_HEIGHT-80, getInventoryColor(inventory[2]))
	} else if current_slot == 3 {
		renderCubeInInventory(530, SCREEN_HEIGHT-80, getInventoryColor(inventory[3]))
	} else if current_slot == 4 {
		renderCubeInInventory(580, SCREEN_HEIGHT-80, getInventoryColor(inventory[4]))
	} else if current_slot == 5 {
		renderCubeInInventory(630, SCREEN_HEIGHT-80, getInventoryColor(inventory[5]))
	} else if current_slot == 6 {
		renderCubeInInventory(680, SCREEN_HEIGHT-80, getInventoryColor(inventory[6]))
	}

}
