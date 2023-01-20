package main

import (
	"github.com/siasalar/Alien-invasion/alieninvasion"
)

const (
	mapFilePath = "map.txt"
)

func main() {
	alieninvasion.Run(mapFilePath)
}
