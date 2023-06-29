package isomorphicstrings

import "testing"

type functionInput struct {
	s string
	t string
}

type testInput struct {
	input    functionInput
	expected bool
}

func TestPivotIndex(t *testing.T) {

	tests := []testInput{
		{
			input: functionInput{
				s: "egg",
				t: "add",
			},
			expected: true,
		},{
			input: functionInput{
				s: "foo",
				t: "bar",
			},
			expected: false,
		},{
			input: functionInput{
				s: "paper",
				t: "title",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		result := isIsomorphic(test.input.s, test.input.t)
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
