package middleofthelinkedlist

import "leetcode/utils/linkedlist"

type ListNode = linkedlist.ListNode

func middleNode(head *ListNode) *ListNode {
	// Floyd Cycle Detection 
	slowCursor := head
	fastCursor := head

	for fastCursor != nil && fastCursor.Next != nil {
		fastCursor = fastCursor.Next.Next
		slowCursor = slowCursor.Next
	}

	return slowCursor
}
