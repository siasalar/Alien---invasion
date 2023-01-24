package invasion

import (
	"log"
	"math/rand"
	"time"
)

// city represents a city in the alien invasion map
type city struct {
	name        string
	connections map[string]string // key:direction value:city
}

// WorldMap represents the alien invasion map
type WorldMap map[string]city // key:cityName value:connectedCity

// alien represents an alien in the alien invasion simulation
type alien struct {
	currentCity city
	isAlive     bool
}

func Run(worldMap WorldMap, numAliens, movementThresh int) {
	aliens := createAliens(worldMap, numAliens)

	// Run the simulation until all aliens are destroyed or each alien has moved at least 10000 times
	counter := movementThresh

	for {
		if counter == 0 {
			log.Printf("alien reached maximum moves: %d move, ending simulation!", movementThresh)
			return
		}
		counter--

		if isAlive := moveAliens(worldMap, aliens); !isAlive {
			return
		}

		// Check for alien fights
		for i := 0; i < numAliens; i++ {
			for j := i + 1; j < numAliens; j++ {
				if aliens[i].currentCity.name != aliens[j].currentCity.name || !aliens[i].isAlive || !aliens[j].isAlive {
					continue
				}

				destroyedCity := aliens[i].currentCity
				log.Printf("alien %d and alien %d fought at %s, destroying the city\n", i, j, destroyedCity.name)

				aliens[i].isAlive = false
				aliens[j].isAlive = false
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

// moveAliens Move each alien randomly
func moveAliens(worldMap WorldMap, aliens []alien) bool {
	for i, alien := range aliens {
		if !alien.isAlive {
			continue
		}

		possibleDirections := getListOfPossibleDirections(alien, worldMap)

		// Move the alien to a randInt direction
		if len(possibleDirections) < 1 {
			log.Printf("alien %d is trapped and cannot move", i)
			return false
		}

		rand.Seed(time.Now().UnixNano())
		nextCity := possibleDirections[rand.Intn(len(possibleDirections))]
		alien.currentCity.name = nextCity
	}

	return true
}

// createAliens Create the aliens and randomly place them on the map
func createAliens(worldMap WorldMap, numAliens int) []alien {
	aliens := make([]alien, numAliens)

	for i := 0; i < numAliens; i++ {
		var cityNames []string
		for cityName := range worldMap {
			cityNames = append(cityNames, cityName)
		}

		rand.Seed(time.Now().UnixNano())

		aliens[i] = alien{
			currentCity: worldMap[cityNames[rand.Intn(len(cityNames))]],
			isAlive:     true,
		}
	}

	return aliens
}

// updateWorldMap update the world map based on destroyed city
func updateWorldMap(worldMap WorldMap, destroyedCity city) WorldMap {
	// find the list of the cities which their connections needs to be updated
	var cities []string
	for _, city := range destroyedCity.connections {
		cities = append(cities, city)
	}

	// delete any roads that lead into or out of destroyed city
	for _, city := range cities {
		for direction, cityName := range worldMap[city].connections {
			if cityName == destroyedCity.name {
				delete(worldMap[city].connections, direction)
			}
		}
	}

	// delete the destroyed city from world map
	delete(worldMap, destroyedCity.name)

	return worldMap
}

// getListOfPossibleDirections gets the list of possible directions of a city
func getListOfPossibleDirections(alien alien, cityMap WorldMap) (possibleDirections []string) {
	c := cityMap[alien.currentCity.name]

	for direction := range c.connections {
		possibleDirections = append(possibleDirections, direction)
	}

	return possibleDirections
}
