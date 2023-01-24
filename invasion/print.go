package invasion

import (
	"fmt"
)

// PrintWorldMap print out whatever is left of the world in the same format as the input file.
func PrintWorldMap(worldMap WorldMap) {
	fmt.Println("Remained world map:")

	for cityName, connectedCity := range worldMap {
		line := cityName

		north, ok := connectedCity.Connections["north"]
		if ok {
			line += " north=" + north
		}

		south, ok := connectedCity.Connections["south"]
		if ok {
			line += " south=" + south
		}

		east, ok := connectedCity.Connections["east"]
		if ok {
			line += " east=" + east
		}

		west, ok := connectedCity.Connections["west"]
		if ok {
			line += " west=" + west
		}

		fmt.Println(line)
	}
}
