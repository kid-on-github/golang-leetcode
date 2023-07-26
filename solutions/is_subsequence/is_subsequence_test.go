package subsequence

import (
	"testing"
)

type functionInput struct {
	s string
	t string
}

type testInput struct {
	input    functionInput
	expected bool
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input: functionInput{
				s: "abc",
				t: "ahbgdc",
			},
			expected: true,
		},
		{
			input: functionInput{
				s: "axc",
				t: "ahbgdc",
			},
			expected: false,
		},
		{
			input: functionInput{
				s: "aaaaaa",
				t: "bbaaaa",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		result := isSubsequence(test.input.s, test.input.t)

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
