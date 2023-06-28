package findpivotindex

import "testing"

type testInput struct {
	input    []int
	expected int
}

func TestPivotIndex(t *testing.T) {

	tests := []testInput{
		{
			input:    []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3},
			expected: -1,
		},
		{
			input:    []int{2, 1, -1},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := pivotIndex(test.input)
		if result != test.expected {
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
