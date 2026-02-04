/**
*
* Список разделов

Учитывая начало связанного списка и значение x, разделите его таким образом, чтобы все узлы, меньшие x, располагались перед узлами, большими или равными x.

Вы должны сохранить исходный относительный порядок расположения узлов в каждом из двух разделов.

Входные данные: head = [1,4,3,2,5,2], x = 3
Выходные данные: [1,2,2,4,3,5]

Пример 2:

Входные данные: head = [2,1], x = 2
Выходные данные: [1,2]



Ограничения:

    Количество узлов в списке находится в диапазоне [0, 200].
    -100 <= Node.val <= 100
    -200 <= x <= 200



*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var m, mEnd, b, bEnd *ListNode

	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			if m == nil {
				m = node
			} else {
				mEnd.Next = node
			}
			mEnd = node
		} else {
			if b == nil {
				b = node
			} else {
				bEnd.Next = node
			}
			bEnd = node
		}
	}

	if mEnd != nil {
		mEnd.Next = b
	}
	if bEnd != nil {
		bEnd.Next = nil
	}

	if m == nil {
		return b
	}

	return m
}

func showList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Print("\n")

}

func main() {
	node6 := &ListNode{Val: 2, Next: nil}
	node5 := &ListNode{Val: 5, Next: node6}
	node4 := &ListNode{Val: 2, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	node0 := &ListNode{Val: 2, Next: node1}

	showList(node0)

	newNode0 := partition(node0, 3)
	showList(newNode0)
}
