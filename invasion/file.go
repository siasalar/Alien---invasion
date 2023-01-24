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

		// Create the city in WorldMap if it doesn't exist yet
		c, ok := worldMap[cityName]
		if !ok {
			c = city{name: cityName}
			worldMap[cityName] = c
		}

		// Parse the city's connections
		c.connections = make(map[string]string)

		for _, connection := range parts[1:] {
			dirAndCity := strings.Split(connection, "=")
			direction := dirAndCity[0]
			connectedCityName := dirAndCity[1]

			// Add the connection to the current city
			c.connections[direction] = connectedCityName

			// Add the city to city map
			worldMap[cityName] = c
		}
	}

	return worldMap, nil
}
