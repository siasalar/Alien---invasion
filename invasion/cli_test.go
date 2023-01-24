package invasion

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNumberOfAliens(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectedNum int
		expectedErr error
	}{
		{
			name:        "valid input",
			args:        []string{"invasion", "5"},
			expectedNum: 5,
			expectedErr: nil,
		},
		{
			name:        "invalid input",
			args:        []string{"invasion", "not_a_number"},
			expectedNum: 0,
			expectedErr: fmt.Errorf("invalid number of aliens: not_a_number"),
		},
		{
			name:        "no input",
			args:        []string{"invasion"},
			expectedNum: 0,
			expectedErr: fmt.Errorf("please provide the number of aliens as a command line argument"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args
			numAliens, err := GetNumberOfAliens()

			require.Equal(t, tc.expectedNum, numAliens)

			if err != nil {
				require.Contains(t, err.Error(), tc.expectedErr.Error())
			}
		})
	}
}
