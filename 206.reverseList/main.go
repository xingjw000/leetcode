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

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	fmt.Println(node1)
	fmt.Println(node1.Next)
	fmt.Println(node1.Next.Next)
	fmt.Println(node1.Next.Next.Next)
	fmt.Println(node1.Next.Next.Next.Next)

	ret := reverseList(&node1)

	fmt.Println(ret)
	fmt.Println(ret.Next)
	fmt.Println(ret.Next.Next)
	fmt.Println(ret.Next.Next.Next)
	fmt.Println(ret.Next.Next.Next.Next)
}
