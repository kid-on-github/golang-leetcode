package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

// func Constructor() Codec {

// }

// Serializes a tree to a single string.
func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	values := ""
	nodeQueue := []TreeNode{*root}

	for len(nodeQueue) > 0 {
		current := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		fmt.Println(current.Val)
		// values += fmt.Sprint(' ',current.Val)

		children := []*TreeNode{current.Left, current.Right}

		for _, child := range children {
			if child != nil {
				nodeQueue = append(nodeQueue, *child)
			}
		}
	}

	return values
}

func main() {
	rr := &TreeNode{Val: 5}
	rl := &TreeNode{Val: 4}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3, Left: rl, Right: rr}
	head := &TreeNode{Val: 1, Left: l, Right: r}
	serialize(head)
}

//  // Deserializes your encoded data to tree.
//  func (this *Codec) deserialize(data string) *TreeNode {

//  }

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
