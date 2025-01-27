package gamePart

import (
	"os"
	"strconv"
	"strings"
)

func SaveCurrentLevel() {
	file, err := os.OpenFile((WorldName + "#1"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		SetNewLog("Failed to save current level")
	}
	defer file.Close()
	//X Y TYPE
	for i := 0; i < int(SCREEN_WIDTH); i += CUBE_SIZE {
		for j := 0; j < int(SCREEN_HEIGHT); j += CUBE_SIZE {
			file.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(j) + " " + strconv.Itoa(int(MAP[i/CUBE_SIZE][j/CUBE_SIZE].ID)) + "\n")
		}
	}

	SetNewLog("Saved current level to " + (WorldName + "#1"))
}

func OpenSaveOfLevel() {
	level, err := os.ReadFile((WorldName + "#1"))
	if err != nil {
		SetNewLog("Failed to open current level")
		return
	}

	levelLines := strings.Split(strings.TrimSpace(string(level)), "\n")
	use := 0

	for i := 0; i < int(SCREEN_WIDTH); i += CUBE_SIZE {
		for j := 0; j < int(SCREEN_HEIGHT); j += CUBE_SIZE {
			if use < len(levelLines) {
				line := levelLines[use]
				parts := strings.Split(line, " ")
				if len(parts) >= 3 {
					// Преобразуем строки в целые числа
					x, _ := strconv.Atoi(parts[0])
					y, _ := strconv.Atoi(parts[1])
					t, _ := strconv.Atoi(parts[2])
					MAP[x/CUBE_SIZE][y/CUBE_SIZE].ID = rune(t)
				}
			}
			use++
		}
	}

	SetNewLog("Successfully opened level")
}
