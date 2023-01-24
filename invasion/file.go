package invasion

import (
	"bufio"
	"os"
	"strings"
)

// ReadCityMapFile parse the map file and create the city map
func ReadCityMapFile(path string) (CityMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	cityMap := make(CityMap)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 1 {
			continue
		}

		cityName := parts[0]

		// Create the city if it doesn't exist yet
		city, ok := cityMap[cityName]
		if !ok {
			city = &City{Name: cityName}
			cityMap[cityName] = city
		}

		// Parse the city's connections
		for _, connection := range parts[1:] {
			dirAndCity := strings.Split(connection, "=")
			direction := dirAndCity[0]
			connectedCityName := dirAndCity[1]

			// Create the connected city if it doesn't exist yet
			connectedCity, ok := cityMap[connectedCityName]
			if !ok {
				connectedCity = &City{Name: connectedCityName}
				cityMap[connectedCityName] = connectedCity
			}

			// Add the connection to the current city
			switch direction {
			case "north":
				city.North = connectedCity
			case "south":
				city.South = connectedCity
			case "east":
				city.East = connectedCity
			case "west":
				city.West = connectedCity
			}
		}
	}

	return cityMap, nil
}
