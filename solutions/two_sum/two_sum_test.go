package twosum

import (
	"reflect"
	"testing"
)

type functionInput struct {
	nums   []int
	target int
}

type testInput struct {
	input    functionInput
	expected []int
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input: functionInput{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			input: functionInput{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			input: functionInput{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		result := twoSum(test.input.nums, test.input.target)

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
