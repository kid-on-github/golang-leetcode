package validpalindrome

import (
	"testing"
)

type testInput struct {
	input    string
	expected bool
}

func TestRunningSum(t *testing.T) {

	tests := []testInput{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    " ",
			expected: true,
		},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)

		if !result == test.expected {
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
