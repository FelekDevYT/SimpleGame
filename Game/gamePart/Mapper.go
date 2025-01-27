package gamePart

func CreateMap() {
	plains()
}

func DrawMap() {
	for i := 0; i < int(SCREEN_WIDTH); i += CUBE_SIZE {
		for j := 0; j < int(SCREEN_HEIGHT); j += CUBE_SIZE {
			switch MAP[i/CUBE_SIZE][j/CUBE_SIZE].ID {
			case 0:
				renderCube(rune(i), rune(j), colors["GRASS"])
			case 1:
				renderCube(rune(i), rune(j), colors["SKY"])
			case 2:
				renderCube(rune(i), rune(j), colors["STONE"])
			case 3:
				renderCube(rune(i), rune(j), colors["SNOW"])
			case 4:
				renderCube(rune(i), rune(j), colors["IRON"])
			case 5:
				renderCube(rune(i), rune(j), colors["GOLD"])
			case 6:
				renderCube(rune(i), rune(j), colors["DIAMOND"])
			}
		}
	}
}
