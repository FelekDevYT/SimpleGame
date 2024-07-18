package gamePart

import (
	"fmt"
	"os"
)

func Quit() {
	fmt.Println("You are quitting")
	SetNewLog("Game has closed")
	os.Exit(0)
}
