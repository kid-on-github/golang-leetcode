package reverselinkedlist

import "leetcode/utils/linkedlist"

type ListNode = linkedlist.ListNode

func reverseList(head *ListNode) *ListNode {

	if (head == nil){
		return head
	}

	prevNode := head
	currentNode := head.Next
	prevNode.Next = nil

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode
}
