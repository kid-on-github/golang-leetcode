package narytreepreordertraversal

type Node struct {
	Val      int
	Children []*Node
}

type LinkedList struct {
	Next     *LinkedList
	TreeNode *Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	head := LinkedList{
		TreeNode: root,
		Next:     nil,
	}

	current := &head

	for current != nil {
		tmpNext := &current.Next
		for _, childNode := range current.TreeNode.Children {
			newLLNode := LinkedList{
				TreeNode: childNode,
				Next:     current.Next,
			}

			current.Next = &newLLNode
			current = current.Next
		}

		current = *tmpNext
	}

	current = &head
	valueList := []int{}
	for current != nil {
		valueList = append(valueList, current.TreeNode.Val)
		current = current.Next
	}

	return valueList
}