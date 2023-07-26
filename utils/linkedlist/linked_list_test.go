package linkedlist

import (
	"testing"
)

func TestArrayToLinkedList(t *testing.T) {

	arrayToConvert := []int{1, 2, 3, 4, 5}

	currentNode := ArrayToLinkedList(arrayToConvert)

	for _, val := range arrayToConvert {

		currentNodeValue := currentNode.Val

		if currentNodeValue != val {
			t.Errorf("\n"+
				"\nTest Failed:"+
				"\n array to convert:    %+v"+
				"\n broke on value:   %+v"+
				"\n received from linked list: %+v \n\n",
				arrayToConvert,
				val,
				currentNodeValue,
			)
		}

		nextNode := currentNode.Next

		if nextNode != nil {
			currentNode = nextNode
		}

	}

	if currentNode.Next != nil {
		t.Errorf("\n"+
			"\nTest Failed:"+
			"\n array to convert:    %+v"+
			"\n tail.Next should be nil:   %+v"+
			"\n received from linked list: %+v \n\n",
			arrayToConvert,
			nil,
			currentNode.Next,
		)
	}
}

func TestLinkedListsMatch(t *testing.T) {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{1, 2, 3, 4, 5, 6}
	array3 := []int{1, 2, 3, 4, 4}

	if !LinkedListsMatch(ArrayToLinkedList(array1), ArrayToLinkedList(array1)) {
		t.Errorf("\n"+
			"\nTest Failed:"+
			"\n list 1:    %+v"+
			"\n list 2:    %+v"+
			"\n the lists match, but LinkedListMatch did not return true\n\n",
			array1,
			array1,
		)
	}

	if LinkedListsMatch(ArrayToLinkedList(array1), ArrayToLinkedList(array2)) {
		t.Errorf("\n"+
			"\nTest Failed:"+
			"\n list 1:    %+v"+
			"\n list 2:    %+v"+
			"\n the lists don't match, but LinkedListMatch returned true\n\n",
			array1,
			array2,
		)
	}

	if LinkedListsMatch(ArrayToLinkedList(array1), ArrayToLinkedList(array3)) {
		t.Errorf("\n"+
			"\nTest Failed:"+
			"\n list 1:    %+v"+
			"\n list 2:    %+v"+
			"\n the lists don't match, but LinkedListMatch returned true\n\n",
			array1,
			array3,
		)
	}

}
