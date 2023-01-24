package invasion

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestGetListOfPossibleDirections(t *testing.T) {
	testCases := []struct {
		name     string
		alien    Alien
		worldMap WorldMap
		want     []string
	}{
		{
			name: "Test All Directions From CityA",
			alien: Alien{
				CurrentCity: City{
					Name: "cityA",
				},
				IsAlive: true,
			},
			worldMap: WorldMap{
				"cityA": City{
					Connections: map[string]string{
						"north": "cityB",
						"south": "cityC",
					},
				},
				"cityB": City{
					Connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": City{
					Connections: map[string]string{
						"north": "cityA",
					},
				},
			},
			want: []string{"north", "south"},
		},
		{
			name: "Test Alien In CityB No Directions Exist",
			alien: Alien{
				CurrentCity: City{
					Name: "cityB",
				},
				IsAlive: true,
			},
			worldMap: WorldMap{
				"cityA": City{
					Connections: map[string]string{
						"north": "cityB",
					},
				},
				"cityB": City{
					Connections: map[string]string{
						"south": "cityA",
					},
				},
			},
			want: []string{"south"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getListOfPossibleDirections(tc.alien, tc.worldMap)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getListOfPossibleDirections(%v, %v) = %v, want %v", tc.alien, tc.worldMap, got, tc.want)
			}
		})
	}
}

func TestUpdateWorldMap(t *testing.T) {
	testCases := []struct {
		name          string
		worldMap      WorldMap
		destroyedCity City
		want          WorldMap
	}{
		{
			name: "destroying cityA which is between cityB and cityC",
			worldMap: WorldMap{
				"cityA": City{
					Name: "cityA",
					Connections: map[string]string{
						"north": "cityB",
						"west":  "cityC",
					},
				},
				"cityB": City{
					Name: "cityB",
					Connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": City{
					Name: "cityC",
					Connections: map[string]string{
						"east": "cityA",
					},
				},
			},
			destroyedCity: City{
				Name: "cityA",
				Connections: map[string]string{
					"north": "cityB",
					"west":  "cityC",
				},
			},
			want: WorldMap{
				"cityB": City{
					Name:        "cityB",
					Connections: map[string]string{},
				},
				"cityC": City{
					Name:        "cityC",
					Connections: map[string]string{},
				},
			},
		},
		{
			name: "destroying CityB which is in the corner of the world",
			worldMap: WorldMap{
				"cityA": City{
					Name: "cityA",
					Connections: map[string]string{
						"north": "cityB",
						"west":  "cityC",
					},
				},
				"cityB": City{
					Name: "cityB",
					Connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": City{
					Name: "cityC",
					Connections: map[string]string{
						"east": "cityA",
					},
				},
			},
			destroyedCity: City{
				Name: "cityB",
				Connections: map[string]string{
					"south": "cityA",
				},
			},
			want: WorldMap{
				"cityA": City{
					Name: "cityA",
					Connections: map[string]string{
						"west": "cityC",
					},
				},
				"cityC": City{
					Name: "cityC",
					Connections: map[string]string{
						"east": "cityA",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := updateWorldMap(tc.worldMap, tc.destroyedCity)

			require.Equal(t, tc.want, got)
		})
	}
}
