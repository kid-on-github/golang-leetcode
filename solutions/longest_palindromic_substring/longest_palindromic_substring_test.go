package longestpalindromicsubstring

import (
	"testing"
)

type testInput struct {
	input    string
	expected string
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input:    "1234aa",
			expected: "aa",
		},
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "asdf",
			expected: "a",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input: "aaaa",
			expected: "aaaa",
		},
	}

	for _, test := range tests {
		result := longestPalindrome(test.input)

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
