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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a_addr_map := map[*ListNode]int{}
	for headA.Next != nil {
		a_addr_map[headA] = headA.Val
		headA = headA.Next
	}

	for headB.Next != nil {
		if _, ok := a_addr_map[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil

}

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{8, &node4}
	node2 := ListNode{1, &node3}
	node1 := ListNode{4, &node2}

	fmt.Println(node1)
	fmt.Println(node1.Next)
	fmt.Println(node1.Next.Next)
	fmt.Println(node1.Next.Next.Next)
	fmt.Println(node1.Next.Next.Next.Next)

	headA := &node1

	nodeB2 := ListNode{1, &node3}
	nodeB1 := ListNode{6, &nodeB2}
	nodeB0 := ListNode{5, &nodeB1}

	headB := &nodeB0

	fmt.Println(headB)
	fmt.Println(headB.Next)
	fmt.Println(headB.Next.Next)
	fmt.Println(headB.Next.Next.Next)
	fmt.Println(headB.Next.Next.Next.Next)
	fmt.Println(headB.Next.Next.Next.Next.Next)

	ret := getIntersectionNode(headA, headB)

	fmt.Println(ret)
	fmt.Println(ret.Next)
	fmt.Println(ret.Next.Next)
}
