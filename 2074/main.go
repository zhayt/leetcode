package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	l := length(head)
	k := 1
	head.Next = reverse(head.Next, l, k)
	return head
}

func reverse(node *ListNode, l int, k int) *ListNode {
RE:
	k++
	l -= k
	if l < k {
		return node
	}

	if k%2 == 1 {
		node = skipGroup(node, k)
		l -= k
		goto RE
	}

	var prev, next *ListNode
	curr := node

	for i := 0; i < k; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	node.Next = reverse(node, l, k)

	return prev
}

func length(node *ListNode) int {
	var count int
	for node != nil {
		count++
		node = node.Next
	}

	return count
}

func skipGroup(node *ListNode, k int) *ListNode {
	for k != 0 {
		k--
		node = node.Next
	}

	return node
}

func main() {
	head := &ListNode{}
	curr := head
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			curr.Val = j
			curr.Next = &ListNode{}
			curr = curr.Next
		}
		if i == 5 {
			curr = nil
		}
	}

	fmt.Println(reverseEvenLengthGroups(head))
	for i := head; i != nil; i = i.Next {
		fmt.Print(head.Val, "->")
	}
}
