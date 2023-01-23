package invasion

import (
	"log"
	"math/rand"
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

func Run(cityMap CityMap, numAliens, movementThresh int) {
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
			if alien.Moves >= movementThresh {
				log.Println("Alien reached maximum moves, ending simulation")
				return
			}
			alien.Moves++

			possibleDirections := getListOfPossibleDirections(alien)

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
					log.Printf("Alien %d and Alien %d fought at %s, destroying the city\n", i, j, aliens[i].CurrentCity.Name)
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

func getListOfPossibleDirections(alien *Alien) []*City {
	possibleDirections := make([]*City, 0)

	directions := map[string]*City{
		"north": alien.CurrentCity.North,
		"south": alien.CurrentCity.South,
		"east":  alien.CurrentCity.East,
		"west":  alien.CurrentCity.West,
	}

	for _, direction := range directions {
		switch {
		case direction != nil:
			possibleDirections = append(possibleDirections, direction)
		}
	}
	return possibleDirections
}
