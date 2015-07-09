package main

//go run gameMode.go

import (
	"fmt"
	"github.com/geeks-dev/go-choices"
)

func main() {

	gamemode := []string{
		"HUMAN",
		"DEVIL HUNTER",
		"SON OF SPARDA",
		"HEAVEN OR HELL",
		"LEGENDARY DARK KNIGHT",
		"DANTE MUST DIE",
		"HELL AND HELL",
		"BLOODY PALACE",
	}
	mode, answered := choices.Ask(
		gamemode,
		"Devil May Cry",
	)

	if answered {
		fmt.Println(gamemode[mode])
	}
}
