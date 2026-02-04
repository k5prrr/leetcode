/**
 * Definition for singly-linked list.
 *
 */
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nEnd := head
	for node := head; node != nil; node = node.Next {
		if nEnd.Val != node.Val {
			nEnd.Next = node
			nEnd = node
		}
	}
	nEnd.Next = nil

	return head
}

func showList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Print("\n")
}

func main() {
	node6 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 3, Next: node6}
	node4 := &ListNode{Val: 3, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 2, Next: node2}
	node0 := &ListNode{Val: -2, Next: node1}

	showList(node0)

	newNode0 := deleteDuplicates(node0)

	showList(newNode0)
}
