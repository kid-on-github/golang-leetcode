package mergetwosortedlists

import (
	LinkedList "leetcode/utils/linkedlist"
	"testing"
)

type functionInput struct {
	list1 []int
	list2 []int
}

type testInput struct {
	input    functionInput
	expected []int
}

func Test(t *testing.T) {

	tests := []testInput{
		{
			input: functionInput{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input: functionInput{
				list1: []int{},
				list2: []int{},
			},
			expected: []int{},
		},
		{
			input: functionInput{
				list1: []int{},
				list2: []int{0},
			},
			expected: []int{0},
		},
		{
			input: functionInput{
				list1: []int{1},
				list2: []int{},
			},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		list1 := LinkedList.ArrayToLinkedList(test.input.list1)
		list2 := LinkedList.ArrayToLinkedList(test.input.list2)
		expected := LinkedList.ArrayToLinkedList(test.expected)

		merged := mergeTwoLists(list1, list2)

		if !LinkedList.LinkedListsMatch(merged, expected) {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n input list1:    %+v"+
				"\n input list2:   %+v"+
				"\n merge result: %+v"+
				"\n expected result: %+v \n\n",
				test.input.list1,
				test.input.list2,
				merged,
				expected,
			)
		}
	}

}
