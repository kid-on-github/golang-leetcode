package mergetwosortedlists

import "leetcode/utils/linkedlist"

type ListNode = linkedlist.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		head = list1
		list1 = list1.Next
	}

	merged := head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			merged.Next = list2
			list2 = list2.Next
		} else {
			merged.Next = list1
			list1 = list1.Next
		}
		merged = merged.Next
	}

	if list1 != nil {
		merged.Next = list1
	} else if list2 != nil {
		merged.Next = list2
	}

	return
}
