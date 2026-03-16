package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head.Next
	for slow != quick {
		if quick == nil || quick.Next == nil {
			return false
		}
		slow = slow.Next
		quick = quick.Next.Next
	}
	return true
}

func main() {
	node5 := ListNode{-4, nil}
	node4 := ListNode{0, &node5}
	node3 := ListNode{2, &node4}
	node2 := ListNode{3, &node3}
	node1 := ListNode{1, &node2}
	node5.Next = &node3

	fmt.Println(node1)
	fmt.Println(node1.Next)
	fmt.Println(node1.Next.Next)
	fmt.Println(node1.Next.Next.Next)
	fmt.Println(node1.Next.Next.Next.Next)

	ret := hasCycle(&node1)

	fmt.Println(ret)
}
