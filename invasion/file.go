package invasion

import (
	"bufio"
	"os"
	"strings"
)

// ReadWorldMapFile parse the map file and create the city map
func ReadWorldMapFile(path string) (WorldMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	worldMap := make(WorldMap)

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

		// Create the city in worldMap if it doesn't exist yet
		city, ok := worldMap[cityName]
		if !ok {
			city = City{Name: cityName}
			worldMap[cityName] = city
		}

		// Parse the city's connections
		city.Connections = make(map[string]string)
		for _, connection := range parts[1:] {
			dirAndCity := strings.Split(connection, "=")
			direction := dirAndCity[0]
			connectedCityName := dirAndCity[1]

			// Add the connection to the current city
			city.Connections[direction] = connectedCityName

			// Add the city to city map
			worldMap[cityName] = city
		}
	}

	return worldMap, nil
}
