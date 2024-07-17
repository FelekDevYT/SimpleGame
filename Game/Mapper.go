package main

func createMap() {
	var points [1000]rune

	for i := 0; rune(i) < (SCREEN_WIDTH / CUBE_SIZE); i++ {
		for points[i] == 0 {
			lvl := int(SCREEN_HEIGHT / CUBE_SIZE)
			//points[i] = getIndex(rune(rand.Intn(int(SCREEN_HEIGHT / CUBE_SIZE)))) - old
			if isLeftBorder(rune(i)) {
				points[i] = getIndex(rune(random(lvl/2, lvl+2))) // - new
			} else {
				points[i] = getIndex(rune(random(lvl/2, lvl+(random(lvl/2, lvl)/2)))) // - new
			}
		}
	}

	setNewLog("[MAPPING-1]Generate points of heights")

	for X := 0; rune(X) < (SCREEN_WIDTH / CUBE_SIZE); X++ {
		for Y := 0; rune(Y) < points[X]; Y++ {
			MAP[X][Y] = 1
		}
	}

	setNewLog("[MAPPING-2]Setting points to MAP")

	for X := 0; rune(X) < (SCREEN_WIDTH / CUBE_SIZE); X++ {
		for Y := 0; rune(Y) < (SCREEN_HEIGHT / CUBE_SIZE); Y++ {
			if Y <= getDefaultIndex(rune(random(4, 7))) && MAP[X][Y] == 0 {
				MAP[X][Y] = 3
			}

			if Y >= getDefaultIndex(9) && MAP[X][Y] == 0 && isOre(iron) {
				MAP[X][Y] = 4
			}
			if Y >= getDefaultIndex(10) && MAP[X][Y] == 0 && isOre(gold) {
				MAP[X][Y] = 5
			}
			if Y >= getDefaultIndex(12) && MAP[X][Y] == 0 && isOre(diamond) {
				MAP[X][Y] = 6
			}

			if Y >= getDefaultIndex(rune(random(8, 11))) && MAP[X][Y] == 0 {
				MAP[X][Y] = 2
			}
		}
	}

	setNewLog("[MAPPING-3]Setting levels of height to MAP")
}

func drawMap() {
	for i := 0; i < int(SCREEN_WIDTH); i += CUBE_SIZE {
		for j := 0; j < int(SCREEN_HEIGHT); j += CUBE_SIZE {
			switch MAP[i/CUBE_SIZE][j/CUBE_SIZE] {
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
