package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToLinkedList(arr []int) (head *ListNode) {
	if len(arr) == 0 {
		return
	}

	head = new(ListNode)
	currentNode := head

	for i, val := range arr {
		currentNode.Val = val

		if i < len(arr)-1 {
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next
		}
	}

	return
}

func LinkedListsMatch(list1 *ListNode, list2 *ListNode) bool {

	for list1 != nil {
		if list2 == nil {
			return false
		}

		if list1.Val != list2.Val {
			return false
		}
		
		list1 = list1.Next
		list2 = list2.Next
	}

	return list2 == nil
}
