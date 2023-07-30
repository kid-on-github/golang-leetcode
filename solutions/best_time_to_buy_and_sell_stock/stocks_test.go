package besttimetobuyandsellstock

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			input:    []int{2, 1, 2, 1, 0, 1, 2},
			expected: 2,
		},
	}

	for _, test := range tests {
		returned := maxProfit(test.input)
		if returned != test.expected {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n input:    %+v"+
				"\n output:   %+v"+
				"\n expected: %+v \n\n",
				test.input,
				returned,
				test.expected,
			)
		}
	}
}
