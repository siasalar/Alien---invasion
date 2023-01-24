package invasion

import (
	"crypto/rand"
	"log"
	"math/big"
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

		randInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(cityNames))))
		if err != nil {
			log.Printf("failed to generate random Int number: %v", err)
			return
		}
		startCity := cityMap[cityNames[randInt.Int64()]]
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

			// Move the alien to a randInt direction
			if len(possibleDirections) < 1 {
				log.Println("Alien is trapped and cannot move")
				continue
			}

			randInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(possibleDirections))))
			if err != nil {
				log.Printf("failed to generate random Int number: %v", err)
				return
			}
			nextCity := possibleDirections[randInt.Int64()]

			alien.CurrentCity = nextCity
		}

		// Check for alien fights
		for i := 0; i < numAliens; i++ {
			for j := i + 1; j < numAliens; j++ {
				if aliens[i].CurrentCity != aliens[j].CurrentCity {
					continue
				}

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

func getListOfPossibleDirections(alien *Alien) []*City {
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

	return possibleDirections
}
