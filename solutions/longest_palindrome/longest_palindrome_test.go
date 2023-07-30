package longestpalindrome

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abccccdd",
			expected: 7,
		},
		{
			input:    "a",
			expected: 1,
		},
	}

	for _, test := range tests {
		returned := longestPalindrome(test.input)
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
