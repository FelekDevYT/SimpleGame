package gamePart

func plains() {
	var points [1000]rune

	for i := 0; rune(i) < (SCREEN_WIDTH / CUBE_SIZE); i++ {
		for points[i] == 0 {
			points[i] = getIndex(rune(random(10, 14))) // - new
		}
	}

	for X := 0; rune(X) < (SCREEN_WIDTH / CUBE_SIZE); X++ {
		for Y := 0; rune(Y) < points[X]; Y++ {
			MAP[X][Y] = 1
		}
	}

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

}
