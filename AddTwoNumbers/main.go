package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := l1
	last := l1
	memory := false

	node1 := l1
	node2 := l2

	for node1 != nil || node2 != nil {
		val := 0

		if node1 != nil {
			val += node1.Val
		}

		if node2 != nil {
			val += node2.Val
		}

		if memory {
			memory = false
			val++
		}

		if val >= 10 {
			val %= 10
			memory = true
		}

		last = node1
		if last == nil {
			last = node2
			result = l2
		}

		if node1 != nil {
			node1.Val = val
			node1 = node1.Next
		}

		if node2 != nil {
			node2.Val = val
			node2 = node2.Next
		}
	}

	if memory {
		last.Next = &ListNode{
			Val: 1,
		}
	}

	return result
}

func main() {
	arr1 := []int{9, 9, 9}
	arr2 := []int{9, 9}

	list1 := &ListNode{}
	list2 := &ListNode{}
	h1 := list1
	h2 := list2

	for i, n := range arr1 {
		newL := &ListNode{Val: n}
		list1.Next = newL
		list1 = newL
		if i == len(arr1)-1 {
			list1.Next = nil
			break
		}
	}
	for i, n := range arr2 {
		newL := &ListNode{Val: n}
		list2.Next = newL
		list2 = newL
		if i == len(arr2)-1 {
			list2.Next = nil
			break
		}
	}

	result := addTwoNumbers(h1.Next, h2.Next)

	for i := result; i != nil; i = i.Next {
		fmt.Print(i.Val, "->")
	}
}
