package main

import (
	"fmt"
	"os"
)

func quit() {
	fmt.Println("You are quitting")
	setNewLog("Game has closed")
	os.Exit(0)
}
