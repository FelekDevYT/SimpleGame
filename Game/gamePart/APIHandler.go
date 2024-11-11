package gamePart

type VEC2 struct {
	x rune
	y rune
}

type BlockEvent struct {
	cords     VEC2
	blockType rune
	isActive  bool
}

type WorldEvent struct {
	isWorldAlreadyOpened bool
}

type GameAPIEvents struct {
	isBlockSet BlockEvent
	WorldEvent WorldEvent
}

func clearAndSetDisable(api *GameAPIEvents) {
	api.isBlockSet.isActive = false
	api.isBlockSet.cords.x = 0
	api.isBlockSet.cords.y = 0
	api.isBlockSet.blockType = 0
	//
	api.WorldEvent.isWorldAlreadyOpened = false
}

func setEnable(api *GameAPIEvents, typeOfEvent rune) {
	switch typeOfEvent {
	case 0:
		api.isBlockSet.isActive = true
		break
	}
}
