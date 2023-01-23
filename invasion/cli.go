package invasion

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetNumberOfAliens get the number of aliens from the command line
func GetNumberOfAliens() (int, error) {
	if len(os.Args) < 2 {
		log.Println("Please provide the number of aliens as a command line argument")
		return 0, fmt.Errorf("please provide the number of aliens as a command line argument")
	}

	numAliens, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Printf("Invalid number of aliens: %s", os.Args[1])
		return 0, fmt.Errorf("invalid number of aliens: %s", os.Args[1])
	}

	return numAliens, err
}
