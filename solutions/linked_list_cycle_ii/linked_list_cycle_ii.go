package linkedlistcycleii

import (
	"leetcode/utils/linkedlist"
)

type ListNode = linkedlist.ListNode

// find start of cycle
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// Floyd Cycle Detection
	slowCursor := head
	fastCursor := head.Next

	for fastCursor != slowCursor {
		if fastCursor == nil || fastCursor.Next == nil {
			return nil
		}

		fastCursor = fastCursor.Next.Next
		slowCursor = slowCursor.Next
	}

	stopSign := slowCursor
	slowCursor = head

	for slowCursor != stopSign {
		fastCursor = stopSign

		for fastCursor.Next != stopSign {
			fastCursor = fastCursor.Next
			if fastCursor == slowCursor {
				return fastCursor
			}
		}

		slowCursor = slowCursor.Next
	}
	return fastCursor
}
