package runningsum

import (
	"reflect"
	"testing"
)

type testInput struct {
	input    []int
	expected []int
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input: []int{1, 2, 3, 4},

			expected: []int{1, 3, 6, 10},
		},
		{
			input: []int{1, 1, 1, 1, 1},

			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{3, 1, 2, 10, 1},

			expected: []int{3, 4, 6, 16, 17},
		},
		{
			input: []int{1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		result := runningSum(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n input:    %+v"+
				"\n output:   %+v"+
				"\n expected: %+v \n\n",
				test.input,
				result,
				test.expected,
			)
		}
	}

}
