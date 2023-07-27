package reverselinkedlist

import (
	LinkedList "leetcode/utils/linkedlist"
	"testing"
)

type testInput struct {
	input    []int
	expected []int
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		inputList := LinkedList.ArrayToLinkedList(test.input)
		expected := LinkedList.ArrayToLinkedList(test.expected)
		reversed := reverseList(inputList)

		if !LinkedList.LinkedListsMatch(reversed, expected) {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n input list:    %+v"+
				"\n reversed result: %+v"+
				"\n expected result: %+v \n\n",
				test.input,
				reversed,
				expected,
			)
		}
	}

}
