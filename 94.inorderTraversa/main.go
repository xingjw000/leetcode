package main

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	// dfs(root, &result)

	node := root
	stack := list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		e := stack.Back()
		node = e.Value.(*TreeNode)
		result = append(result, node.Val)
		stack.Remove(e)
		node = node.Right
	}
	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, result)

	*result = append(*result, node.Val)

	dfs(node.Right, result)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node5 := TreeNode{3, nil, nil}
	node4 := TreeNode{2, &node5, nil}
	node3 := TreeNode{1, nil, &node4}

	fmt.Println(inorderTraversal(&node3))

}
