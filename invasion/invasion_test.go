package invasion

import (
	"reflect"
	"testing"
)

func TestGetListOfPossibleDirections(t *testing.T) {
	type testCase struct {
		name        string
		alien       *Alien
		expectedOut []*City
	}

	northCity := &City{Name: "North City"}
	southCity := &City{Name: "South City"}
	eastCity := &City{Name: "East City"}
	westCity := &City{Name: "West City"}

	testCases := []testCase{
		{
			name:        "all directions are not nil",
			alien:       &Alien{CurrentCity: &City{North: northCity, South: southCity, East: eastCity, West: westCity}},
			expectedOut: []*City{northCity, southCity, eastCity, westCity},
		},
		{
			name:        "north and south directions are nil",
			alien:       &Alien{CurrentCity: &City{East: eastCity, West: westCity}},
			expectedOut: []*City{eastCity, westCity},
		},
		{
			name:        "east and west directions are nil",
			alien:       &Alien{CurrentCity: &City{North: northCity, South: southCity}},
			expectedOut: []*City{northCity, southCity},
		},
		{
			name:        "all directions are nil",
			alien:       &Alien{CurrentCity: &City{}},
			expectedOut: []*City{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getListOfPossibleDirections(tc.alien)
			if !reflect.DeepEqual(result, tc.expectedOut) {
				t.Errorf("Expected %v, got %v", tc.expectedOut, result)
			}
		})
	}
}
