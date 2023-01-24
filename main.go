package main

import (
	"github.com/siasalar/alien-invasion/invasion"
)

const (
	mapFilePath         = "map.txt"
	alienMovementThresh = 10000
)

func main() {
	numAliens, err := invasion.GetNumberOfAliens()
	if err != nil {
		panic(err)
	}

	worldMap, err := invasion.ReadWorldMapFile(mapFilePath)
	if err != nil {
		panic(err)
	}

	invasion.Run(worldMap, numAliens, alienMovementThresh)

	invasion.PrintWorldMap(worldMap)
}


