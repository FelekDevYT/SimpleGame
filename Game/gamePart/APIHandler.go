package gamePart

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

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

//THE UP LEVEL FUNCS

type Mod struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
}

func GetModInfo(modpath string) Mod {
	// Открываем JSON файл
	file, _ := os.Open(modpath + "\\mod.json") //https://open.spotify.com/playlist/2X1mh4DY0PSJEJ9DY0j6RH?si=dYMelAP-Szq8yh2s1oQrqg
	defer file.Close()

	// Читаем содержимое файла
	byteValue, _ := ioutil.ReadAll(file)

	// Декодируем JSON в структуру
	var mod Mod
	if err := json.Unmarshal(byteValue, &mod); false {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}
	return mod
}
