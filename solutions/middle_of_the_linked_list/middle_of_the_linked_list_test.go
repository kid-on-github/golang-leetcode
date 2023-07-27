package middleofthelinkedlist

import (
	LinkedList "leetcode/utils/linkedlist"
	"testing"
)

func Test(t *testing.T) {

	tests := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4},
	}

	for _, test := range tests {
		expectedMiddle := LinkedList.ArrayToLinkedList(test)
		returnedMiddle := middleNode(expectedMiddle)

		for i := 0; i < len(test)/2; i++ {
			expectedMiddle = expectedMiddle.Next
		}

		if returnedMiddle != expectedMiddle {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n input list:    %+v"+
				"\n returned middle value: %+v"+
				"\n expected middle value: %+v \n\n",
				test,
				returnedMiddle.Val,
				expectedMiddle.Val,
			)
		}
	}

}
