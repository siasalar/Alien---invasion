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

	cityMap, err := invasion.ReadCityMapFile(mapFilePath)
	if err != nil {
		panic(err)
	}

	invasion.Run(cityMap, numAliens, alienMovementThresh)
}
