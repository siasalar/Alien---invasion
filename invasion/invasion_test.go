package invasion

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetListOfPossibleDirections(t *testing.T) {
	testCases := []struct {
		name     string
		alien    alien
		worldMap WorldMap
		want     []string
	}{
		{
			name: "Test All Directions From CityA",
			alien: alien{
				currentCity: city{
					name: "cityA",
				},
				isAlive: true,
			},
			worldMap: WorldMap{
				"cityA": city{
					connections: map[string]string{
						"north": "cityB",
						"south": "cityC",
					},
				},
				"cityB": city{
					connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": city{
					connections: map[string]string{
						"north": "cityA",
					},
				},
			},
			want: []string{"north", "south"},
		},
		{
			name: "Test alien In CityB No Directions Exist",
			alien: alien{
				currentCity: city{
					name: "cityB",
				},
				isAlive: true,
			},
			worldMap: WorldMap{
				"cityA": city{
					connections: map[string]string{
						"north": "cityB",
					},
				},
				"cityB": city{
					connections: map[string]string{
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
		destroyedCity city
		want          WorldMap
	}{
		{
			name: "destroying cityA which is between cityB and cityC",
			worldMap: WorldMap{
				"cityA": city{
					name: "cityA",
					connections: map[string]string{
						"north": "cityB",
						"west":  "cityC",
					},
				},
				"cityB": city{
					name: "cityB",
					connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": city{
					name: "cityC",
					connections: map[string]string{
						"east": "cityA",
					},
				},
			},
			destroyedCity: city{
				name: "cityA",
				connections: map[string]string{
					"north": "cityB",
					"west":  "cityC",
				},
			},
			want: WorldMap{
				"cityB": city{
					name:        "cityB",
					connections: map[string]string{},
				},
				"cityC": city{
					name:        "cityC",
					connections: map[string]string{},
				},
			},
		},
		{
			name: "destroying CityB which is in the corner of the world",
			worldMap: WorldMap{
				"cityA": city{
					name: "cityA",
					connections: map[string]string{
						"north": "cityB",
						"west":  "cityC",
					},
				},
				"cityB": city{
					name: "cityB",
					connections: map[string]string{
						"south": "cityA",
					},
				},
				"cityC": city{
					name: "cityC",
					connections: map[string]string{
						"east": "cityA",
					},
				},
			},
			destroyedCity: city{
				name: "cityB",
				connections: map[string]string{
					"south": "cityA",
				},
			},
			want: WorldMap{
				"cityA": city{
					name: "cityA",
					connections: map[string]string{
						"west": "cityC",
					},
				},
				"cityC": city{
					name: "cityC",
					connections: map[string]string{
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
