package main

import (
	"fmt"
)

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：head = [1,2,2,1]
// 输出：true
// 示例 2：

// 输入：head = [1,2]
// 输出：false

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	p := head
	n := 0
	for p != nil {
		p = p.Next
		n++
	}

	half := n / 2
	var pre *ListNode = nil
	p = head
	next := head
	for i := 0; i < half; i++ {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}

	pre_half := pre
	post_half := p
	if n%2 != 0 {
		post_half = p.Next
	}
	for pre_half != nil && post_half != nil {
		if pre_half.Val != post_half.Val {
			return false
		}
		pre_half = pre_half.Next
		post_half = post_half.Next
	}
	return true
}

func main() {
	node4 := ListNode{1, nil}
	node3 := ListNode{2, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	// fmt.Println(node1)
	// fmt.Println(node1.Next)
	// fmt.Println(node1.Next.Next)
	// fmt.Println(node1.Next.Next.Next)

	ret := isPalindrome(&node1)

	fmt.Println(ret)
}
