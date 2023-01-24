package invasion

import (
	"fmt"
)

// PrintWorldMap print out whatever is left of the world in the same format as the input file.
func PrintWorldMap(worldMap WorldMap) {
	fmt.Println("Remained world map:")

	for cityName, connectedCity := range worldMap {
		line := cityName

		north, ok := connectedCity.connections["north"]
		if ok {
			line += " north=" + north
		}

		south, ok := connectedCity.connections["south"]
		if ok {
			line += " south=" + south
		}

		east, ok := connectedCity.connections["east"]
		if ok {
			line += " east=" + east
		}

		west, ok := connectedCity.connections["west"]
		if ok {
			line += " west=" + west
		}

		fmt.Println(line)
	}
}
