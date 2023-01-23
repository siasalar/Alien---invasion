package main

import (
	"github.com/siasalar/alien-invasion/invasion"
)

const (
	mapFilePath = "map.txt"
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

	invasion.Run(cityMap, numAliens)
}
