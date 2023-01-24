package invasion

import (
	"log"
	"math/rand"
	"time"
)

// City represents a city in the alien invasion map
type City struct {
	Name        string
	Connections map[string]string // key:direction value:city
}

// WorldMap represents the alien invasion map
type WorldMap map[string]City // key:cityName value:connectedCity

// Alien represents an alien in the alien invasion simulation
type Alien struct {
	CurrentCity City
	IsAlive     bool
}

func Run(worldMap WorldMap, numAliens, movementThresh int) {
	// Create the aliens and randomly place them on the map
	aliens := make([]Alien, numAliens)
	for i := 0; i < numAliens; i++ {
		var cityNames []string
		for cityName := range worldMap {
			cityNames = append(cityNames, cityName)
		}

		rand.Seed(time.Now().UnixNano())
		aliens[i] = Alien{
			CurrentCity: worldMap[cityNames[rand.Intn(len(cityNames))]],
			IsAlive:     true,
		}
	}

	// Run the simulation until all aliens are destroyed or each alien has moved at least 10000 times
	counter := movementThresh
	for {
		if counter == 0 {
			log.Printf("Alien reached maximum moves: %d move, ending simulation!", movementThresh)
			return
		}
		counter--

		// Move each alien randomly
		for i, alien := range aliens {
			if alien.IsAlive {
				possibleDirections := getListOfPossibleDirections(alien, worldMap)

				// Move the alien to a randInt direction
				if len(possibleDirections) < 1 {
					log.Printf("Alien %d is trapped and cannot move", i)
					return
				}

				rand.Seed(time.Now().UnixNano())
				nextCity := possibleDirections[rand.Intn(len(possibleDirections))]
				alien.CurrentCity.Name = nextCity
			}
		}

		// Check for alien fights
		for i := 0; i < numAliens; i++ {
			for j := i + 1; j < numAliens; j++ {
				if aliens[i].CurrentCity.Name != aliens[j].CurrentCity.Name || !aliens[i].IsAlive || !aliens[j].IsAlive {
					continue
				}

				destroyedCity := aliens[i].CurrentCity
				log.Printf("Alien %d and Alien %d fought at %s, destroying the city\n", i, j, destroyedCity.Name)

				aliens[i].IsAlive = false
				aliens[j].IsAlive = false
				i++ // current alien[i] is dead

				worldMap = updateWorldMap(worldMap, destroyedCity)

				numAliens -= 2
				if numAliens == 0 {
					log.Println("All aliens have been destroyed, ending simulation")
					return
				}
			}
		}
	}
}

func updateWorldMap(worldMap WorldMap, destroyedCity City) WorldMap {
	// find the list of the cities which their connections needs to be updated
	var cities []string
	for _, city := range destroyedCity.Connections {
		cities = append(cities, city)
	}

	// delete any roads that lead into or out of destroyed city
	for _, city := range cities {
		for direction, cityName := range worldMap[city].Connections {
			if cityName == destroyedCity.Name {
				delete(worldMap[city].Connections, direction)
			}
		}
	}

	// delete the destroyed city from world map
	delete(worldMap, destroyedCity.Name)

	return worldMap
}

func getListOfPossibleDirections(alien Alien, cityMap WorldMap) (possibleDirections []string) {
	city := cityMap[alien.CurrentCity.Name]

	for direction := range city.Connections {
		possibleDirections = append(possibleDirections, direction)
	}

	return possibleDirections
}
