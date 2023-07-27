package linkedlistcycleii

import (
	"testing"
)

func TestCycle(t *testing.T) {

	nodeThree := &ListNode{
		Val: 3,
	}
	nodeTwo := &ListNode{
		Next: nodeThree,
		Val:  2,
	}
	nodeOne := &ListNode{
		Next: nodeTwo,
		Val:  1,
	}
	head := ListNode{
		Next: nodeOne,
		Val:  0,
	}

	nodeThree.Next = nodeOne

	// head -> nodeOne -> nodeTwo -> nodeThree -> nodeOne

	cycleStart := nodeOne
	result := detectCycle(&head)

	if result != cycleStart {
		t.Error("\nTest Failed: the result did not equal the node at the start of the cycle\n\n")
	}

}

func TestNoCycle(t *testing.T) {

	head := new(ListNode)
	nodeOne := new(ListNode)
	nodeTwo := new(ListNode)

	head.Next = nodeOne
	nodeOne.Next = nodeTwo

	// head -> nodeOne -> nodeTwo

	result := detectCycle(head)

	if result != nil {
		t.Error("\nTest Failed: received a non-nil response for a linked list with no cycle\n\n")
	}
}

func TestHeadPointsToHead(t *testing.T) {

	head := new(ListNode)
	head.Next = head

	// head -> head

	result := detectCycle(head)
	if result != head {
		t.Error("\nTest Failed: the head points to its self but head was not returned\n\n")
	}
}

func TestIfNodeOnePointsToNodeOne(t *testing.T) {

	nodeOne := &ListNode{
		Val: 1,
	}
	head := ListNode{
		Next: nodeOne,
		Val:  0,
	}

	nodeOne.Next = nodeOne

	// head -> nodeOne -> nodeOne

	cycleStart := nodeOne
	result := detectCycle(&head)

	if result != cycleStart {
		t.Error("\nTest Failed: nodeOne points to its self but nodeOne was not returned\n\n", result)
	}

}
