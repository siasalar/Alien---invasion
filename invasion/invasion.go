package invasion

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// City represents a city in the alien invasion map
type City struct {
	Name                     string
	North, South, East, West *City
}

// CityMap represents the alien invasion map
type CityMap map[string]*City

// Alien represents an alien in the alien invasion simulation
type Alien struct {
	CurrentCity *City
	Moves       int
}

func Run(cityMap CityMap, numAliens int) {

	// Create the aliens and randomly place them on the map
	aliens := make([]*Alien, numAliens)
	for i := 0; i < numAliens; i++ {
		cityNames := make([]string, 0, len(cityMap))
		for cityName := range cityMap {
			cityNames = append(cityNames, cityName)
		}
		startCity := cityMap[cityNames[rand.Intn(len(cityNames))]]
		aliens[i] = &Alien{CurrentCity: startCity, Moves: 0}
	}

	// Run the simulation until all aliens are destroyed or each alien has moved at least 10000 times
	for {
		// Move each alien randomly
		for _, alien := range aliens {
			if alien.Moves >= 10000 {
				log.Println("Alien reached maximum moves, ending simulation")
				return
			}
			alien.Moves++

			// Get a list of possible directions to move
			possibleDirections := make([]*City, 0)
			if alien.CurrentCity.North != nil {
				possibleDirections = append(possibleDirections, alien.CurrentCity.North)
			}
			if alien.CurrentCity.South != nil {
				possibleDirections = append(possibleDirections, alien.CurrentCity.South)
			}
			if alien.CurrentCity.East != nil {
				possibleDirections = append(possibleDirections, alien.CurrentCity.East)
			}
			if alien.CurrentCity.West != nil {
				possibleDirections = append(possibleDirections, alien.CurrentCity.West)
			}

			// Move the alien to a random direction
			if len(possibleDirections) < 1 {
				log.Println("Alien is trapped and cannot move")
				continue
			}

			nextCity := possibleDirections[rand.Intn(len(possibleDirections))]
			alien.CurrentCity = nextCity
		}

		// Check for alien fights
		for i := 0; i < numAliens; i++ {
			for j := i + 1; j < numAliens; j++ {
				if aliens[i].CurrentCity == aliens[j].CurrentCity {
					fmt.Printf("Alien %d and Alien %d fought at %s, destroying the city\n", i, j, aliens[i].CurrentCity.Name)
					delete(cityMap, aliens[i].CurrentCity.Name)
					aliens[i].CurrentCity = nil
					aliens[j].CurrentCity = nil
					numAliens -= 2
					if numAliens == 0 {
						log.Println("All aliens have been destroyed, ending simulation")
						return
					}
				}
			}
		}
	}
}

//// createCityMap parse the map file and create the city map
//func createCityMap(scanner *bufio.Scanner) CityMap {
//	m := make(CityMap)
//	for scanner.Scan() {
//		line := scanner.Text()
//		parts := strings.Split(line, " ")
//		cityName := parts[0]
//
//		// Create the city if it doesn't exist yet
//		city, ok := m[cityName]
//		if !ok {
//			city = &City{Name: cityName}
//			m[cityName] = city
//		}
//
//		// Parse the city's connections
//		for _, connection := range parts[1:] {
//			dirAndCity := strings.Split(connection, "=")
//			direction := dirAndCity[0]
//			connectedCityName := dirAndCity[1]
//
//			// Create the connected city if it doesn't exist yet
//			connectedCity, ok := m[connectedCityName]
//			if !ok {
//				connectedCity = &City{Name: connectedCityName}
//				m[connectedCityName] = connectedCity
//			}
//
//			// Add the connection to the current city
//			switch direction {
//			case "north":
//				city.North = connectedCity
//			case "south":
//				city.South = connectedCity
//			case "east":
//				city.East = connectedCity
//			case "west":
//				city.West = connectedCity
//			}
//		}
//	}
//
//	return m
//}

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
		parts := strings.Split(line, " ")
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
